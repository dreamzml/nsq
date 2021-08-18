package nsqadmin

import(
	"fmt"
	"time"
	"net/http"
	"encoding/json"
	"github.com/jmoiron/sqlx"
	"github.com/julienschmidt/httprouter"
	"github.com/nsqio/nsq/internal/protocol"
	"github.com/nsqio/nsq/internal/dao"
	"github.com/nsqio/nsq/internal/http_api"
	"github.com/nsqio/nsq/internal/clusterinfo"
)

type RestChannelsController struct {
}

func (s *httpServer) restChannelsHandler(w http.ResponseWriter, req *http.Request, ps httprouter.Params) (interface{}, error) {
	var messages []string
	var restChannels []clusterinfo.RestChannel

	queryArgsArr := map[string]interface{}{
		"limit": "1000",
		"offset": "0",
		"is_deleted": "0",
	}
	
	//查询
	sql := "SELECT id,topic,channel,method,rest_url,content_type,created_at FROM rest_channels WHERE is_deleted=:is_deleted limit :limit offset :offset"
	nstmt, err := dao.MDB.PrepareNamed(sql)
	if err != nil {
		fmt.Println(err)
	}
	if err := nstmt.Select(&restChannels, queryArgsArr); err != nil {
		fmt.Println(err)
	}

	return struct {
		RestChannels  []clusterinfo.RestChannel `json:"restChannels"`
		Message string   `json:"message"`
	}{restChannels, maybeWarnMsg(messages)}, nil
}

func (s *httpServer) createRestChannelHandler(w http.ResponseWriter, req *http.Request, ps httprouter.Params) (interface{}, error) {
	var messages []string

	var body clusterinfo.RestChannel

	if !s.isAuthorizedAdminRequest(req) {
		return nil, http_api.Err{403, "FORBIDDEN"}
	}

	err := json.NewDecoder(req.Body).Decode(&body)
	if err != nil {
		return nil, http_api.Err{400, err.Error()}
	}

	if !protocol.IsValidTopicName(body.Topic) {
		return nil, http_api.Err{400, "INVALID_TOPIC"}
	}

	if len(body.Channel) > 0 && !protocol.IsValidChannelName(body.Channel) {
		return nil, http_api.Err{400, "INVALID_CHANNEL"}
	}

	err = s.ci.CreateTopicChannel(body.Topic, body.Channel,
		s.nsqadmin.getOpts().NSQLookupdHTTPAddresses)
	if err != nil {
		pe, ok := err.(clusterinfo.PartialErr)
		if !ok {
			s.nsqadmin.logf(LOG_ERROR, "failed to create topic/channel - %s", err)
			return nil, http_api.Err{502, fmt.Sprintf("UPSTREAM_ERROR: %s", err)}
		}
		s.nsqadmin.logf(LOG_WARN, "%s", err)
		messages = append(messages, pe.Error())
	}

	body.CreateAt = time.Now()

	fmt.Println(body, time.Now().Format("2006-01-02 15:04:05"))
	s.notifyAdminAction("create_topic", body.Topic, "", "", req)
	if len(body.Channel) > 0 {
		s.notifyAdminAction("create_channel", body.Topic, body.Channel, "", req)
	}
	sql := "INSERT INTO rest_channels (`topic`, `channel`, `method`, rest_url, content_type, created_at) VALUES (:topic, :channel, :method, :rest_url, :content_type, :created_at)"
	if result, err := dao.MDB.NamedExec(sql, body); err != nil{
		fmt.Println("insert error", err)
		messages = append(messages, err.Error())
	}else{
		id, _ := result.LastInsertId()
		s.nsqadmin.logf(LOG_INFO, "create rest channel success, new id is %d", id)
	}

	return struct {
		Message string `json:"message"`
	}{maybeWarnMsg(messages)}, nil
}

func (s *httpServer) deleteRestChannelHandler(w http.ResponseWriter, req *http.Request, ps httprouter.Params) (interface{}, error) {
	var messages []string

	if !s.isAuthorizedAdminRequest(req) {
		return nil, http_api.Err{403, "FORBIDDEN"}
	}

	var channel clusterinfo.RestChannel
	id := ps.ByName("id")

	if err := dao.MDB.Get(&channel, "SELECT id,topic,channel,method,rest_url,content_type,created_at FROM rest_channels WHERE id=? LIMIT 1;", id); err !=nil{
		s.nsqadmin.logf(LOG_ERROR, "not fond delete record, id: %d, error:  %s",id, err)
		return nil, http_api.Err{502, fmt.Sprintf("not fond record: %s", err)}
	}

	topicName := channel.Topic
	channelName := channel.Channel

	err := s.ci.DeleteChannel(topicName, channelName,
		s.nsqadmin.getOpts().NSQLookupdHTTPAddresses,
		s.nsqadmin.getOpts().NSQDHTTPAddresses)
	if err != nil {
		pe, ok := err.(clusterinfo.PartialErr)
		if !ok {
			s.nsqadmin.logf(LOG_ERROR, "failed to delete channel - %s", err)
			return nil, http_api.Err{502, fmt.Sprintf("UPSTREAM_ERROR: %s", err)}
		}
		s.nsqadmin.logf(LOG_WARN, "%s", err)
		messages = append(messages, pe.Error())
	}
	s.notifyAdminAction("delete_channel", topicName, channelName, "", req)

	query, args, _ := sqlx.In("UPDATE rest_channels SET is_deleted=1 WHERE id = ?;", id)
	query = dao.MDB.Rebind(query)
	_, err = dao.MDB.Exec(query, args...)
	if err != nil {
		fmt.Println("delete fail failed:", query, args, err)
		messages = append(messages, err.Error())
	}

	return struct {
		Message string `json:"message"`
	}{maybeWarnMsg(messages)}, nil
}



package nsqlookupd

import(
	"fmt"
	"time"
	"net/http"
	"github.com/julienschmidt/httprouter"
	"github.com/nsqio/nsq/internal/dao"
)

func (s *httpServer) doRestChannels(w http.ResponseWriter, req *http.Request, ps httprouter.Params) (interface{}, error) {

	var restChannels []restChannel
	queryArgsArr := map[string]interface{}{
		"limit": "1000",
		"offset": "0",
		"is_deleted": "0",
	}
	
	//查询
	sql := "SELECT id,topic,channel,method,rest_url,created_at FROM rest_channels WHERE is_deleted=:is_deleted limit :limit offset :offset"
	nstmt, err := dao.MDB.PrepareNamed(sql)
	if err != nil {
		fmt.Println(err)
	}
	if err := nstmt.Select(&restChannels, queryArgsArr); err != nil {
		fmt.Println(err)
	}

	return map[string]interface{}{
		"rest_channels": restChannels,
	}, nil
}

type restChannel struct{
	ID int `db:"id"`
	Topic string `db:"topic" json:"topic"`
	Channel string `db:"channel" json:"channel"`
	Method string `db:"method" json:"method"`
	RestUrl string `db:"rest_url" json:"rest_url"`
	CreateAt time.Time `db:"created_at" json:"created_at"`
}
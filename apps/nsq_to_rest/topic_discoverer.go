package main

import (
	"os"
	"regexp"
	"sync"
	"time"
	"reflect"
	"github.com/nsqio/go-nsq"
	"github.com/nsqio/nsq/internal/clusterinfo"
	"github.com/nsqio/nsq/internal/http_api"
	"github.com/nsqio/nsq/internal/lg"
)

type TopicDiscoverer struct {
	logf     lg.AppLogFunc
	opts     *Options
	ci       *clusterinfo.ClusterInfo
	rests    map[string]*RequestWorker
	hupChan  chan os.Signal
	termChan chan os.Signal
	wg       sync.WaitGroup
	cfg      *nsq.Config
}

func newTopicDiscoverer(logf lg.AppLogFunc, opts *Options, cfg *nsq.Config, hupChan chan os.Signal, termChan chan os.Signal) *TopicDiscoverer {
	client := http_api.NewClient(nil, opts.HTTPClientConnectTimeout, opts.HTTPClientRequestTimeout)
	return &TopicDiscoverer{
		logf:     logf,
		opts:     opts,
		ci:       clusterinfo.New(nil, client),
		rests:    make(map[string]*RequestWorker),
		hupChan:  hupChan,
		termChan: termChan,
		cfg:      cfg,
	}
}

func (t *TopicDiscoverer) updateTopics(restChannels map[string]*clusterinfo.RestChannel) {
	var oldKeys = reflect.ValueOf(t.rests).MapKeys()
	for key, rest := range restChannels {
		oldKeys = remove(oldKeys, key)
		if _, ok := t.rests[key]; ok {
			continue
		}

		// if !t.isTopicAllowed(topic) {
		// 	t.logf(lg.WARN, "skipping topic %s (doesn't match pattern %s)", topic, t.opts.TopicPattern)
		// 	continue
		// }

		fl, err := NewRequestWorker(t.logf, t.opts, rest, t.cfg, t.ci)
		if err != nil {
			t.logf(lg.ERROR, "couldn't create NewRequestWorker for new topic %s: %s", key, err)
			continue
		}
		t.rests[key] = fl

		t.wg.Add(1)
		go func(fl *RequestWorker) {
			fl.router()
			t.wg.Done()
		}(fl)
	}

	//关闭已删除的restchannel
	for _, remainKey := range oldKeys{
		key := remainKey.String()
		t.logf(lg.INFO, "close rest channel %s", key)
		close(t.rests[key].termChan)
		delete(t.rests, key)
	}
}

//删除字符串数组中指定的元素
func remove(list []reflect.Value, str string) ([]reflect.Value) {
	for k,v := range list{
	    if v.String() == str{
			return append(list[:k], (list[k+1:])...)
	    }
	}
	return list
}


func (t *TopicDiscoverer) run() {
	var ticker <-chan time.Time
	if len(t.opts.Topics) == 0 {
		ticker = time.Tick(t.opts.TopicRefreshInterval)
	}
	//t.updateTopics(t.opts.Topics)
forloop:
	for {
		select {
		case <-ticker:
			newTopics, err := t.ci.GetRestChannels(t.opts.NSQLookupdHTTPAddrs)
			if err != nil {
				t.logf(lg.ERROR, "could not retrieve topic list: %s", err)
				continue
			}
			t.updateTopics(newTopics)
		case <-t.termChan:
			for _, fl := range t.rests {
				close(fl.termChan)
			}
			t.logf(lg.INFO, "app stop...")
			break forloop
		case <-t.hupChan:
			for _, fl := range t.rests {
				fl.hupChan <- true
			}
		}
	}
	t.wg.Wait()
}

func (t *TopicDiscoverer) isTopicAllowed(topic string) bool {
	if t.opts.TopicPattern == "" {
		return true
	}
	match, err := regexp.MatchString(t.opts.TopicPattern, topic)
	if err != nil {
		return false
	}
	return match
}

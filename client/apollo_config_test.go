package client

import (
	"fmt"
	"sync"
	"testing"

	"github.com/apolloconfig/agollo/v4"
	"github.com/apolloconfig/agollo/v4/env/config"
	"github.com/apolloconfig/agollo/v4/storage"
)

func TestApolloConfig(t *testing.T) {
	conf := &config.AppConfig{
		AppID:         "SampleApp",
		Cluster:       "default",
		NamespaceName: "application",
		IP:            "http://localhost:8080",
	}
	client, _ := agollo.StartWithConfig(func() (*config.AppConfig, error) {
		return conf, nil
	})
	fmt.Println("初始化配置成功")
	cache := client.GetConfigCache(conf.NamespaceName)
	value, _ := cache.Get("timeout")
	fmt.Println(value)

	listener := &CustomChangeListener{}
	listener.wg.Add(2)
	client.AddChangeListener(listener)
	listener.wg.Wait()
}

type CustomChangeListener struct {
	wg sync.WaitGroup
}

func (c *CustomChangeListener) OnChange(changeEvent *storage.ChangeEvent) {
	for key, value := range changeEvent.Changes {
		fmt.Println("change key: ", key, ", value: ", value)
	}
	fmt.Println(changeEvent.Namespace)
	c.wg.Done()
}

func (c *CustomChangeListener) OnNewestChange(event *storage.FullChangeEvent) {
	//TODO implement me
	// panic("implement me")
}

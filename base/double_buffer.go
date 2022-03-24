package base

import "sync/atomic"

type Config struct {
	WhiteList map[int]struct{}
}

func getConfig() *Config {
	return config.Load().(*Config)
}

var config atomic.Value

// updateConfig 全量更新，使用双 buffer/RCU 完全消除读阻塞
func updateConfig() {
	conf := &Config{
		WhiteList: make(map[int]struct{}),
	}

	for i := 0; i < 1000; i++ {
		conf.WhiteList[i] = struct{}{}
	}
	config.Store(conf)
}

// updateConfig2 部分更新
func updateConfig2() {
	oldConf := getConfig()
	newConf := &Config{
		WhiteList: make(map[int]struct{}),
	}

	// 复制旧的配置
	for k, v := range oldConf.WhiteList {
		newConf.WhiteList[k] = v
	}

	// add 新的值
	newConf.WhiteList[1111] = struct{}{}
	newConf.WhiteList[2222] = struct{}{}
	config.Store(newConf)
}

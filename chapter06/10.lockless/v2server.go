package main

import (
	"fmt"
	"time"
)

type WebServerV2 struct {
	config *Config // 单颗粒度
	// 指针记得初始化
}

func (ws *WebServerV2) reLoad() {
	ws.config = &Config{
		Content: fmt.Sprintf("%d", time.Now().UnixNano()),
	}
	// 不是改content的值，而是将整个config换成一个新的
}

func (ws *WebServerV2) ReloadWorker() {
	for {
		time.Sleep(10 * time.Millisecond)
		ws.reLoad()
	}
}

func (ws *WebServerV2) Visit() string {
	return ws.config.Content
}

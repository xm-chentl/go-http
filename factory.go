package gohttp

import (
	"fmt"
	"sync"
)

// DEFAULT 默认key
const DEFAULT = "default"

var (
	mx            sync.Mutex
	keyOfInstance = make(map[string]IHttp)
)

// Default 默认实例
func Default() IHttp {
	mx.Lock()
	defer mx.Unlock()

	ins, ok := keyOfInstance[DEFAULT]
	if !ok {
		panic(
			fmt.Sprintf("gohttp default key is not instance"),
		)
	}

	return ins
}

// SetDefault 设置默认实例
func SetDefault(ins IHttp) {
	mx.Lock()
	defer mx.Unlock()

	keyOfInstance[DEFAULT] = ins
}

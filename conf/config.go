package conf

import (
	"log"
	"path"
	"runtime"
	"sync"

	"github.com/BurntSushi/toml"
)

const (
	appFileName = "app.toml"
)

var (
	//服务主配置 单例
	app  = new(AppConfig)
	once sync.Once
)

// GetAppConfig 懒加载方式获取配置单例
func GetAppConfig() *AppConfig {
	lazyInit()
	return app
}

// lazyInit 初始化app单例
func lazyInit() {
	once.Do(func() { mustInit() })
}

// mustInit 创建配置实例，必须成功，否则报错
func mustInit() {
	// 获取当前文件目录
	_, filename, _, _ := runtime.Caller(1)
	dir := path.Dir(filename)
	if _, err := toml.DecodeFile(dir+"/"+appFileName, app); err != nil {
		log.Fatal("parse toml file error: ", err)
	}
}

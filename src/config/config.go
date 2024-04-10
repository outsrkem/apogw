package config

import (
	"flag"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

// InitConfig 初始化配置
func InitConfig() *Config {
	var _cfg Config
	var cfgPath string
	var printVersion bool

	flag.StringVar(&cfgPath, "f", "config.yaml", "Configuration file path")
	flag.BoolVar(&printVersion, "version", false, "print program version")
	flag.Parse()

	if printVersion {
		versions, _ := newVersions(Version, GoVersion, GitCommit)
		versions.Print(versions)
	}

	log.Println("Read configuration file:", cfgPath)

	configData, err := os.ReadFile(cfgPath)
	if err != nil {
		log.Println("读取配置文件失败:", err)
		os.Exit(1)
	}

	err = yaml.Unmarshal(configData, &_cfg)
	if err != nil {
		log.Println("解析配置文件失败:", err)
		os.Exit(1)
	}
	// proxy := _cfg.Apigw.Rroxy
	// redis := _cfg.Apigw.Redis
	// for _, apigw := range proxy {
	//     fmt.Println("")
	//     fmt.Println(apigw.Name)
	//     for _, server := range apigw.Server {
	//         fmt.Println(server.Location.Path)
	//         fmt.Println(server.Location.Backend.Host)
	//         fmt.Println(server.Location.Backend.Url)
	//     }
	// }

	return &_cfg
}

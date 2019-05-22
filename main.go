package main

import (
	"apiserver/router/middleware"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"net/http"
	"os"
	"time"

	"apiserver/config"
	"apiserver/model"
	v "apiserver/pkg/version"
	"apiserver/router"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/spf13/pflag"
)

var (
	cfg     = pflag.StringP("config", "c", "", "apiserver config file path")
	version = pflag.BoolP("version", "v", false, "show version info.")
)

func main() {
	// 命令解析
	pflag.Parse()
	if *version {
		_v := v.Get()
		_marshaled, err := json.MarshalIndent(&_v, "", "  ")
		if err != nil {
			fmt.Printf("%v\n", err)
			os.Exit(1)
		}
		fmt.Println(string(_marshaled))
		return
	}

	// 初始化配置
	if err := config.Init(*cfg); err != nil {
		panic(err)
	}
	// init db
	model.DB.Init()
	defer model.DB.Close()

	// 设置启动模式
	gin.SetMode(viper.GetString("runmode"))
	// Create the Gin engine.
	g := gin.New()

	// Routes.
	router.Load(
		// Cores.
		g,

		// Middlwares.
		middleware.RequestId(),
		middleware.Logging(),
	)

	// Ping the server to make sure the router is working.
	go func() {
		if err := pingServer(); err != nil {
			log.Fatal("The router has no response, or it might took too long to start up.", err)
		}
		log.Info("The router has been deployed successfully.")
	}()

	// Start to listening the incoming requests
	cert := viper.GetString("tls.cert")
	key := viper.GetString("tls.key")
	if cert != "" && key != "" {
		go func() {
			log.Infof("Start to listening the incoming requests on https address: %s", viper.GetString("tls.addr"))
			log.Info(http.ListenAndServeTLS(viper.GetString("tls.addr"), cert, key, g).Error())
		}()
	}

	log.Infof("Start to listening the incoming requests on http address: %s", viper.GetString("addr"))
	log.Info(http.ListenAndServe(viper.GetString("addr"), g).Error())
}

// pingServer pings the http server to make sure the router is working.
func pingServer() error {
	for i := 0; i < viper.GetInt("max_ping_count"); i++ {
		// Ping the server by sending a GET request to `/health`.
		log.Infof("read url from file, the url is: %s", viper.GetString("url"))
		resp, err := http.Get(viper.GetString("url") + "/sd/health")
		if err == nil && resp.StatusCode == 200 {
			return nil
		}

		// Sleep for a second to continue the next ping.
		log.Info("Waiting for the router, retry in 1 second.")
		time.Sleep(time.Second)
	}
	return errors.New("Cannot connect to the router.")
}

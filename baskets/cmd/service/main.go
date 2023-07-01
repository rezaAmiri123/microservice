package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/rezaAmiri123/microservice/baskets/internal/agent"
	"github.com/spf13/viper"
)

func main() {
	config, err := LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	ag, err := agent.NewAgent(config.Config)
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, syscall.SIGINT, syscall.SIGTERM)
	<-sigc
	ag.Shutdown()
}

type cfg struct {
	agent.Config
	// GrpcServerTLSConfig tls.TLSConfig
}

func LoadConfig(path string) (config cfg, err error) {
	// viper.AddConfigPath(path)
	viper.AddConfigPath("/mallbots")
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	if err = viper.ReadInConfig(); err != nil {
		log.Println("we have error")

		log.Println(err.Error())
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			log.Println("error is not ConfigFileNotFoundError")
			return
		}
		err = nil
	}
	// if err != nil {
	// 	return
	// }

	err = viper.Unmarshal(&config.Config)
	log.Println(config.Config)
	return
}

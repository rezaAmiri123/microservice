package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/rezaAmiri123/microservice/service_message/internal/agent"
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
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config.Config)
	return
}

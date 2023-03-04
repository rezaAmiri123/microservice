package main

import (
	"github.com/rezaAmiri123/microservice/users/internal/agent"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/signal"
	"syscall"
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

//func main() {
//	nc, err := nats.Connect("localhost")
//	if err != nil {
//		fmt.Println("Error connecting to NATS:", err)
//		return
//	}
//	defer nc.Close()
//
//	js, err := nc.JetStream()
//	if err != nil {
//		fmt.Println("Error connecting to JetStream:", err)
//		return
//	}
//	_, err = js.AddStream(&nats.StreamConfig{
//		Name:     "stream",
//		Subjects: []string{fmt.Sprintf("%s.>", "stream")},
//	})
//	// Define the consumer configuration
//	ccfg := &nats.ConsumerConfig{
//		Durable:        "my-durable",
//		DeliverSubject: "my-deliver-subject",
//		AckPolicy:      nats.AckExplicitPolicy,
//		AckWait:        5000,
//		MaxDeliver:     10,
//		FilterSubject:  "my-filter-subject",
//	}
//
//	// Create the consumer
//	consumer, err := js.AddConsumer("stream", ccfg)
//	if err != nil {
//		fmt.Println("Error creating consumer:", err)
//		return
//	}
//	fmt.Println(consumer)
//	// Consume messages from the stream
//	//for {
//	//	msg, err := consumer.NextMsg()
//	//	if err != nil {
//	//		fmt.Println("Error receiving message:", err)
//	//		continue
//	//	}
//	//
//	//	// Process the message
//	//	fmt.Println("Received message:", string(msg.Data))
//	//
//	//	// Acknowledge the message
//	//	err = msg.Ack()
//	//	if err != nil {
//	//		fmt.Println("Error acknowledging message:", err)
//	//		continue
//	//	}
//	//}
//}

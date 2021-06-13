package main

import (
	"flag"

	"github.com/Shopify/sarama"

	"github.com/streadway/amqp"

	"cloud-native-programming-with-golang/Chapter05/src/bookingservice/listener"
	"cloud-native-programming-with-golang/Chapter05/src/bookingservice/rest"
	"cloud-native-programming-with-golang/Chapter05/src/lib/configuration"
	"cloud-native-programming-with-golang/Chapter05/src/lib/msgqueue"
	"cloud-native-programming-with-golang/Chapter05/src/lib/msgqueue/kafka"
	"cloud-native-programming-with-golang/Chapter05/src/lib/persistence/dblayer"

	msgqueue_amqp "cloud-native-programming-with-golang/Chapter05/src/lib/msgqueue/amqp"
)

func panicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	var eventListener msgqueue.EventListener
	var eventEmitter msgqueue.EventEmitter

	confPath := flag.String("conf", "../lib/configuration/config.json", "flag to set the path to the configuration json file")
	flag.Parse()

	//extract configuration
	config, _ := configuration.ExtractConfiguration(*confPath)

	switch config.MessageBrokerType {
	case "amqp":
		conn, err := amqp.Dial(config.AMQPMessageBroker)
		panicIfErr(err)

		eventListener, err = msgqueue_amqp.NewAMQPEventListener(conn, "events", "booking")
		panicIfErr(err)

		eventEmitter, err = msgqueue_amqp.NewAMQPEventEmitter(conn, "events")
		panicIfErr(err)
	case "kafka":
		conf := sarama.NewConfig()
		conf.Producer.Return.Successes = true
		conn, err := sarama.NewClient(config.KafkaMessageBrokers, conf)
		panicIfErr(err)

		eventListener, err = kafka.NewKafkaEventListener(conn, []int32{})
		panicIfErr(err)

		eventEmitter, err = kafka.NewKafkaEventEmitter(conn)
		panicIfErr(err)
	default:
		panic("Bad message broker type: " + config.MessageBrokerType)
	}

	dbhandler, _ := dblayer.NewPersistenceLayer(config.Databasetype, config.DBConnection)

	processor := listener.EventProcessor{eventListener, dbhandler}
	go processor.ProcessEvents()

	rest.ServeAPI(config.RestfulEndpoint, dbhandler, eventEmitter)
}

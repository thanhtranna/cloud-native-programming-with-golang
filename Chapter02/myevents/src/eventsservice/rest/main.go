package main

import (
	"flag"
	"fmt"
	"log"

	"cloud-native-programming-with-golang/Chapter02/myevents/src/eventsservice/lib/configuration"
	"cloud-native-programming-with-golang/Chapter02/myevents/src/eventsservice/lib/persistence/dblayer"
	"cloud-native-programming-with-golang/Chapter02/myevents/src/eventsservice/rest/rest"
)

func main() {
	confPath := flag.String("conf", `.\configuration\config.json`, "flag to set the path to the configuration json file")
	flag.Parse()

	//extract configuration
	config, _ := configuration.ExtractConfiguration(*confPath)

	fmt.Println("Connecting to database")
	dbhandler, _ := dblayer.NewPersistenceLayer(config.Databasetype, config.DBConnection)
	//RESTful API start
	log.Fatal(rest.ServeAPI(config.RestfulEndpoint, dbhandler))
}

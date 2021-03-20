package listener

import (
	"fmt"
	"log"

	"gopkg.in/mgo.v2/bson"

	"cloud-native-programming-with-golang/Chapter04/src/contracts"
	"cloud-native-programming-with-golang/Chapter04/src/lib/msgqueue"
	"cloud-native-programming-with-golang/Chapter04/src/lib/persistence"
)

type EventProcessor struct {
	EventListener msgqueue.EventListener
	Database      persistence.DatabaseHandler
}

func (ep *EventProcessor) ProcessEvents() {
	log.Println("listening or events")

	received, errors, err := ep.EventListener.Listen("eventCreated")
	if err != nil {
		panic(err)
	}
	for {
		select {
		case evt := <-received:
			fmt.Printf("got event %T: %s\n", evt, evt)
			ep.handleEvent(evt)
		case err = <-errors:
			fmt.Printf("got error while receiving event: %s\n", err)
		}
	}
}

func (ep *EventProcessor) handleEvent(event msgqueue.Event) {
	switch e := event.(type) {
	case *contracts.EventCreatedEvent:
		log.Printf("event %s created: %s ", e.ID, e)
		if !bson.IsObjectIdHex(e.ID) {
			log.Printf("event %v did not contain valid object ID", e)
			return
		}

		ep.Database.AddEvent(persistence.Event{ID: bson.ObjectIdHex(e.ID), Name: e.Name})
	case *contracts.LocationCreatedEvent:
		log.Printf("location %s created: %v", e.ID, e)
		// TODO: No persistence for locations, yet

	default:
		log.Printf("unknown event type: %T", e)
	}
}

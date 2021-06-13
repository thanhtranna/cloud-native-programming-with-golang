package rest

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"cloud-native-programming-with-golang/Chapter04/src/lib/msgqueue"
	"cloud-native-programming-with-golang/Chapter04/src/lib/persistence"
)

func ServeAPI(listenAddr string, database persistence.DatabaseHandler, eventEmitter msgqueue.EventEmitter) {
	r := mux.NewRouter()
	r.Methods(http.MethodPost).Path("/events/{eventID}/bookings").Handler(&CreateBookingHandler{eventEmitter, database})

	srv := http.Server{
		Handler:      r,
		Addr:         listenAddr,
		WriteTimeout: 2 * time.Second,
		ReadTimeout:  1 * time.Second,
	}

	log.Println("Server listening on port ", listenAddr)

	srv.ListenAndServe()
}

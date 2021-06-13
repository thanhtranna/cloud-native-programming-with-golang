package dblayer

import (
	"testing"
	"time"

	"gopkg.in/mgo.v2/bson"

	"cloud-native-programming-with-golang/Chapter04/src/lib/persistence"
)

func TestExamineDatabaseCalls(t *testing.T) {
	events := generateDemoEventData()
	layer, nil := NewPersistenceLayer(MONGODB, "mongodb://127.0.0.1")
	eventids := make([]string, 2)
	var err error
	for i, event := range events {
		eventids[i], err = layer.AddEvent(event)
		if err != nil {
			t.Fatal(err)
		}
	}
	users := generateDemoUsersData()
	userids := make([]string, 2)
	bookings := generateDemoBookingData(eventids[0], eventids[1])
	for i, user := range users {
		userids[i], err = layer.AddUser(user)
		if err != nil {
			t.Fatal(err)
		}

		finduser, err := layer.FindUser(user.First, user.Last)
		if err != nil {
			t.Fatal("Could not find user", user, err)
		}
		t.Log(finduser)

		err = layer.AddBookingForUser(userids[i], bookings[i])
		if err != nil {
			t.Fatal("Could not add booking", bookings[i], "for userid", bson.ObjectId(userids[i]), "error", err)
		}

		bookings, err := layer.FindBookingsForUser(userids[i])
		if err != nil {
			t.Fatal("Could not find booking for user", userids[i], err)
		}
		t.Log(bookings)
	}
	allevents, err := layer.FindAllAvailableEvents()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(allevents)
}

func generateDemoEventData() []persistence.Event {
	return []persistence.Event{
		{
			Name:      "Pink Floyd Concert",
			Duration:  90,
			StartDate: time.Now().AddDate(0, -6, 0).UnixNano(),
			EndDate:   time.Now().AddDate(0, -2, 0).UnixNano(),
			Location: persistence.Location{
				Name:      "East side opera house",
				Address:   "23 Murphy Street East",
				Country:   "Canada",
				OpenTime:  8,
				CloseTime: 22,
				Halls: []persistence.Hall{
					{
						Name:     "Olive West",
						Location: "Second floor, west wing",
						Capacity: 80,
					},
					{
						Name:     "Golden Leaf",
						Location: "Third floor",
						Capacity: 80,
					},
				},
			},
		},
		{
			Name:      "BackStreet boys Concert",
			Duration:  120,
			StartDate: time.Now().AddDate(0, -8, 0).UnixNano(),
			EndDate:   time.Now().AddDate(0, -2, 0).UnixNano(),
			Location: persistence.Location{
				Name:      "West side opera house",
				Address:   "12 Kevin Street West",
				Country:   "US",
				OpenTime:  7,
				CloseTime: 21,
				Halls: []persistence.Hall{
					{
						Name:     "Picasso",
						Location: "First floor",
						Capacity: 95,
					},
					{
						Name:     "Van Gogh",
						Location: "Third floor",
						Capacity: 120,
					},
				},
			},
		},
	}
}

func generateDemoUsersData() []persistence.User {
	return []persistence.User{
		{
			First: "Joe",
			Last:  "Smith",
			Age:   32,
		},
		{
			First: "Jane",
			Last:  "Doe",
			Age:   34,
		},
	}
}

func generateDemoBookingData(eventid1, eventid2 string) []persistence.Booking {
	return []persistence.Booking{
		{
			Date:    time.Now().UnixNano(),
			Seats:   4,
			EventID: eventid1,
		},
		{
			Date:    time.Now().UnixNano(),
			Seats:   4,
			EventID: eventid2,
		},
	}
}

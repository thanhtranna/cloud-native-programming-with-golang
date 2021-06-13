package mongolayer

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"cloud-native-programming-with-golang/Chapter04/src/lib/persistence"
)

const (
	DB        = "myevents"
	USERS     = "users"
	EVENTS    = "events"
	LOCATIONS = "locations"
)

type MongoDBLayer struct {
	session *mgo.Session
}

func NewMongoDBLayer(connection string) (persistence.DatabaseHandler, error) {
	s, err := mgo.Dial(connection)
	return &MongoDBLayer{
		session: s,
	}, err
}

func (mgoLayer *MongoDBLayer) AddUser(u persistence.User) (string, error) {
	s := mgoLayer.getFreshSession()
	defer s.Close()
	u.ID = bson.NewObjectId()
	return u.ID.Hex(), s.DB(DB).C(USERS).Insert(u)
}

func (mgoLayer *MongoDBLayer) AddEvent(e persistence.Event) (string, error) {
	s := mgoLayer.getFreshSession()
	defer s.Close()

	if !e.ID.Valid() {
		e.ID = bson.NewObjectId()
	}

	if e.Location.ID == "" {
		e.Location.ID = bson.NewObjectId()
	}

	return e.ID.Hex(), s.DB(DB).C(EVENTS).Insert(e)
}

func (mgoLayer *MongoDBLayer) AddLocation(l persistence.Location) (persistence.Location, error) {
	s := mgoLayer.getFreshSession()
	defer s.Close()
	l.ID = bson.NewObjectId()
	err := s.DB(DB).C(LOCATIONS).Insert(l)
	return l, err
}

func (mgoLayer *MongoDBLayer) AddBookingForUser(id string, bk persistence.Booking) error {
	s := mgoLayer.getFreshSession()
	defer s.Close()
	return s.DB(DB).C(USERS).UpdateId(bson.ObjectIdHex(id), bson.M{"$addToSet": bson.M{"bookings": []persistence.Booking{bk}}})
}

func (mgoLayer *MongoDBLayer) FindUser(first string, last string) (persistence.User, error) {
	s := mgoLayer.getFreshSession()
	defer s.Close()
	u := persistence.User{}
	err := s.DB(DB).C(USERS).Find(bson.M{"first": first, "last": last}).One(&u)
	//fmt.Printf("Found %v \n", u.String())
	return u, err
}

func (mgoLayer *MongoDBLayer) FindBookingsForUser(id string) ([]persistence.Booking, error) {
	s := mgoLayer.getFreshSession()
	defer s.Close()
	u := persistence.User{}
	err := s.DB(DB).C(USERS).FindId(bson.ObjectId(id)).One(&u)
	return u.Bookings, err
}

func (mgoLayer *MongoDBLayer) FindEvent(id string) (persistence.Event, error) {
	s := mgoLayer.getFreshSession()
	defer s.Close()
	e := persistence.Event{}

	err := s.DB(DB).C(EVENTS).FindId(bson.ObjectIdHex(id)).One(&e)
	return e, err
}

func (mgoLayer *MongoDBLayer) FindEventByName(name string) (persistence.Event, error) {
	s := mgoLayer.getFreshSession()
	defer s.Close()
	e := persistence.Event{}
	err := s.DB(DB).C(EVENTS).Find(bson.M{"name": name}).One(&e)
	return e, err
}

func (mgoLayer *MongoDBLayer) FindAllAvailableEvents() ([]persistence.Event, error) {
	s := mgoLayer.getFreshSession()
	defer s.Close()
	events := []persistence.Event{}
	err := s.DB(DB).C(EVENTS).Find(nil).All(&events)
	return events, err
}

func (l *MongoDBLayer) FindLocation(id string) (persistence.Location, error) {
	s := l.getFreshSession()
	defer s.Close()
	location := persistence.Location{}
	err := s.DB(DB).C(LOCATIONS).Find(bson.M{"_id": bson.ObjectId(id)}).One(&location)
	return location, err
}

func (l *MongoDBLayer) FindAllLocations() ([]persistence.Location, error) {
	s := l.getFreshSession()
	defer s.Close()
	locations := []persistence.Location{}
	err := s.DB(DB).C(LOCATIONS).Find(nil).All(&locations)
	return locations, err
}

func (mgoLayer *MongoDBLayer) getFreshSession() *mgo.Session {
	return mgoLayer.session.Copy()
}

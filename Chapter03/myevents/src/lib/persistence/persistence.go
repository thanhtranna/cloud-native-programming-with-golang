package persistence

type DatabaseHandler interface {
	AddEvent(Event) (string, error)
	FindEvent([]byte) (Event, error)
	FindEventByName(string) (Event, error)
	FindAllAvailableEvents() ([]Event, error)
}

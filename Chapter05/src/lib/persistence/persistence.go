package persistence

type DatabaseHandler interface {
	AddUser(User) (string, error)
	AddEvent(Event) (string, error)
	AddBookingForUser(string, Booking) error
	AddLocation(Location) (Location, error)
	FindUser(string, string) (User, error)
	FindBookingsForUser(string) ([]Booking, error)
	FindEvent(string) (Event, error)
	FindEventByName(string) (Event, error)
	FindAllAvailableEvents() ([]Event, error)
	FindLocation(string) (Location, error)
	FindAllLocations() ([]Location, error)
}

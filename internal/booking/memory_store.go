package booking

type MemoryStore struct {
	//map seats to booking
	bookings map[string]Booking
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		bookings: map[string]Booking{},
	}
}

func (s *MemoryStore) Book(b Booking) error {
	//seat taken or replace seat

	if _, exists := s.bookings[b.SeatID]; exists {
		return ErrSeatAlreadyBooked
	}
	s.bookings[b.SeatID] = b
	return nil
}

// iterating movies and append movie by ID
func (s *MemoryStore) ListBookings(movieID string) []Booking {

	var result []Booking //finding bookings for the movie
	for _, b := range s.bookings {
		if b.MovieID == movieID {
			result = append(result, b)
		}
	}
	return result
}

package bookingOffline

import "time"

type BookingOfflineCore struct {
	ID        int
	ClassID   string
	UserID    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

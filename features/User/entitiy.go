package user

import (
	"time"
)

type Core struct {
	ID        int
	Name      string
	Role      string
	Email     string
	Password  string
	Member_id []Membership
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Membership struct {
	ID               int
	MembershipStatus string
	Expired          time.Time
	PersonalID       []PersonalData
	CreatedAt        time.Time
	UpdatedAT        time.Time
}

type PersonalData struct {
	ID        int
	Name      string
	Umur      int
	Gender    string
	Weight    int
	Height    int
	Goals     int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Bussiness interface {
	CreateUser(data Core) (resp Core, err error)
}

type Data interface {
	InsertUser(data Core) (resp Core, err error)
}

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
	Umur             int
	Gender           string
	Weight           int
	Height           int
	Goals            int
	MembershipStatus string
	Expired          time.Time
	CreatedAt        time.Time
	UpdatedAT        time.Time
}

type Bussiness interface {
	RegisterUser(data Core) (err error)
	LoginUser(data Core) (user Core, err error)
	GetUsers(data Core) (users []Core, err error)
	GetUserById(id int) (user Core, err error)
	UpdateUser(data Core) error
}

type Data interface {
	InsertUserData(data Core) (id int, err error)
	CheckUser(data Core) (user Core, err error)
	GetData(Core) (user []Core, err error)
	GetDataById(id int) (user Core, err error)
	UpdateUser(data Core) error
	GetUserByEmail(email string) (bool, error)
}

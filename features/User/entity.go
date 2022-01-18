package user

type User struct {
	ID         uint
	Username   string `json: "username"`
	Role       string `json: "role"`
	Email      string `json: "email"`
	Password   string `json: "password"`
	Token      string `json:"token"`
	Created_at string `json: "created_at"`
	Updated_at string `json: "updated_at"`
}

type Bussiness interface {
	GetAllUser(search string) (user []User)
	CreateUser(data User) (resp User, err error)
	EditUser(id int) (usr User, err error)
	LoginUser(user User) (usr User, err error)
}

type Data interface {
	SelectAllUser(search string) (user []User)
	InsertUser(data User) (resp User, err error)
	UpdateUser(id int) (usr User, err error)
	CheckAccount(User) (user User, err error)
}

package data

import (
	user "capstone/backend/features/User"
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name      string `gorm:"type:varchar(224)"`
	Role      string `gorm:"type:varchar(50)`
	Email     string
	Password  string
	Member_id Membership `gorm:"many2many:user_membership"`
}

type Membership struct {
	gorm.Model
	MembershipStatus string
	Expired          time.Time
}

// DTO

func toUserRecord(user user.Core) User {
	return User{
		Model: gorm.Model{
			ID: user.Id,
		},
		Name:      user.Name,
		Role:      user.Role,
		Email:     user.Email,
		Password:  user.Password,
		Member_id: user.Member_id,
		Member_id: toSkillsetRecords(user.Member_id),
	}
}

func toUserCore(u User) user.Core {
	return user.Core{
		Id:        u.ID,
		Name:      u.Name,
		Role:      u.Role,
		Email:     u.Role,
		Password:  u.Password,
		Member_id: u.Member_id,
		Skillsets: toSkillsetCoreList(u.Skillsets),
	}
}

func toUserCoreList(uList []User) []user.Core {
	convertedUser := []user.Core{}

	for _, user := range uList {
		convertedUser = append(convertedUser, toUserCore(user))
	}

	return convertedUser
}

func SeparateUserData(data user.Core) (User, []Skillset, []Experience) {
	user := User{
		Name:     data.Name,
		Role:     data.Role,
		Email:    data.Email,
		Password: data.Password,
	}

	return user, user.Skillsets, user.Experiences
}

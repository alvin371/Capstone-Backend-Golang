package bussiness

import user "capstone/backend/features/User"

type userUseCase struct {
	userData user.Data
}

func NewUserBussiness(uData user.Data) user.Bussiness {
	return &userUseCase{
		userData: uData,
	}
}

func (userU *userUseCase) CreateUser(data user.Core) (resp user.Core, err error) {
	resp, err = userU.userData.InsertUser(data)
	if err != nil {
		return user.Core{}, err
	}
	return user.Core{}, err
}

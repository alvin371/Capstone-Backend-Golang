package bussiness

// var (
// 	mockData             mocks.Data
// 	onlineClassBussiness onlineClass.Bussiness
// 	onlineClassDatas     []onlineClass.OnlineClassCore
// 	onlineClassData      onlineClass.OnlineClassCore
// 	onlineClassLogin     onlineClass.OnlineClassCore
// )

// func TestMain(m *testing.M) {
// 	onlineClassBussiness = NewBussinessOnlineClass(&mockData)

// 	onlineClassDatas = []onlineClass.OnlineClassCore{
// 		{
// 			ID:      1,
// 			Name:    "Alvin",
// 			Day:     "Senin",
// 			Link:    "https//",
// 			Time:    "13.00 - 14.00",
// 			Trainer: "Verrel",
// 			Image:   "local",
// 		},
// 	}

// 	onlineClassData = onlineClass.OnlineClassCore{
// 		Name:    "Alvin",
// 		Day:     "Senin",
// 		Link:    "https//",
// 		Time:    "13.00 - 14.00",
// 		Trainer: "Verrel",
// 		Image:   "local",
// 	}
// 	os.Exit(m.Run())
// }

// func TestCreateonlineClass(t *testing.T) {
// 	t.Run("validate create account", func(t *testing.T) {
// 		mockData.On("InsertAccount", mock.AnythingOfType("onlineClass.OnlineClassCore")).Return(nil).Once()
// 		err := onlineClassBussiness.CreateClass(onlineClass.OnlineClassCore{})
// 		assert.Nil(t, err)
// 	})
// 	t.Run("error create account", func(t *testing.T) {
// 		mockData.On("InsertAccount", mock.AnythingOfType("onlineClass.OnlineClassCore")).Return(errors.New("Error")).Once()
// 		err := onlineClassBussiness.CreateClass(onlineClass.OnlineClassCore{})
// 		assert.NotNil(t, err)
// 	})
// }
// func TestGetAllonlineClass(t *testing.T) {
// 	t.Run("validate get accounts", func(t *testing.T) {
// 		mockData.On("SelectAccount", mock.AnythingOfType("onlineClass.OnlineClassCore")).Return(onlineClassDatas, nil).Once()
// 		resp, err := onlineClassBussiness.GetAllClass(onlineClass.OnlineClassCore{})
// 		assert.Nil(t, err)
// 		assert.Equal(t, len(resp), 1)
// 	})

// 	t.Run("error get accounts", func(t *testing.T) {
// 		mockData.On("SelectAccount", mock.AnythingOfType("onlineClass.OnlineClassCore")).Return(nil, errors.New("error"))
// 		resp, err := onlineClassBussiness.GetAllClass(onlineClass.OnlineClassCore{})
// 		assert.NotNil(t, err)
// 		assert.Nil(t, resp)
// 	})
// }

// func TestGetonlineClassByID(t *testing.T) {
// 	t.Run("validate get account by id", func(t *testing.T) {
// 		mockData.On("SelectAccountByID", mock.AnythingOfType("int")).Return(onlineClass.OnlineClassCore{}, nil).Once()
// 		resp, err := onlineClassBussiness.GetClassById(3)
// 		assert.Nil(t, err)
// 		assert.NotNil(t, resp)
// 	})
// 	t.Run("error get account by id", func(t *testing.T) {
// 		mockData.On("SelectAccountByID", mock.AnythingOfType("int")).Return(onlineClass.OnlineClassCore{}, errors.New("error")).Once()
// 		resp, err := onlineClassBussiness.GetClassById(3)
// 		assert.NotNil(t, err)
// 		assert.Equal(t, 0, int(resp.ID))
// 	})
// }

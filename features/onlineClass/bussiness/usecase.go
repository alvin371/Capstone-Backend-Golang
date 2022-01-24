package bussiness

import (
	"capstone/backend/features/onlineClass"
)

type OnlineClassUseCase struct {
	OcData onlineClass.Data
}

func NewBussinessOnlineClass(ocData onlineClass.Data) onlineClass.Bussiness {
	return &OnlineClassUseCase{
		OcData: ocData,
	}
}

func (ou *OnlineClassUseCase) GetAllClass(data onlineClass.OnlineClassCore) (class []onlineClass.OnlineClassCore, err error) {
	onlineClass, err := ou.OcData.SelectAllClass(data)
	if err != nil {
		return nil, err
	}
	return onlineClass, nil
}
func (ou *OnlineClassUseCase) GetClassById(id int) (onlineClass.OnlineClassCore, error) {
	onlineClassData, err := ou.OcData.SelectClassById(id)

	if err != nil {
		return onlineClass.OnlineClassCore{}, err
	}
	return onlineClassData, nil
}
func (ou *OnlineClassUseCase) CreateClass(data onlineClass.OnlineClassCore) (err error) {
	if err := ou.OcData.InsertClass(data); err != nil {
		return err
	}
	return nil
}
func (ou *OnlineClassUseCase) EditClass(id int) (onlineClass onlineClass.OnlineClassCore, err error) {
	onlineClassData, err := ou.OcData.UpdateClass(id)

	if err != nil {
		return onlineClass, err
	}
	return onlineClassData, nil
}
func (ou *OnlineClassUseCase) DeleteClass(id int) (onlineClass onlineClass.OnlineClassCore, err error) {
	onlineClassData, err := ou.OcData.DestryoClass(id)

	if err != nil {
		return onlineClass, err
	}
	return onlineClassData, nil
}

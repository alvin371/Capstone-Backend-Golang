package bussiness

import (
	offlineClass "capstone/backend/features/offlineClass"
)

type offlineClassUseCase struct {
	OcData offlineClass.Data
}

func NewBussinessOfflineClass(ocData offlineClass.Data) offlineClass.Bussiness {
	return &offlineClassUseCase{
		OcData: ocData,
	}
}

func (ou *offlineClassUseCase) GetAllClass(data offlineClass.OfflineClassCore) (class []offlineClass.OfflineClassCore, err error) {
	offlineClass, err := ou.OcData.SelectAllClass(data)
	if err != nil {
		return nil, err
	}
	return offlineClass, nil
}
func (ou *offlineClassUseCase) GetClassById(id int) (offlineClass.OfflineClassCore, error) {
	offlineClassData, err := ou.OcData.SelectClassById(id)

	if err != nil {
		return offlineClass.OfflineClassCore{}, err
	}
	return offlineClassData, nil
}
func (ou *offlineClassUseCase) CreateClass(data offlineClass.OfflineClassCore) (err error) {
	if err := ou.OcData.InsertClass(data); err != nil {
		return err
	}
	return nil
}
func (ou *offlineClassUseCase) EditClass(id int, data offlineClass.OfflineClassCore) (offlineClass offlineClass.OfflineClassCore, err error) {
	offlineClassData, err := ou.OcData.UpdateClass(id, data)

	if err != nil {
		return offlineClass, err
	}
	return offlineClassData, nil
}
func (ou *offlineClassUseCase) DeleteClass(id int) (offlineClass offlineClass.OfflineClassCore, err error) {
	offlineClassData, err := ou.OcData.DestroyClass(id)

	if err != nil {
		return offlineClass, err
	}
	return offlineClassData, nil
}

package Services

import (
	"net/http"
	"ralali-crud-cake-test/Model"
	"ralali-crud-cake-test/Model/Database"
	"ralali-crud-cake-test/Repository"
)

type (
	ICakeServicesHandler interface {
		checkCake(id uint64) (serviceError *Model.ServiceErrorDto)
		GetCake(pagination uint) (data []Database.Cakes, serviceError *Model.ServiceErrorDto)
		GetCakeById(id uint64) (data Database.Cakes, serviceError *Model.ServiceErrorDto)
		AddCake(data Model.CakeRequestDto) (serviceError *Model.ServiceErrorDto)
		UpdateCake(id uint64, data Model.CakeRequestDto) (serviceError *Model.ServiceErrorDto)
		DeleteCake(id uint64) (serviceError *Model.ServiceErrorDto)
	}

	CakeServicesHandler struct {
		repo Repository.ICakeRepositoryHandler
	}
)

func CakeServicesControllerProvider(repo Repository.ICakeRepositoryHandler) *CakeServicesHandler {
	return &CakeServicesHandler{
		repo: repo,
	}
}

func (h *CakeServicesHandler) checkCake(id uint64) (serviceError *Model.ServiceErrorDto) {
	var (
		err    error
		exists bool
	)
	if exists, err = h.repo.CheckExists(id); err != nil {
		return Model.NewInternalServiceError(err)
	}

	if !exists {
		return Model.NewCustomServiceError("id is not found", err, http.StatusNotFound)
	}

	return nil
}

func (h *CakeServicesHandler) GetCake(pagination uint) (data []Database.Cakes, serviceError *Model.ServiceErrorDto) {
	var err error

	if data, err = h.repo.GetCake(pagination); err != nil {
		return nil, Model.NewInternalServiceError(err)
	}

	return data, nil
}

func (h *CakeServicesHandler) GetCakeById(id uint64) (data Database.Cakes, serviceError *Model.ServiceErrorDto) {
	var (
		err error
	)

	if serviceError = h.checkCake(id); serviceError != nil {
		return data, serviceError
	}

	if data, err = h.repo.GetCakeById(id); err != nil {
		return data, Model.NewInternalServiceError(err)
	}

	return data, nil
}

func (h *CakeServicesHandler) AddCake(data Model.CakeRequestDto) (serviceError *Model.ServiceErrorDto) {
	var err error

	if err = h.repo.AddCake(data); err != nil {
		return Model.NewInternalServiceError(err)
	}

	return nil
}

func (h *CakeServicesHandler) UpdateCake(id uint64, data Model.CakeRequestDto) (serviceError *Model.ServiceErrorDto) {
	var (
		err error
	)

	if serviceError = h.checkCake(id); serviceError != nil {
		return serviceError
	}

	if err = h.repo.UpdateCake(id, data); err != nil {
		return Model.NewInternalServiceError(err)
	}

	return nil
}

func (h *CakeServicesHandler) DeleteCake(id uint64) (serviceError *Model.ServiceErrorDto) {
	var err error

	if serviceError = h.checkCake(id); serviceError != nil {
		return serviceError
	}

	if err = h.repo.DeleteCake(id); err != nil {
		return Model.NewInternalServiceError(err)
	}

	return nil
}

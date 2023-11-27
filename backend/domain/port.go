package domain

import (
	"errors"
	"fmt"
	"saasteamtest/backend/models"
)

var ErrIndexNotFound = errors.New("Not Found")

type PortHandler interface {
	Create(request models.PortRequest) (*models.Port, error)
	Update(request models.PortRequest) (*models.Port, error)
}

type PortServiceInterface interface {
	Save(request models.PortRequest) (*models.Port, error)
	Update(request models.PortRequest) (*models.Port, error)
}

type PortService struct {
	portHandler PortHandler
}

func (ps PortService) Update(request models.PortRequest) (*models.Port, error) {
	updatedPort, err := ps.portHandler.Update(request)
	if err != nil {
		if errors.Is(err, ErrIndexNotFound) {
			return nil, err
		}
		return nil, fmt.Errorf("create: %w", err)
	}
	return updatedPort, nil
}

func NewPortService(p1 PortHandler) PortServiceInterface {
	return PortService{
		portHandler: p1,
	}
}

func (ps PortService) Save(request models.PortRequest) (*models.Port, error) {

	savedPort, err := ps.portHandler.Create(request)
	if err != nil {
		return nil, fmt.Errorf("create: %w", err)
	}
	return savedPort, nil
}

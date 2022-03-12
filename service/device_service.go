package service

import (
	"context"

	"gitlab.com/golang-team-template/monolith/model"
	devicerepo "gitlab.com/golang-team-template/monolith/storage/device"
)

// DeviceService interface
type DeviceService interface {
	CreateDevice(ctx context.Context, userID, modelName string) (*model.Device, error)
}

type deviceServiceImpl struct {
	deviceRepo devicerepo.Repository
}

//NewDeviceService creates a new device service
func NewDeviceService(deviceRepo devicerepo.Repository) DeviceService {
	return &deviceServiceImpl{
		deviceRepo: deviceRepo,
	}
}

func (s *deviceServiceImpl) CreateDevice(ctx context.Context, userID, modelName string) (*model.Device, error) {
	device := model.Device{
		GUID:   "12345",
		UserID: userID,
		Model:  modelName,
	}

	if err := s.deviceRepo.SaveDevice(ctx, &device); err != nil {
		return nil, err
	}

	return &device, nil
}

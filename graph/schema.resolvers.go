package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	db_model "github.com/sugoi-wada/home-device-admin/db/model"
	"github.com/sugoi-wada/home-device-admin/graph/generated"
	"github.com/sugoi-wada/home-device-admin/graph/model"
)

func (r *queryResolver) CpDevices(ctx context.Context) ([]*model.CPDevice, error) {
	db_cp_devices := []*db_model.CPDevice{}
	r.DB.Joins("State").Find(&db_cp_devices)

	cp_devices := []*model.CPDevice{}
	for _, device := range db_cp_devices {
		cp_devices = append(cp_devices, &model.CPDevice{
			ID:                  device.ID,
			DeviceID:            device.DeviceID,
			Nickname:            device.Nickname,
			Power:               device.State.Power,
			Feature:             device.State.Feature,
			Speed:               device.State.Speed,
			Temp:                device.State.Temp,
			InsideTemp:          device.State.InsideTemp,
			Nanoex:              device.State.Nanoex,
			People:              device.State.People,
			OutsideTemp:         device.State.OutsideTemp,
			Pm25:                device.State.PM25,
			OnTimer:             device.State.OnTimer,
			OffTimer:            device.State.OffTimer,
			VerticalDirection:   device.State.VerticalDirection,
			HorizontalDirection: device.State.HorizontalDirection,
			Fast:                device.State.Fast,
			Econavi:             device.State.Econavi,
			Volume:              device.State.Volume,
			DisplayLight:        device.State.DisplayLight,
			Sleep:               device.State.Sleep,
			Dry:                 device.State.Dry,
			SelfClean:           device.State.SelfClean,
		})
	}

	return cp_devices, nil
}

func (r *queryResolver) CpDevice(ctx context.Context, id *uint) (*model.CPDevice, error) {
	dbCpDevice := db_model.CPDevice{}
	result := r.DB.Joins("State").First(&dbCpDevice, id)

	if result.Error != nil {
		return nil, result.Error
	}

	cpDevice := &model.CPDevice{
		ID:                  dbCpDevice.ID,
		DeviceID:            dbCpDevice.DeviceID,
		Nickname:            dbCpDevice.Nickname,
		Power:               dbCpDevice.State.Power,
		Feature:             dbCpDevice.State.Feature,
		Speed:               dbCpDevice.State.Speed,
		Temp:                dbCpDevice.State.Temp,
		InsideTemp:          dbCpDevice.State.InsideTemp,
		Nanoex:              dbCpDevice.State.Nanoex,
		People:              dbCpDevice.State.People,
		OutsideTemp:         dbCpDevice.State.OutsideTemp,
		Pm25:                dbCpDevice.State.PM25,
		OnTimer:             dbCpDevice.State.OnTimer,
		OffTimer:            dbCpDevice.State.OffTimer,
		VerticalDirection:   dbCpDevice.State.VerticalDirection,
		HorizontalDirection: dbCpDevice.State.HorizontalDirection,
		Fast:                dbCpDevice.State.Fast,
		Econavi:             dbCpDevice.State.Econavi,
		Volume:              dbCpDevice.State.Volume,
		DisplayLight:        dbCpDevice.State.DisplayLight,
	}

	return cpDevice, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

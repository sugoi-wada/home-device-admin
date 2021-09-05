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
			GatewayID:           device.GatewayID,
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
		})
	}

	return cp_devices, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

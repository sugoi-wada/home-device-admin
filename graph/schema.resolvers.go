package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/sugoi-wada/home-device-admin/graph/generated"
	"github.com/sugoi-wada/home-device-admin/graph/model"
)

func (r *queryResolver) CpDevices(ctx context.Context) ([]*model.CPDevice, error) {
	cp_devices := []*model.CPDevice{}
	r.DB.Find(&cp_devices)
	return cp_devices, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

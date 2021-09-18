package repo

import (
	"github.com/sugoi-wada/home-device-admin/client/cp_client"
	"gorm.io/gorm"
)

type CPDeviceRepo struct {
	DB     *gorm.DB
	Client *cp_client.Client
}

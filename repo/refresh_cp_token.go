package repo

import (
	"fmt"
	"os"

	"github.com/sugoi-wada/home-device-admin/client/cp_client"
	"github.com/sugoi-wada/home-device-admin/db/model"
	"gorm.io/gorm/clause"
)

func (repo CPDeviceRepo) RefreshCPToken() {
	userLoginResponse, err := repo.Client.UserLogin(cp_client.UserLoginRequest{
		Email:    os.Getenv("CP_EMAIL"),
		Password: os.Getenv("CP_PASSWORD"),
		AppToken: os.Getenv("CP_APP_TOKEN"),
	})

	if err != nil {
		fmt.Println(fmt.Errorf("CPログインに失敗しました。 %v", err))
		return
	}

	cpUser := model.CPUser{
		Email:        os.Getenv("CP_EMAIL"),
		CPToken:      userLoginResponse.CPToken,
		ExpireTime:   userLoginResponse.ExpireTime,
		RefreshToken: userLoginResponse.RefreshToken,
		MVersion:     userLoginResponse.MVersion,
	}

	repo.DB.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "email"}},
		DoUpdates: clause.AssignmentColumns([]string{"cp_token", "expire_time", "refresh_token", "m_version", "updated_at"}),
	}).Create(&cpUser)
}

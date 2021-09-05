package worker

import (
	"fmt"
	"os"

	"github.com/sugoi-wada/home-device-admin/client/cp_client"
	"github.com/sugoi-wada/home-device-admin/db/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type RefreshCPToken struct {
	DB *gorm.DB
}

func (data RefreshCPToken) Run() {
	fmt.Println("[Run] Refresh CPToken...")
	client := cp_client.NewClient()

	userLoginResponse, err := client.UserLogin(cp_client.UserLoginRequest{
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

	data.DB.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "email"}},
		DoUpdates: clause.AssignmentColumns([]string{"cp_token", "expire_time", "refresh_token", "m_version", "updated_at"}),
	}).Create(&cpUser)
}

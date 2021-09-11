package worker

import (
	"fmt"
	"os"

	"github.com/sugoi-wada/home-device-admin/client/cp_client"
	"github.com/sugoi-wada/home-device-admin/db/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type FetchCPDeviceList struct {
	DB     *gorm.DB
	Client *cp_client.Client
}

func (data FetchCPDeviceList) Run() {
	fmt.Println("[Run] Update cp devices status...")

	var cpUser model.CPUser
	userResult := data.DB.First(&cpUser, "email = ?", os.Getenv("CP_EMAIL"))
	if userResult.Error != nil {
		fmt.Println(fmt.Errorf("CPTokenの検索に失敗したため、CPデバイス一覧の取得をキャンセルします。 %v", userResult.Error))
		return
	}

	deviceListResponse, err := data.Client.DeviceList(cpUser.CPToken)

	if err != nil {
		fmt.Println(fmt.Errorf("CPデバイス一覧の取得に失敗しました。 %v", err))
		return
	}

	for _, gw := range deviceListResponse.GWList {
		for _, device := range gw.Devices {
			device := model.CPDevice{
				GatewayID: gw.GWID,
				Auth:      gw.Auth,
				DeviceID:  device.DeviceID,
				Nickname:  device.NickName,
			}
			data.DB.Clauses(clause.OnConflict{
				Columns:   []clause.Column{{Name: "gateway_id"}, {Name: "device_id"}},
				DoUpdates: clause.AssignmentColumns([]string{"auth", "nickname", "updated_at"}),
			}).Create(&device)
		}
	}
}

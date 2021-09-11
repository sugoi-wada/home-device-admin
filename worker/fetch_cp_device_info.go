package worker

import (
	"fmt"
	"os"

	"github.com/sugoi-wada/home-device-admin/client/cp_client"
	"github.com/sugoi-wada/home-device-admin/client/cp_client/cmd"
	"github.com/sugoi-wada/home-device-admin/db/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type FetchCPDeviceInfo struct {
	DB *gorm.DB
}

func (data FetchCPDeviceInfo) Run() {
	fmt.Println("[Run] Update cp devices info...")
	client := cp_client.NewClient()

	var cpUser model.CPUser
	userResult := data.DB.First(&cpUser, "email = ?", os.Getenv("CP_EMAIL"))
	if userResult.Error != nil {
		fmt.Println(fmt.Errorf("CPTokenの検索に失敗したため、CPデバイス情報の更新をキャンセルします。 %v", userResult.Error))
		return
	}

	var cpDevices []model.CPDevice
	devicesResult := data.DB.Find(&cpDevices)
	if devicesResult.Error != nil {
		fmt.Println(fmt.Errorf("CPDevice一覧の取得に失敗したため、CPデバイス情報の更新をキャンセルします。 %v", devicesResult.Error))
		return
	}

	allCommandTypes := cp_client.AllCommandTypes()

	for _, device := range cpDevices {
		commandStatusMap := map[string]cp_client.CommandTypeInfo{}

		for _, commandTypes := range [][]cp_client.CommandType{allCommandTypes[:cp_client.MaxCommandCount], allCommandTypes[cp_client.MaxCommandCount:]} {
			deviceInfoResponse, err := client.DeviceInfo(cpUser.CPToken, device.Auth, cp_client.DeviceInfoRequest{
				DeviceID:     device.DeviceID,
				CommandTypes: commandTypes,
			})
			if err != nil {
				fmt.Println(fmt.Errorf("[Response Error] CPデバイス情報の取得に失敗しました。 DeviceID:%d %v", device.ID, err))
				continue
			}

			for _, info := range deviceInfoResponse.Devices[0].Info {
				commandStatusMap[info.CommandType.CommandType] = info
			}
		}

		state := model.CPDeviceState{
			CPDeviceID:          device.ID,
			Power:               commandStatusMap[cmd.Power].Localize(),
			Feature:             commandStatusMap[cmd.Feature].Localize(),
			Speed:               commandStatusMap[cmd.Speed].Localize(),
			Temp:                commandStatusMap[cmd.Temp].Localize(),
			InsideTemp:          commandStatusMap[cmd.InsideTemp].Localize(),
			Nanoex:              commandStatusMap[cmd.Nanoex].Localize(),
			People:              commandStatusMap[cmd.People].Localize(),
			OutsideTemp:         commandStatusMap[cmd.OutsideTemp].Localize(),
			PM25:                commandStatusMap[cmd.PM25].Localize(),
			OnTimer:             commandStatusMap[cmd.OnTimer].Localize(),
			OffTimer:            commandStatusMap[cmd.OffTimer].Localize(),
			VerticalDirection:   commandStatusMap[cmd.VerticalDirection].Localize(),
			HorizontalDirection: commandStatusMap[cmd.HorizontalDirection].Localize(),
			Fast:                commandStatusMap[cmd.Fast].Localize(),
			Econavi:             commandStatusMap[cmd.Econavi].Localize(),
			Volume:              commandStatusMap[cmd.Volume].Localize(),
			DisplayLight:        commandStatusMap[cmd.DisplayLight].Localize(),
			Sleep:               commandStatusMap[cmd.Sleep].Localize(),
			Dry:                 commandStatusMap[cmd.Dry].Localize(),
			SelfClean:           commandStatusMap[cmd.SelfClean].Localize(),
		}

		data.DB.Clauses(clause.OnConflict{
			Columns: []clause.Column{{Name: "cp_device_id"}},
			DoUpdates: clause.AssignmentColumns([]string{
				"power",
				"feature",
				"speed",
				"temp",
				"inside_temp",
				"nanoex",
				"people",
				"outside_temp",
				"pm25",
				"on_timer",
				"off_timer",
				"vertical_direction",
				"horizontal_direction",
				"fast",
				"econavi",
				"volume",
				"display_light",
				"sleep",
				"dry",
				"self_clean",
				"updated_at",
			}),
		}).Create(&state)
	}
}

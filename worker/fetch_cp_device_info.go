package worker

import (
	"fmt"
	"os"

	"github.com/sugoi-wada/home-device-admin/client/cp_client"
	"github.com/sugoi-wada/home-device-admin/client/cp_client/command"
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
		commandStatusMap := map[string]string{}

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
				commandStatusMap[info.CommandType.CommandType] = info.Status
			}
		}

		state := model.CPDeviceState{
			CPDeviceID:          device.ID,
			Power:               command.EnumParams(command.Power, commandStatusMap[command.Power]),
			Feature:             command.EnumParams(command.Feature, commandStatusMap[command.Feature]),
			Speed:               command.EnumParams(command.Speed, commandStatusMap[command.Speed]),
			Temp:                commandStatusMap[command.Temp],
			InsideTemp:          commandStatusMap[command.InsideTemp],
			Nanoex:              command.EnumParams(command.Nanoex, commandStatusMap[command.Nanoex]),
			People:              command.EnumParams(command.People, commandStatusMap[command.People]),
			OutsideTemp:         commandStatusMap[command.OutsideTemp],
			PM25:                commandStatusMap[command.PM25],
			OnTimer:             commandStatusMap[command.OnTimer],
			OffTimer:            commandStatusMap[command.OffTimer],
			VerticalDirection:   command.EnumParams(command.VerticalDirection, commandStatusMap[command.VerticalDirection]),
			HorizontalDirection: command.EnumParams(command.HorizontalDirection, commandStatusMap[command.HorizontalDirection]),
			Fast:                command.EnumParams(command.Fast, commandStatusMap[command.Fast]),
			Econavi:             command.EnumParams(command.Econavi, commandStatusMap[command.Econavi]),
			Volume:              command.EnumParams(command.Volume, commandStatusMap[command.Volume]),
			DisplayLight:        command.EnumParams(command.DisplayLight, commandStatusMap[command.DisplayLight]),
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
				"updated_at",
			}),
		}).Create(&state)
	}
}

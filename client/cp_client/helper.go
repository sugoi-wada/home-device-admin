package cp_client

import "github.com/sugoi-wada/home-device-admin/client/cp_client/command"

const (
	MaxCommandCount = 10
)

func AllCommandTypes() []CommandType {
	return []CommandType{
		{CommandType: command.Power},
		{CommandType: command.Feature},
		{CommandType: command.Speed},
		{CommandType: command.Temp},
		{CommandType: command.InsideTemp},
		{CommandType: command.Sleep},
		{CommandType: command.Nanoex},
		{CommandType: command.OnTimer},
		{CommandType: command.OffTimer},
		{CommandType: command.VerticalDirection},
		{CommandType: command.HorizontalDirection},
		{CommandType: command.Dry},
		{CommandType: command.SelfClean},
		{CommandType: command.People},
		{CommandType: command.OutsideTemp},
		{CommandType: command.Fast},
		{CommandType: command.Econavi},
		{CommandType: command.Volume},
		{CommandType: command.DisplayLight},
		{CommandType: command.PM25},
	}
}

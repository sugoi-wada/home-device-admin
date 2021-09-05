package cp_client

import (
	"github.com/sugoi-wada/home-device-admin/client/cp_client/command"
)

const (
	MaxCommandCount = 9
)

func AllCommandTypes() []CommandType {
	return []CommandType{
		{CommandType: command.Power},
		{CommandType: command.Feature},
		{CommandType: command.Speed},
		{CommandType: command.Temp},
		{CommandType: command.InsideTemp},
		{CommandType: command.Nanoex},
		{CommandType: command.People},
		{CommandType: command.OutsideTemp},
		{CommandType: command.PM25},
		{CommandType: command.OnTimer},
		{CommandType: command.OffTimer},
		{CommandType: command.VerticalDirection},
		{CommandType: command.HorizontalDirection},
		{CommandType: command.Fast},
		{CommandType: command.Econavi},
		{CommandType: command.Volume},
		{CommandType: command.DisplayLight},
	}
}

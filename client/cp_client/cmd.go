package cp_client

import "github.com/sugoi-wada/home-device-admin/client/cp_client/cmd"

const (
	MaxCommandCount = 10

	UnknownValue = "-1638.3"
)

func AllCommandTypes() []CommandType {
	return []CommandType{
		{CommandType: cmd.Power},
		{CommandType: cmd.Feature},
		{CommandType: cmd.Speed},
		{CommandType: cmd.Temp},
		{CommandType: cmd.InsideTemp},
		{CommandType: cmd.Sleep},
		{CommandType: cmd.Nanoex},
		{CommandType: cmd.OnTimer},
		{CommandType: cmd.OffTimer},
		{CommandType: cmd.VerticalDirection},
		{CommandType: cmd.HorizontalDirection},
		{CommandType: cmd.Dry},
		{CommandType: cmd.SelfClean},
		{CommandType: cmd.People},
		{CommandType: cmd.OutsideTemp},
		{CommandType: cmd.Fast},
		{CommandType: cmd.Econavi},
		{CommandType: cmd.Volume},
		{CommandType: cmd.DisplayLight},
		{CommandType: cmd.PM25},
	}
}

func (c CommandTypeInfo) Localize() string {
	if params, found := cmd.EnumParams(c.CommandType.CommandType); found {
		return params[c.Status]
	}

	if c.Status == UnknownValue {
		return ""
	}

	return c.Status
}

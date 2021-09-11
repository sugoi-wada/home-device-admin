package cmd

const (
	Power               = "0x00"
	Feature             = "0x01"
	Speed               = "0x02"
	Temp                = "0x03"
	InsideTemp          = "0x04"
	Sleep               = "0x05"
	Nanoex              = "0x08"
	OnTimer             = "0x0B"
	OffTimer            = "0x0C"
	VerticalDirection   = "0x0F"
	HorizontalDirection = "0x11"
	Dry                 = "0x17"
	SelfClean           = "0x18"
	People              = "0x19"
	OutsideTemp         = "0x21"
	Fast                = "0x1A"
	Econavi             = "0x1B"
	Volume              = "0x1E"
	DisplayLight        = "0x1F"
	PM25                = "0x37"
)

var switchState = map[string]string{
	"0": "關閉",
	"1": "開啟",
}

var enumParams = map[string]map[string]string{
	Power: {
		"0": "停止",
		"1": "運轉",
	},
	Feature: {
		"0": "冷氣",
		"1": "除濕",
		"2": "清淨",
		"3": "自動",
		"4": "暖氣",
	},
	Speed: {
		"0": "自動",
		"1": "最弱",
		"2": "弱",
		"3": "中",
		"4": "強",
		"5": "最強",
	},
	VerticalDirection: {
		"0": "自動",
		"1": "最上",
		"2": "上",
		"3": "真ん中",
		"4": "下",
		"5": "最下",
	},
	HorizontalDirection: {
		"0": "自動",
		"1": "左右真ん中",
		"2": "左右内側",
		"3": "左右外側",
		"4": "左",
		"5": "やや左",
		"6": "やや右",
		"7": "右",
	},
	People: {
		"0": "關",
		"1": "對人",
		"2": "不對人",
		"3": "自動",
	},
	DisplayLight: {
		"0": "亮",
		"1": "暗",
		"2": "ECO燈滅",
	},
	Sleep:     switchState,
	Fast:      switchState,
	Volume:    switchState,
	Nanoex:    switchState,
	Econavi:   switchState,
	SelfClean: switchState,
	Dry:       switchState,
}

func EnumParams(cmd string) (map[string]string, bool) {
	if params, found := enumParams[cmd]; found {
		return params, found
	}

	return nil, false
}

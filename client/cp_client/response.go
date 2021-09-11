package cp_client

// Response Root types

type Response struct {
	State    string `json:"State"`
	StateMsg string `json:"StateMsg"`
	MVersion string `json:"MVersion"`
}

type DeviceListResponse struct {
	Response
	GWList []struct {
		GatewayID string `json:"GWID"`
		NickName  string `json:"NickName"`
		Auth      string `json:"auth"`
		HSType    string `json:"HSType"`
		ModelID   string `json:"ModelID"`
		City      string `json:"City"`
		Area      string `json:"Area"`
		Devices   []struct {
			DeviceID   string `json:"DeviceID"`
			NickName   string `json:"NickName"`
			DeviceType string `json:"DeviceType"`
			AreaID     string `json:"AreaID"`
			ModelType  string `json:"ModelType"`
			Model      string `json:"Model"`
		} `json:"Devices"`
	} `json:"GWList"`
	CommandList []struct {
		ModelType string `json:"ModelType"`
		JSON      []struct {
			DeviceType      int    `json:"DeviceType"`
			DeviceName      string `json:"DeviceName"`
			ModelType       string `json:"ModelType"`
			ProtocalType    string `json:"ProtocalType"`
			ProtocalVersion string `json:"ProtocalVersion"`
			Timestamp       string `json:"Timestamp"`
			List            []struct {
				CommandType   string          `json:"CommandType"`
				CommandName   string          `json:"CommandName"`
				ParameterType string          `json:"ParameterType"`
				ParameterUnit string          `json:"ParameterUnit"`
				Parameters    [][]interface{} `json:"Parameters"`
			} `json:"list"`
		} `json:"JSON"`
	} `json:"CommandList"`
}

type DeviceInfoResponse struct {
	Status    string       `json:"status"`
	Devices   []DeviceInfo `json:"devices"`
	UpdatedAt string       `json:"updated_at"`
}

type DeviceInfoRequest struct {
	DeviceID     string        `json:"DeviceID"`
	CommandTypes []CommandType `json:"CommandTypes"`
}

type UserLoginRequest struct {
	Email    string `json:"MemId"`
	Password string `json:"PW"`
	AppToken string `json:"AppToken"`
}

type UserLoginResponse struct {
	Response
	CPToken      string `json:"CPToken"`
	ExpireTime   string `json:"ExpireTime"`
	RefreshToken string `json:"RefreshToken"`
}

// Nested types

type DeviceInfo struct {
	DeviceID int32             `json:"DeviceID"`
	Info     []CommandTypeInfo `json:"Info"`
}

type CommandTypeInfo struct {
	CommandType
	Status string `json:"status"`
}

type CommandType struct {
	CommandType string `json:"CommandType"`
}

package cp_client

import (
	"fmt"

	"github.com/go-resty/resty/v2"
	"golang.org/x/xerrors"
)

const (
	baseURL = "https://ems2.panasonic.com.tw"
)

type Client struct {
	RestyClient *resty.Client
}

func NewClient() *Client {
	return &Client{
		RestyClient: resty.New().SetHostURL(baseURL).SetDebug(true),
	}
}

func (c *Client) newRequest() *resty.Request {
	return c.RestyClient.R().EnableTrace().SetHeader("Accept", "application/json")
}

type Response struct {
	State    string `json:"State"`
	StateMsg string `json:"StateMsg"`
	MVersion string `json:"MVersion"`
}

type DeviceListResponse struct {
	Response
	GWList     []Gateway   `json:"GWList"`
	PanaModels []PanaModel `json:"CommandList"`
}

type Gateway struct {
	GWID     string   `json:"GWID"`
	NickName string   `json:"NickName"`
	Auth     string   `json:"auth"`
	HSType   string   `json:"HSType"`
	ModelID  string   `json:"ModelID"`
	City     string   `json:"City"`
	Area     string   `json:"Area"`
	Devices  []Device `json:"Devices"`
}

type Device struct {
	DeviceID   string `json:"DeviceID"`
	NickName   string `json:"NickName"`
	DeviceType string `json:"DeviceType"`
	AreaID     string `json:"AreaID"`
	ModelType  string `json:"ModelType"`
	Model      string `json:"Model"`
}

type PanaModel struct {
	ModelType    string        `json:"ModelType"`
	PanaProducts []PanaProduct `json:"JSON"`
}

type PanaProduct struct {
	DeviceType      int32            `json:"DeviceType"`
	DeviceName      string           `json:"DeviceName"`
	ModelType       string           `json:"ModelType"`
	ProtocalType    string           `json:"ProtocalType"`
	ProtocalVersion string           `json:"ProtocalVersion"`
	Timestamp       string           `json:"Timestamp"`
	Commands        []ProductCommand `json:"list"`
}

type ProductCommand struct {
	CommandType   string `json:"CommandType"`
	CommandName   string `json:"CommandName"`
	ParameterUnit string `json:"ParameterUnit"`
	// ParameterType: enum
	// Parameters
}

type DeviceInfoResponse struct {
	Status    string       `json:"status"`
	Devices   []DeviceInfo `json:"devices"`
	UpdatedAt string       `json:"updated_at"`
}

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

func (c *Client) UserLogin(request UserLoginRequest) (*UserLoginResponse, error) {
	res, err := c.newRequest().
		SetBody(&request).
		SetResult(&UserLoginResponse{}).
		Post("/api/userlogin1")

	if err != nil {
		return nil, xerrors.Errorf("%w", err)
	}

	return res.Result().(*UserLoginResponse), nil
}

func (c *Client) DeviceList(cpToken string) (*DeviceListResponse, error) {
	res, err := c.newRequest().
		SetHeader("CPToken", cpToken).
		SetResult(&DeviceListResponse{}).
		Get("/api/UserGetRegisteredGWList1")

	if err != nil {
		return nil, xerrors.Errorf("%w", err)
	}

	return res.Result().(*DeviceListResponse), nil
}

func (c *Client) DeviceInfo(cpToken string, auth string, request DeviceInfoRequest) (*DeviceInfoResponse, error) {
	if len(request.CommandTypes) > MaxCommandCount {
		return nil, xerrors.New(fmt.Sprintf("一度に渡す CommandTypes の数が多すぎます。%d個以下にしてください。\n", MaxCommandCount))
	}

	res, err := c.newRequest().
		SetHeader("CPToken", cpToken).
		SetHeader("auth", auth).
		SetBody(&[]DeviceInfoRequest{request}).
		SetResult(&DeviceInfoResponse{}).
		Post("/api/DeviceGetInfo")

	if err != nil {
		return nil, xerrors.Errorf("%w", err)
	}

	return res.Result().(*DeviceInfoResponse), nil
}

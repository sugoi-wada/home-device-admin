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

package cp_client

import (
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


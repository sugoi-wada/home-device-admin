package repo

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-resty/resty/v2"
	"github.com/jarcoal/httpmock"

	"github.com/stretchr/testify/suite"
	"github.com/sugoi-wada/home-device-admin/client/cp_client"
	"github.com/sugoi-wada/home-device-admin/db/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type FetchCPDeviceInfoTestSuite struct {
	suite.Suite
	repo *CPDeviceRepo
	mock sqlmock.Sqlmock
}

func (suite *FetchCPDeviceInfoTestSuite) SetupTest() {
	db, mock, _ := sqlmock.New()
	suite.mock = mock
	repo := &CPDeviceRepo{}
	repo.DB, _ = gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), config.GetConf())
	rst := resty.New()
	repo.Client = &cp_client.Client{
		RestyClient: rst,
	}
	httpmock.ActivateNonDefault(rst.GetClient())
	suite.repo = repo
}

func (suite *FetchCPDeviceInfoTestSuite) TearDownTest() {
	db, _ := suite.repo.DB.DB()
	db.Close()
	httpmock.DeactivateAndReset()
}

func TestFetchCPDeviceInfoTestSuite(t *testing.T) {
	suite.Run(t, new(FetchCPDeviceInfoTestSuite))
}

func (suite *FetchCPDeviceInfoTestSuite) FetchCPDeviceInfoが成功するはず() {
	httpmock.RegisterResponder("GET", "/api/DeviceGetInfo",
		func(req *http.Request) (*http.Response, error) {
			reqBody := &cp_client.DeviceInfoRequest{}
			bodyBytes, _ := ioutil.ReadAll(req.Body)

			err := json.Unmarshal(bodyBytes, &reqBody)
			if err != nil {
				return nil, err
			}

			var infos []cp_client.CommandTypeInfo
			for _, commandType := range reqBody.CommandTypes {
				infos = append(infos, cp_client.CommandTypeInfo{Status: "0", CommandType: cp_client.CommandType{CommandType: commandType.CommandType}})
			}

			deviceID, err := strconv.ParseInt(reqBody.DeviceID, 10, 32)
			if err != nil {
				return nil, err
			}

			responder, err := httpmock.NewJsonResponse(200, &cp_client.DeviceInfoResponse{
				Status: "success",
				Devices: []cp_client.DeviceInfo{{
					DeviceID: int32(deviceID),
					Info:     infos,
				}},
			})
			if err != nil {
				return nil, err
			}

			return responder, nil
		})

	userRows := sqlmock.
		NewRows([]string{"id", "email", "cp_token", "expire_time", "refresh_token", "m_version", "created_at", "updated_at"}).
		AddRow(1, "example@example.com", "token", "20210918114235", "", "20210910140206", time.Now(), time.Now())
	suite.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "cp_users" WHERE email = $1 ORDER BY "cp_users"."id" LIMIT 1`,
	)).WithArgs(os.Getenv("CP_EMAIL")).WillReturnRows(userRows)

	deviceRows := sqlmock.
		NewRows([]string{"id", "gateway_id", "device_id", "auth", "nickname", "created_at", "updated_at"}).
		AddRow(1, "test.gateway.id", "1", "test.auth.id", "", time.Now(), time.Now())
	suite.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "cp_devices"`)).WillReturnRows(deviceRows)

	suite.mock.ExpectBegin()
	suite.mock.ExpectQuery(regexp.QuoteMeta(
		`INSERT INTO "cp_device_states" ("cp_device_id","power","feature","speed","temp","inside_temp","nanoex","people","outside_temp","pm25","on_timer","off_timer","vertical_direction","horizontal_direction","fast","econavi","volume","display_light","sleep","dry","self_clean","created_at","updated_at") VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18,$19,$20,$21,$22,$23) ON CONFLICT ("cp_device_id") DO UPDATE SET "power"="excluded"."power","feature"="excluded"."feature","speed"="excluded"."speed","temp"="excluded"."temp","inside_temp"="excluded"."inside_temp","nanoex"="excluded"."nanoex","people"="excluded"."people","outside_temp"="excluded"."outside_temp","pm25"="excluded"."pm25","on_timer"="excluded"."on_timer","off_timer"="excluded"."off_timer","vertical_direction"="excluded"."vertical_direction","horizontal_direction"="excluded"."horizontal_direction","fast"="excluded"."fast","econavi"="excluded"."econavi","volume"="excluded"."volume","display_light"="excluded"."display_light","sleep"="excluded"."sleep","dry"="excluded"."dry","self_clean"="excluded"."self_clean","updated_at"="excluded"."updated_at" RETURNING "id"`,
	)).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	suite.mock.ExpectCommit()
	suite.repo.FetchCPDeviceInfo()

	if err := suite.mock.ExpectationsWereMet(); err != nil {
		suite.T().Errorf("there were unfulfilled expectations: %s", err)
	}
}

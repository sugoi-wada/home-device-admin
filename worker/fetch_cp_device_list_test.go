package worker

import (
	"os"
	"regexp"
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

type FetchCPDeviceListTestSuite struct {
	suite.Suite
	worker *FetchCPDeviceList
	mock   sqlmock.Sqlmock
}

func (suite *FetchCPDeviceListTestSuite) SetupTest() {
	db, mock, _ := sqlmock.New()
	suite.mock = mock
	worker := &FetchCPDeviceList{}
	worker.DB, _ = gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), config.GetConf())
	rst := resty.New()
	worker.Client = &cp_client.Client{
		RestyClient: rst,
	}
	httpmock.ActivateNonDefault(rst.GetClient())
	suite.worker = worker
}

func (suite *FetchCPDeviceListTestSuite) TearDownTest() {
	db, _ := suite.worker.DB.DB()
	db.Close()
	httpmock.DeactivateAndReset()
}

func TestFetchCPDeviceListTestSuite(t *testing.T) {
	suite.Run(t, new(FetchCPDeviceListTestSuite))
}

func (suite *FetchCPDeviceListTestSuite) Testデバイス一覧取得タスクが成功するはず() {
	responder, _ := httpmock.NewJsonResponder(200, httpmock.File("../client/cp_client/mock/device_list.json"))
	httpmock.RegisterResponder("GET", "/api/UserGetRegisteredGWList1", responder)
	userRows := sqlmock.
		NewRows([]string{"id", "email", "cp_token", "expire_time", "refresh_token", "m_version", "created_at", "updated_at"}).
		AddRow(1, "example@example.com", "token", "20210918114235", "", "20210910140206", time.Now(), time.Now())
	suite.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "cp_users" WHERE email = $1 ORDER BY "cp_users"."id" LIMIT 1`,
	)).WithArgs(os.Getenv("CP_EMAIL")).WillReturnRows(userRows)
	suite.mock.ExpectBegin()
	suite.mock.ExpectQuery(regexp.QuoteMeta(
		`INSERT INTO "cp_devices" ("gateway_id","device_id","auth","nickname","created_at","updated_at") VALUES ($1,$2,$3,$4,$5,$6) ON CONFLICT ("gateway_id","device_id") DO UPDATE SET "auth"="excluded"."auth","nickname"="excluded"."nickname","updated_at"="excluded"."updated_at" RETURNING "id"`,
	)).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	suite.mock.ExpectCommit()
	suite.worker.Run()

	if err := suite.mock.ExpectationsWereMet(); err != nil {
		suite.T().Errorf("there were unfulfilled expectations: %s", err)
	}
}

package repo

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-resty/resty/v2"
	"github.com/jarcoal/httpmock"

	"github.com/stretchr/testify/suite"
	"github.com/sugoi-wada/home-device-admin/client/cp_client"
	"github.com/sugoi-wada/home-device-admin/db/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type RefreshCPTokenWorkerTestSuite struct {
	suite.Suite
	repo *CPDeviceRepo
	mock sqlmock.Sqlmock
}

func (suite *RefreshCPTokenWorkerTestSuite) SetupTest() {
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
	suite.repo = repo
	httpmock.ActivateNonDefault(rst.GetClient())
}

func (suite *RefreshCPTokenWorkerTestSuite) TearDownTest() {
	db, _ := suite.repo.DB.DB()
	db.Close()
	httpmock.DeactivateAndReset()
}

func TestRefreshCPTokenTestSuite(t *testing.T) {
	suite.Run(t, new(RefreshCPTokenWorkerTestSuite))
}

func (suite *RefreshCPTokenWorkerTestSuite) RefreshCPTokenが成功するはず() {
	responder, _ := httpmock.NewJsonResponder(200, httpmock.File("../client/cp_client/mock/user_login.json"))
	httpmock.RegisterResponder("POST", "/api/userlogin1", responder)

	rows := sqlmock.NewRows([]string{"id"}).AddRow(1)

	suite.mock.ExpectBegin()
	suite.mock.ExpectQuery(regexp.QuoteMeta(
		`INSERT INTO "cp_users" ("email","cp_token","expire_time","refresh_token","m_version","created_at","updated_at") VALUES ($1,$2,$3,$4,$5,$6,$7) ON CONFLICT ("email") DO UPDATE SET "cp_token"="excluded"."cp_token","expire_time"="excluded"."expire_time","refresh_token"="excluded"."refresh_token","m_version"="excluded"."m_version","updated_at"="excluded"."updated_at" RETURNING "id"`,
	)).WillReturnRows(rows)
	suite.mock.ExpectCommit()
	suite.repo.RefreshCPToken()

	if err := suite.mock.ExpectationsWereMet(); err != nil {
		suite.T().Errorf("there were unfulfilled expectations: %s", err)
	}
}

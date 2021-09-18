package job

import (
	"fmt"

	"github.com/sugoi-wada/home-device-admin/repo"
)

type RefreshCPToken struct {
	Repo *repo.CPDeviceRepo
}

func (data RefreshCPToken) Run() {
	fmt.Println("[Run] Refresh CPToken start.")
	data.Repo.RefreshCPToken()
	fmt.Println("[Run] Refresh CPToken finished.")
}

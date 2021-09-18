package job

import (
	"fmt"

	"github.com/sugoi-wada/home-device-admin/repo"
)

type FetchCPDeviceInfo struct {
	Repo *repo.CPDeviceRepo
}

func (data FetchCPDeviceInfo) Run() {
	fmt.Println("[Run] Update cp devices info start.")
	data.Repo.FetchCPDeviceInfo()
	fmt.Println("[Run] Update cp devices info finished.")
}

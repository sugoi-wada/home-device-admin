package job

import (
	"fmt"

	"github.com/sugoi-wada/home-device-admin/repo"
)

type FetchCPDeviceList struct {
	Repo *repo.CPDeviceRepo
}

func (data FetchCPDeviceList) Run() {
	fmt.Println("[Run] Update cp devices status start.")
	data.Repo.FetchCPDeviceList()
	fmt.Println("[Run] Update cp devices status finished.")
}

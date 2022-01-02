package datetimemodule

import (
	"time"

	"github.com/sqwoteg/sqgo-dwmbar/config"
	"github.com/sqwoteg/sqgo-dwmbar/model"
)

func GetFormattedTime() string {
	return time.Now().Format(config.DATE_FORMAT)
}

func Worker(moduleName string, ch chan model.BarModuleData) {
	for {
		ch <- model.BarModuleData{
			Name:              moduleName,
			Output:            GetFormattedTime(),
			ImmediatelyUpdate: false,
		}
		time.Sleep(time.Second * 10)
	}
}

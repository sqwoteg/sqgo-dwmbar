package datetimemodule

import (
	"time"

	model "../../model"
)

const DATE_FORMAT = "02.01.2006 Mon 15:04"

func GetFormattedTime() string {
	return time.Now().Format(DATE_FORMAT)
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

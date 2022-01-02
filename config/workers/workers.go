package workers

import (
	datetimemodule "github.com/sqwoteg/sqgo-dwmbar/barmodules/datetime"
	"github.com/sqwoteg/sqgo-dwmbar/model"
)

var Workers = map[string]func(string, chan model.BarModuleData){
	// "module_name": module_worker
	"time": datetimemodule.Worker,
}

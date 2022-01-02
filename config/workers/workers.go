package workers

import (
	datetimemodule "github.com/sqwoteg/sqgo-dwmbar/barmodules/datetime"
	layoutmodule "github.com/sqwoteg/sqgo-dwmbar/barmodules/layout"
	"github.com/sqwoteg/sqgo-dwmbar/model"
)

var Workers = map[string]func(string, chan model.BarModuleData){
	// "module_name": module_worker
	// order does not matter, you can configure it in config.go file
	"layout": layoutmodule.Worker,
	"time":   datetimemodule.Worker,
}

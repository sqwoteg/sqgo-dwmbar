package main

import (
	datetimemodule "./barmodules/datetime"

	"./model"
)

const SEPARATOR = "    "
const AFTER_STATUS_TEXT = " "

var moduleWorkers = map[string]func(string, chan model.BarModuleData){
	// "module_name": module_worker
	"time": datetimemodule.Worker,
}

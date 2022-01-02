package main

import (
	"log"
	"strings"
	"time"

	"github.com/BurntSushi/xgb"
	"github.com/BurntSushi/xgb/xproto"
	config "github.com/sqwoteg/sqgo-dwmbar/config"
	ModuleWorkers "github.com/sqwoteg/sqgo-dwmbar/config/workers"
	"github.com/sqwoteg/sqgo-dwmbar/model"
)

var barStatus = make(map[string]model.BarModuleData)

var moduleOrderMap = make(map[string]int)

var modules_len = len(config.MODULES_ORDER)

func setBarStatus(elements map[string]model.BarModuleData, x *xgb.Conn, xRoot xproto.Window) {
	var parts = make([]string, modules_len)
	for _, module := range elements {
		if strings.TrimSpace(module.Output) != "" {
			parts[moduleOrderMap[module.Name]] = module.Output
		}
	}

	status := []byte(strings.Join(parts, config.SEPARATOR) + config.AFTER_STATUS_TEXT)
	xproto.ChangeProperty(x, xproto.PropModeReplace, xRoot, xproto.AtomWmName, xproto.AtomString, 8, uint32(len(status)), status)
}

func setStatusWorker(x *xgb.Conn, xRoot xproto.Window) {
	for {
		setBarStatus(barStatus, x, xRoot)
		time.Sleep(time.Second * 5)
	}
}

func main() {
	// generate module order map by config
	for i, v := range config.MODULES_ORDER {
		moduleOrderMap[v] = i
	}

	// establish connection to X server
	x, err := xgb.NewConn()
	if err != nil {
		log.Fatal(err)
	}
	defer x.Close()
	xRoot := xproto.Setup(x).DefaultScreen(x).Root

	// periodically update status
	go setStatusWorker(x, xRoot)

	// channel to receive updates from modules
	updatesChannel := make(chan model.BarModuleData)

	// run every module's worker
	for workerName, worker := range ModuleWorkers.Workers {
		go worker(workerName, updatesChannel)
	}

	// receive updates from updatesChannel
	var m model.BarModuleData
	for {
		m = <-updatesChannel

		barStatus[m.Name] = model.BarModuleData{Name: m.Name, Output: m.Output}

		if m.ImmediatelyUpdate {
			setBarStatus(barStatus, x, xRoot)
		}
	}
}

package main

import (
	"os/exec"
	"strings"
	"time"

	"./model"
)

var barStatus = make(map[string]model.BarModuleData)

func setBarStatus(elements map[string]model.BarModuleData) {
	var parts []string
	for _, module := range elements {
		if strings.TrimSpace(module.Output) != "" {
			parts = append(parts, module.Output)
		}
	}

	status := strings.Join(parts, SEPARATOR) + AFTER_STATUS_TEXT
	exec.Command("xsetroot", "-name", status).Run()
}

func setStatusWorker() {
	for {
		setBarStatus(barStatus)
		time.Sleep(time.Second * 5)
	}
}

func main() {
	go setStatusWorker() // periodically update status

	updatesChannel := make(chan model.BarModuleData)

	for workerName, worker := range moduleWorkers {
		go worker(workerName, updatesChannel)
	}

	var m model.BarModuleData

	for {
		m = <-updatesChannel

		barStatus[m.Name] = model.BarModuleData{Name: m.Name, Output: m.Output}

		if m.ImmediatelyUpdate {
			setBarStatus(barStatus)
		}
	}
}

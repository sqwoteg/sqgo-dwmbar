package layoutmodule

import (
	"bufio"
	"os/exec"

	"github.com/sqwoteg/sqgo-dwmbar/config"
	"github.com/sqwoteg/sqgo-dwmbar/model"
)

func Worker(moduleName string, ch chan model.BarModuleData) {

	cmd := exec.Command(config.XKB_SWITCH_CMD[0], config.XKB_SWITCH_CMD[1:]...)
	cmdReader, _ := cmd.StdoutPipe()

	scanner := bufio.NewScanner(cmdReader)
	go func() {
		for scanner.Scan() {
			ch <- model.BarModuleData{
				Name:              moduleName,
				Output:            scanner.Text(),
				ImmediatelyUpdate: true,
			}
		}
	}()

	err := cmd.Start()
	if err != nil {
		ch <- model.BarModuleData{
			Name:              moduleName,
			Output:            err.Error(),
			ImmediatelyUpdate: false,
		}
	}

	err = cmd.Wait()
	if err != nil {
		ch <- model.BarModuleData{
			Name:              moduleName,
			Output:            err.Error(),
			ImmediatelyUpdate: false,
		}
	}

}

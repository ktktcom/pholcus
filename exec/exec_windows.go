// +build windows
package exec

import (
	"os"
	"os/exec"
	"os/signal"

	"github.com/ktktcom/pholcus/app/scheduler"
	"github.com/ktktcom/pholcus/config"

	"github.com/ktktcom/pholcus/cmd" // cmd版
	"github.com/ktktcom/pholcus/gui" // gui版
	"github.com/ktktcom/pholcus/web" // web版
)

func Run(which string) {
	exec.Command("cmd.exe", "/c", "title", config.APP_FULL_NAME).Start()
	defer func() {
		scheduler.SaveDeduplication()
	}()

	// 选择运行界面
	switch which {
	case "gui":
		gui.Run()

	case "cmd":
		cmd.Run()

	case "web":
		fallthrough
	default:
		ctrl := make(chan os.Signal, 1)
		signal.Notify(ctrl, os.Interrupt, os.Kill)
		go web.Run()
		<-ctrl
	}
}

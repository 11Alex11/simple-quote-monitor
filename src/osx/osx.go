package osx
import (

	"os/exec"
)

func NotifyOSX(title string,message string){
  alert := exec.Command("terminal-notifier", "-title", title, "-message", message)
  alert.Run()

}

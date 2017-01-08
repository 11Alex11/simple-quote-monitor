package linux
import (
	"fmt"
	"os/exec"
	"os"
)

func NotifyLinux(title string,message string){
		//Get icon image path
		folderPath, err := os.Getwd()
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
		folderPath+="/images/modules_financial.png"
  alert := exec.Command("notify-send", "-i", folderPath,title, message)
  alert.Run()

}

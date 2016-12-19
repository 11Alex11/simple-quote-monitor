package windows
import (
  "log"
  "os"
  "fmt"

  "gopkg.in/toast.v1"
)

// Output message for Windows 10 using go.toast library
func NotifyWindows(title string, message string){
  folderPath, err := os.Getwd()
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
	myNote := toast.Notification{
		AppID: "Quote Monitor",
		Title: title,
		Message: message,
		Icon: folderPath + "\\images\\modules_financial.png",
	}
	err = myNote.Push()
	if err != nil{
		log.Fatalln(err)
	}
}

// +build !linux

package notifydriver

import (
	"github.com/deckarep/gosx-notifier"
	"github.com/gen2brain/beeep"
	"log"
	"runtime"
)

func Notify(title string, body string, img_path string) {
	switch os := runtime.GOOS; os {
	case "darwin":
		OSXNotify(title, body, img_path)
	default:
		BeeepNotify(title, body, img_path)
	}
}

func OSXNotify(title string, body string, img_path string) {
	note := gosxnotifier.NewNotification("Check your Apple Stock!")
	note.Title = title
	note.Subtitle = body
	note.AppIcon = img_path
	note.Sound = gosxnotifier.Ping
	//note.ContentImage = "assets/warning.png"

	err := note.Push()

	//If necessary, check error
	if err != nil {
		log.Println("Uh oh!")
	}
}

func BeeepNotify(title string, body string, img_path string) {
	err := beeep.Notify(title, body, "")
	if err != nil {
		panic(err)
	}
}

package notifydriver

import (
	"github.com/deckarep/gosx-notifier"
	"github.com/gen2brain/beeep"
	"log"
	"runtime"
    "bufio"
    "strings"
    "github.com/dbatbold/beep"
)

func Notify(title string, body string, img_path string) {
	switch os := runtime.GOOS; os {
	case "darwin":
		OSXNotify(title, body, img_path)
	case "linux":
		LinuxNotify(title, body, img_path)
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
	err := beeep.Notify("Notify Test", "Message body", "assets/information.png")
	if err != nil {
		panic(err)
	}
}

func LinuxNotify(title string, body string, img_path string) {
    music := beep.NewMusic("") // output can be file as "music.wav"
    volume := 25

    if err := beep.OpenSoundDevice("default"); err != nil {
        log.Fatal(err)
    }
    if err := beep.InitSoundDevice(); err != nil {
        log.Fatal(err)
    }
    beep.PrintSheet = false
    defer beep.CloseSoundDevice()

    musicScore := `
		VP SA8 SR9 A9HRDE q
    `

    reader := bufio.NewReader(strings.NewReader(musicScore))
    go music.Play(reader, volume)
    music.Wait()
    beep.FlushSoundBuffer()
}


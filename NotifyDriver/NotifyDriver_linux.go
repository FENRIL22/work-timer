// +build linux

package notifydriver

import (
	"bufio"
	"github.com/dbatbold/beep"
	"github.com/gen2brain/beeep"
	"log"
	"strings"
)

func Notify(title string, body string, img_path string) {
	LinuxNotify(title, body, img_path)
}

func BeeepNotify(title string, body string, img_path string) {
	err := beeep.Notify(title, body, "")
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
	BeeepNotify(title, body, img_path)
	music.Wait()
	beep.FlushSoundBuffer()
}

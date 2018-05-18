package main

import (
	"fmt"
	ui "github.com/airking05/termui"
	"github.com/fenril22/work-timer/NotifyDriver"
)

func main() {
	if err := ui.Init(); err != nil {
		panic(err)
	}
	defer ui.Close()

	var Tm Timer
	Tm = NewCountDownTimer()

	g := ui.NewGauge()
	g.Percent = 0
	g.Width = 50
	g.BorderLabel = "Timer"

	strs := []string{
		"[1] IntervalTimer-25-5",
		"[-] IntervalTimer-45-15",
		"[3] CountDown-5",
		"[4] CountDown-15",
		"[-] CountDown-25",
		"[-] CountDown-45",
	}

	ls := ui.NewList()
	ls.Items = strs
	ls.ItemFgColor = ui.ColorYellow
	ls.BorderLabel = "TimerList"
	ls.Height = 7

	ui.Body.AddRows(
		ui.NewRow(
			ui.NewCol(12, 0, g),
		),
		ui.NewRow(
			ui.NewCol(12, 0, ls),
		),
	)

	ui.Body.Align()
	ui.Render(ui.Body)

	ui.Handle("/sys/kbd/<space>", func(e ui.Event) {
		Tm.Freeze()
	})

	ui.Handle("/sys/kbd", func(e ui.Event) {
		str := e.Data.(ui.EvtKbd).KeyStr
		switch str {
		case "1":
			Tm = NewIntervalTimer()
		case "3":
			t := NewCountDownTimer()
			t.Time = 60 * 5
			t.Init()
			Tm = t
		case "4":
			Tm = NewCountDownTimer()
		case "9":
			t := NewCountDownTimer()
			t.Time = 5
			t.Init()
			Tm = t
		case "0":
			notifydriver.Notify("Test", "Test Notify", "")
		}
	})

	ui.Handle("/timer/1s", func(e ui.Event) {
		Tm.Tick()

		g.Label = fmt.Sprint(Tm)
		g.BarColor = Tm.BarColor()
		g.Percent = Tm.Percent()

		ui.Body.Align()
		ui.Render(ui.Body)

	})

	ui.Handle("/sys/kbd/q", func(ui.Event) {
		ui.StopLoop()
	})

	ui.Loop()
}

type Timer interface {
	Tick()
	Freeze()
	Reset()
	BarColor() ui.Attribute
	Percent() int
	String() string
}

type IntervalTimer struct {
	iswork  bool
	WTime   int
	STime   int
	counter int
	freeze  bool
}

func NewIntervalTimer() *IntervalTimer {
	s := new(IntervalTimer)
	s.WTime = 60 * 25
	s.STime = 60 * 15
	s.Init()

	return s
}

func (s *IntervalTimer) Init() {
	s.counter = s.WTime
	s.freeze = true
	s.iswork = true
}

func (s *IntervalTimer) Tick() {
	if !s.freeze {
		if s.counter > 0 {
			s.counter -= 1
		} else {
			s.freeze = !s.freeze
			notifydriver.Notify("Timer Arrived", fmt.Sprintf("%d & %d", s.WTime%60, s.STime)+"Min Arrived", "")
		}
	}
}

func (s *IntervalTimer) Freeze() {
}

func (s *IntervalTimer) Reset() {
}

func (s *IntervalTimer) Percent() int {
	return 25
}

func (s *IntervalTimer) BarColor() ui.Attribute {
	return ui.ColorWhite
}

func (s *IntervalTimer) String() string {
	h := s.counter / 60
	e := s.counter % 60

	if !s.freeze {
		return fmt.Sprintf("%02d:%02d", h, e)
	} else {
		return fmt.Sprintf("- %02d:%02d -", h, e)
	}
}

type CountDownTimer struct {
	Time    int
	counter int
	freeze  bool
}

func NewCountDownTimer() *CountDownTimer {
	s := new(CountDownTimer)
	s.Time = 60 * 15
	s.Init()

	return s
}

func (s *CountDownTimer) Init() {
	s.freeze = true
	s.counter = s.Time
}

func (s *CountDownTimer) Tick() {
	if !s.freeze {
		if s.counter > 0 {
			s.counter -= 1
		} else {
			s.freeze = !s.freeze
			notifydriver.Notify("Timer Arrived", fmt.Sprint(s.Time%60)+" Min Arrived", "")
		}
	}
}

func (s *CountDownTimer) Freeze() {
	s.freeze = !s.freeze
}

func (s *CountDownTimer) Reset() {
	s.Init()
}

func (s *CountDownTimer) Percent() int {
	return s.counter * 100 / s.Time
}

func (s *CountDownTimer) BarColor() ui.Attribute {
	return ui.ColorMagenta
	/*
			ColorDefault Attribute = iota
		    ColorBlack
		    ColorRed
		    ColorGreen
		    ColorYellow
		    ColorBlue
		    ColorMagenta
		    ColorCyan
		    ColorWhite
	*/
}

func (s *CountDownTimer) String() string {
	h := s.counter / 60
	e := s.counter % 60

	if !s.freeze {
		return fmt.Sprintf("%02d:%02d", h, e)
	} else {
		return fmt.Sprintf("- %02d:%02d -", h, e)
	}

}

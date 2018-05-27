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

	var im_enable bool = true
	var Tm Timer
	Tm = NewCountDownTimer()

	g := ui.NewGauge()
	g.Percent = 0
	g.Width = 50
	g.BorderLabel = "Timer"

	strs := []string{
		"[1] IntervalTimer-25-5",
		"[2] IntervalTimer-48-12",
		"[3] CountDown-2",
		"[4] CountDown-5",
		"[5] CountDown-15",
		//"[-] CountDown-25",
		//"[-] CountDown-45",
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
		if im_enable {
			str := e.Data.(ui.EvtKbd).KeyStr
			switch str {
			case "1":
				Tm = NewIntervalTimer()
			case "2":
				t := NewIntervalTimer()
				t.WTime = 60 * 48
				t.STime = 60 * 12
				t.Init()
				Tm = t
			case "3":
				t := NewCountDownTimer()
				t.Time = 60 * 2
				t.Init()
				Tm = t
			case "4":
				t := NewCountDownTimer()
				t.Time = 60 * 5
				t.Init()
				Tm = t
			case "5":
				Tm = NewCountDownTimer()
			case "8":
				t := NewIntervalTimer()
				t.WTime = 10
				t.STime = 5
				t.Init()
				Tm = t
			case "9":
				t := NewCountDownTimer()
				t.Time = 5
				t.Init()
				Tm = t
			case "0":
				notifydriver.Notify("Test", "Test Notify", "")
			}
		}
	})

	ui.Handle("/timer/1s", func(e ui.Event) {
		Tm.Tick()

		if im_enable {
			g.BorderLabel = "Timer"
		} else {
			g.BorderLabel = "Timer - locked -"
		}
		g.Label = fmt.Sprint(Tm)
		g.BarColor = Tm.BarColor()
		g.Percent = Tm.Percent()

		ui.Body.Align()
		ui.Render(ui.Body)

	})

	ui.Handle("/sys/kbd/C-c", func(ui.Event) {
		ui.StopLoop()
	})

	ui.Handle("/sys/kbd/C-\\", func(ui.Event) {
		im_enable = !im_enable
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
	term	int
	freeze  bool
}

func NewIntervalTimer() *IntervalTimer {
	s := new(IntervalTimer)
	s.WTime = 60 * 25
	s.STime = 60 * 5
	s.Init()

	return s
}

func (s *IntervalTimer) Init() {
	s.counter = s.WTime
	s.freeze = true
	s.iswork = true
	s.term = 1
}

func (s *IntervalTimer) Tick() {
	if !s.freeze {
		if s.iswork {
			s.tick_work()
		} else {
			s.tick_sleep()
		}
	}
}

func (s *IntervalTimer) tick_work() {
	if s.counter > 0 {
		s.counter -= 1
	} else {
		s.iswork = false
		notifydriver.Notify("Go To Sleep", fmt.Sprintf("%d", s.WTime%60)+" Min", "")
		if s.term >= 4 {
			notifydriver.Notify("Looooong Work", "sleep long time?", "")
		}
	}
}

func (s *IntervalTimer) tick_sleep() {
	if s.counter < s.STime {
		s.counter++
	} else {
		s.iswork = true
		s.counter = s.WTime
		s.term++
		notifydriver.Notify("Go To Work", fmt.Sprintf("%d", s.STime)+" Min", "")
	}
}

func (s *IntervalTimer) Freeze() {
	s.freeze = !s.freeze
}

func (s *IntervalTimer) Reset() {
	s.Init()
}

func (s *IntervalTimer) Percent() int {
	if s.iswork {
		return s.counter * 100 / s.WTime
	} else {
		return s.counter * 100 / s.STime
	}
}

func (s *IntervalTimer) BarColor() ui.Attribute {
	if s.iswork {
		return ui.ColorRed
	} else {
		return ui.ColorGreen
	}
}

func (s *IntervalTimer) String() string {
	h := s.counter / 60
	e := s.counter % 60

	if !s.freeze {
		return fmt.Sprintf("%02d:%02d | %d", h, e, s.term)
	} else {
		return fmt.Sprintf("- %02d:%02d | %d -", h, e, s.term)
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

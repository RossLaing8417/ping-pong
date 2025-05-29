package ui

import "github.com/gdamore/tcell/v2"

type Screen struct {
	tscreen      tcell.Screen
	defaultStyle tcell.Style
}

func NewScreen() (*Screen, error) {
	screen, err := tcell.NewScreen()
	if err != nil {
		return nil, err
	}
	if err = screen.Init(); err != nil {
		return nil, err
	}
	style := tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite)
	screen.SetStyle(style)
	return &Screen{
		tscreen:      screen,
		defaultStyle: style,
	}, nil
}

func (s *Screen) Destroy() {
	s.tscreen.Fini()
}

func (s *Screen) PollEvent() tcell.Event {
	return s.tscreen.PollEvent()
}

func (s *Screen) Sync() {
	s.tscreen.Sync()
}

func (s *Screen) Clear() {
	s.tscreen.Clear()
}

func (s *Screen) Render() {
	s.tscreen.SetContent(0, 0, 'H', nil, s.defaultStyle)
	s.tscreen.SetContent(1, 0, 'i', nil, s.defaultStyle)
	s.tscreen.SetContent(2, 0, '!', nil, s.defaultStyle)
	s.tscreen.Show()
}

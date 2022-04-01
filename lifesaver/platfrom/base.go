package platform

import "log"

type PlatformCtrl interface {
	Snapshot(roomID string, options ...Option) (*Snapshot, error)
	ParserType() string
	WithRoomOn() Option
	DefaultOptions(p PlatformCtrl) []Option
}

type Base struct {
}

func (b *Base) ParserType() string {
	return "ffmpeg"
}

func (b *Base) DefaultOptions(p PlatformCtrl) []Option {
	return []Option{
		p.WithRoomOn(),
	}
}

func (b *Base) WithRoomOn() Option {
	return func(s *Snapshot) {
		log.Println("aa")
	}
}

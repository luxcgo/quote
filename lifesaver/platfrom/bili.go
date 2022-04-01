package platform

import (
	"errors"
	"log"
)

type BilibiliCtrl struct {
	Base
}

func (c *BilibiliCtrl) StreamURL(roomID string) (string, error) {
	snapshot, err := c.Snapshot(roomID)
	if err != nil {
		return "", err
	}
	if snapshot.RoomOn {
		return snapshot.StreamURL, nil
	}
	return "", errors.New("not on air")
}

func (c *BilibiliCtrl) Snapshot(roomID string, options ...Option) (*Snapshot, error) {
	s := &Snapshot{
		RoomID: roomID,
	}

	if options == nil {
		options = c.DefaultOptions(c)
	}

	for _, option := range options {
		option(s)
	}

	return s, nil
}

func (c *BilibiliCtrl) WithRoomOn() Option {

	return func(s *Snapshot) {
		log.Println("hhh")
	}
}

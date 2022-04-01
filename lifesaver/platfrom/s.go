package platform

type Snapshot struct {
	RoomID    string
	RoomName  string
	RoomOn    bool
	StreamURL string
	Err       error
}

type Option func(*Snapshot)

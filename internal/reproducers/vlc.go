package reproducers

import vlcctrl "github.com/CedArctic/go-vlc-ctrl"

type vlcReproducer struct {
	myVLC vlcctrl.VLC
}

func NewVlcReproducer() Reproducer {
	return &vlcReproducer{}
}

func (r vlcReproducer) Init() error {
	//TODO: create config in reproducers
	vlc, err := vlcctrl.NewVLC("127.0.0.1", 8080, "pass123")
	r.myVLC = vlc

	return err
}

func (r vlcReproducer) Start() error {
	return r.Start()
}

func (r vlcReproducer) Stop() error {
	return r.Stop()
}

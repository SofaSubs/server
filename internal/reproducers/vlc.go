package reproducers

import vlcctrl "github.com/CedArctic/go-vlc-ctrl"

type vlcReproducer struct {
	myVLC vlcctrl.VLC
}

func NewVlcReproducer() (Reproducer, error) {
	vlc, err := initVlcReproducer()

	if err != nil {
		return nil, err
	}

	return &vlcReproducer{
		myVLC: *vlc,
	}, nil
}

func initVlcReproducer() (*vlcctrl.VLC, error) {
	//TODO: create config in reproducers
	vlc, err := vlcctrl.NewVLC("127.0.0.1", 8080, "pass123")

	if err != nil {
		return nil, err
	}

	return &vlc, err
}

func (r vlcReproducer) Play() error {
	return r.myVLC.Play()
}

func (r vlcReproducer) Stop() error {
	return r.myVLC.Stop()
}

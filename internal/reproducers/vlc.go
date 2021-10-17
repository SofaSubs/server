package reproducers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	vlcctrl "github.com/CedArctic/go-vlc-ctrl"
)

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

func (r vlcReproducer) Pause() error {
	return r.myVLC.Pause()
}

func (r vlcReproducer) Stop() error {
	return r.myVLC.Stop()
}

func (r vlcReproducer) GetTime() (int, error) {
	client := &http.Client{}
	request, err := http.NewRequest("GET", "http://127.0.0.1:8080/requests/status.json", nil)

	if err != nil {
		return 0, err
	}
	request.SetBasicAuth("", "pass123")
	response, err := client.Do(request)
	if err != nil {
		return 0, err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return 0, err
	}

	var time Time
	json.Unmarshal(body, &time)

	return time.Time, nil
}

type Time struct {
	Time int `json:"time"`
}

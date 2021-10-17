package reproducers

type Reproducer interface {
	Play() error
	Pause() error
	Stop() error
	GetTime() (int, error)
}

package reproducers

type Reproducer interface {
	Play() error
	Stop() error
}

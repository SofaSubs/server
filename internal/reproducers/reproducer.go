package reproducers

type Reproducer interface {
	Init() error
	Start() error
	Stop() error
}

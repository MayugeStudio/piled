package vm

type PiledVM struct {
	inner     Stack
}

func NewVM() *PiledVM {
	return &PiledVM{
		inner: NewStack(),
	}
}


package opengl_exercise

type Input struct {
	keys [256]bool
}

func NewInput() (*Input, error) {
	return &Input{}, nil
}

// KeyDown saves state in the key array if a key is pressed
func (i *Input) KeyDown(key uint32) {
	i.keys[key] = true
}

// KeyUp clears state in the key array if a key is released
func (i *Input) KeyUp(key uint32) {
	i.keys[key] = false
}

// IsKeyDown returns what state the key is in
func (i *Input) IsKeyDown(key uint32) bool {
	return i.keys[key]
}

package components

//Base implements common IRender functionality
type Base struct {
	// Origin position
	x0 int
	y0 int
	// Destination position
	x1 int
	y1 int
	// padding between children
	padding int
}

func (l *Base) GetDimensions() (int, int, int, int) {
	return l.x0, l.y0, l.x1, l.y1
}

func (l *Base) GetPadding() int {
	return l.padding
}

func (l *Base) SetDimensions(x0, y0, x1, y1 int) error {
	// Todo: Validate parameters
	l.x0 = x0
	l.y0 = y0
	l.x1 = x1
	l.y1 = y1
	return nil
}

func (l *Base) SetPadding(p int) error {
	// Todo: Validate parameters
	l.padding = p
	return nil
}

package components

type IRenderer interface {
	GetDimensions() (int, int, int, int)
	GetPadding() int
	SetDimensions(x0, y0, x1, y1 int) error
	SetPadding(p int) error
	Render() error
}

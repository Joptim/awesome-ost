package components

type DummyRenderer struct {
	Base
}

func (d *DummyRenderer) Render() error {
	return nil
}

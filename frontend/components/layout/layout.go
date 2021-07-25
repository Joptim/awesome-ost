package layout

import (
	"fmt"
	"github.com/joptim/awesome-ost/frontend/components"
)

// Layout implements a vertical layout with items centered vertically
// items childrenSizes can be adjusted with height percentages
type Layout struct {
	components.Base
	// Child objects
	children      []components.IRenderer
	childrenSizes []int
}

func (l *Layout) GetChildren() []components.IRenderer {
	return l.children
}

// SetChildren Resets childrenSizes default values
func (l *Layout) SetChildren(children []components.IRenderer) error {
	// Todo: Validate input
	l.children = children
	return nil
}

// GetChildrenSizes return evenly spaced sizes by default
func (l *Layout) GetChildrenSizes() []int {
	// Todo: Return evenly spaced sizes if slice not set
	return l.childrenSizes
}

// SetChildrenSizes set the size of each child. Returns an error if
// sizes are non-positive or do not add up 12.
func (l *Layout) SetChildrenSizes(sizes []int) error {
	if len(sizes) != len(l.children) {
		return fmt.Errorf(
			"length of childrenSizes (%d) do not match length of children instances (%d)",
			len(sizes),
			len(l.children),
		)
	}
	sum := 0
	for _, h := range sizes {
		if h <= 0 {
			return fmt.Errorf("invalid height %d", h)
		}
		sum += h
	}
	// Todo: Remove magic number 12
	expected := 12
	if sum != expected {
		return fmt.Errorf("invalid children sizes %v. Values do not add up %d", sizes, expected)
	}

	l.childrenSizes = make([]int, len(sizes))
	copy(l.childrenSizes, sizes)
	return nil
}

// Todo: define New method that sets default values for layout dimensions

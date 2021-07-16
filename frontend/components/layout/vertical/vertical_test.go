package vertical

import (
	"github.com/joptim/awesome-ost/frontend/components"
	"testing"
)

func TestVertical_Render(t *testing.T) {
	// SetUp
	vLayout := New()
	if err := vLayout.SetDimensions(4, 0, 10, 12); err != nil {
		t.Fatalf("cannot vLayout.SetDimensions")
	}
	if err := vLayout.SetPadding(1); err != nil {
		t.Fatalf("cannot vLayout.SetPadding")
	}
	children := []components.IRenderer{
		&components.DummyRenderer{},
		&components.DummyRenderer{},
		&components.DummyRenderer{},
	}
	if err := vLayout.SetChildren(children); err != nil {
		t.Fatalf("cannot vLayout.SetChildren")
	}
	childrenSizes := []int{3, 6, 3}
	if err := vLayout.SetChildrenSizes(childrenSizes); err != nil {
		t.Fatalf("cannot vLayout.SetChildrenSizes")
	}

	if err := vLayout.Render(); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Check Dimensions
	expectedHeights := [][]int{{1, 3}, {4, 8}, {9, 11}}
	for i, expected := range expectedHeights {
		x0, y0, x1, y1 := children[i].GetDimensions()
		ok := y0 == expected[0] && y1 == expected[1]
		if !ok {
			t.Errorf(
				"with child %d, got (y0, y1) == (%d, %d), expected (%d, %d)",
				i,
				y0,
				y1,
				expected[0],
				expected[1],
			)
		}
		ok = x0 == 4 && x1 == 4
		if !ok {
			t.Errorf(
				"with child %d, got (x0, x1) == (%d, %d), expected (%d, %d)",
				i,
				x0,
				x1,
				4,
				4,
			)
		}
	}
}

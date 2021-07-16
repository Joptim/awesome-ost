package horizontal

import (
	"github.com/joptim/awesome-ost/frontend/components"
	"testing"
)

func TestHorizontal_Render(t *testing.T) {
	// SetUp
	hLayout := New()
	if err := hLayout.SetDimensions(0, 4, 12, 10); err != nil {
		t.Fatalf("cannot hLayout.SetDimensions")
	}
	if err := hLayout.SetPadding(1); err != nil {
		t.Fatalf("cannot hLayout.SetPadding")
	}
	children := []components.IRenderer{
		&components.DummyRenderer{},
		&components.DummyRenderer{},
		&components.DummyRenderer{},
	}
	if err := hLayout.SetChildren(children); err != nil {
		t.Fatalf("cannot hLayout.SetChildren")
	}
	childrenSizes := []int{3, 6, 3}
	if err := hLayout.SetChildrenSizes(childrenSizes); err != nil {
		t.Fatalf("cannot hLayout.SetChildrenSizes")
	}

	if err := hLayout.Render(); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Check Dimensions
	expectedWidths := [][]int{{1, 3}, {4, 8}, {9, 11}}
	for i, expected := range expectedWidths {
		x0, y0, x1, y1 := children[i].GetDimensions()
		ok := x0 == expected[0] && x1 == expected[1]
		if !ok {
			t.Errorf(
				"with child %d, got (x0, x1) == (%d, %d), expected (%d, %d)",
				i,
				x0,
				x1,
				expected[0],
				expected[1],
			)
		}
		ok = y0 == 4 && y1 == 4
		if !ok {
			t.Errorf(
				"with child %d, got (x0, x1) == (%d, %d), expected (%d, %d)",
				i,
				y0,
				y1,
				4,
				4,
			)
		}
	}
}

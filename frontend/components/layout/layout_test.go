package layout

import (
	"github.com/joptim/awesome-ost/frontend/components"
	"reflect"
	"testing"
)

func TestLayout_SetChildren(t *testing.T) {
	expectedChildren := []components.IRenderer{
		&components.DummyRenderer{},
		&components.DummyRenderer{},
		&components.DummyRenderer{},
	}
	layout := &Layout{}
	err := layout.SetChildren(expectedChildren)
	if err != nil {
		t.Fatalf("with ")
	}
	actualChildren := layout.GetChildren()
	if !reflect.DeepEqual(expectedChildren, actualChildren) {
		t.Fatalf(
			"with children %v, got children %v",
			expectedChildren,
			actualChildren,
		)
	}
}

func TestLayout_SetChildrenSizes(t *testing.T) {
	layout := &Layout{}
	err := layout.SetChildren(
		[]components.IRenderer{
			&components.DummyRenderer{},
			&components.DummyRenderer{},
			&components.DummyRenderer{},
		},
	)
	if err != nil {
		t.Fatalf("cannot layout.SetChildren")
	}
	expectedSizes := []int{3, 6, 3}
	err = layout.SetChildrenSizes(expectedSizes)
	if err != nil {
		t.Errorf("with sizes %v, got error %v, expected nil error", expectedSizes, err)
	}
	actualSizes := layout.GetChildrenSizes()
	if !reflect.DeepEqual(expectedSizes, actualSizes) {
		t.Errorf("with sizes %v, got sizes %v", expectedSizes, actualSizes)
	}

}

func TestLayout_SetChildrenSizes_ReturnsErrorIfSizesDontAddTwelve(t *testing.T) {
	layout := &Layout{}
	err := layout.SetChildren(
		[]components.IRenderer{
			&components.DummyRenderer{},
			&components.DummyRenderer{},
		},
	)
	if err != nil {
		t.Fatalf("cannot layout.SetChildren")
	}
	sizes := []int{1, 5}
	err = layout.SetChildrenSizes(sizes)
	if err == nil {
		t.Errorf("with sizes %v, got nil error, expected non-nil error", sizes)
	}
}

func TestLayout_GetChildrenSizes_ReturnsErrorOnDifferentSizes(t *testing.T) {
	layout := &Layout{}
	err := layout.SetChildren(
		[]components.IRenderer{
			&components.DummyRenderer{},
		},
	)
	if err != nil {
		t.Fatalf("cannot layout.SetChildren")
	}
	sizes := []int{6, 6}
	err = layout.SetChildrenSizes(sizes)
	if err == nil {
		t.Errorf("with sizes %v, got nil error, expected non-nil error", sizes)
	}
}

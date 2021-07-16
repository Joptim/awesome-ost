package components

import "testing"

func TestBase_GetDimensions(t *testing.T) {
	base := &Base{}
	err := base.SetDimensions(1, 2, 3, 4)
	if err != nil {
		t.Fatalf("with dimensions (1,2,3,4), got error %v, expected nil-error", err)
	}
	x0, y0, x1, y1 := base.GetDimensions()
	if x0 != 1 || y0 != 2 || x1 != 3 || y1 != 4 {
		t.Errorf(
			"with dimensions (1,2,3,4), got dimensions (%d, %d, %d, %d)",
			x0,
			y0,
			x1,
			y1,
		)
	}
}

func TestBase_GetPadding(t *testing.T) {
	base := &Base{}
	expectedPadding := 5
	err := base.SetPadding(expectedPadding)
	if err != nil {
		t.Fatalf(
			"with padding %d, got error %v, expected nil-error",
			expectedPadding,
			err,
		)
	}
	actualPadding := base.GetPadding()
	if expectedPadding != actualPadding {
		t.Errorf(
			"with padding %d, got error %v, expected nil-error",
			expectedPadding,
			actualPadding,
		)
	}
}

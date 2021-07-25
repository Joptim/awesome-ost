package media

import (
	"path/filepath"
	"reflect"
	"testing"
)

func TestMedia_List(t *testing.T) {
	media := New()
	actual, err := media.List(filepath.Join(
		"..",
		"..",
		"test",
		"server",
		"media",
		"testdata",
	))
	if err != nil {
		t.Fatalf("expected nil-error, got %v", err)
	}
	expected := map[string][]string{
		"folder1": {"file1.txt", "file2.txt"},
		"folder2": {"file3.txt"},
	}
	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("expected %v mapping, got %v", expected, actual)
	}
}

func TestMedia_GetPath(t *testing.T) {
	media := New()
	root := filepath.Join("..", "..", "test", "server", "media", "testdata")
	_, err := media.List(root)
	if err != nil {
		t.Fatalf("expected nil-error, got %v", err)
	}

	table := []struct {
		id       string
		expected string
	}{
		{"file1.txt", filepath.Join(root, "folder1", "file1.txt")},
		{"file2.txt", filepath.Join(root, "folder1", "file2.txt")},
		{"file3.txt", filepath.Join(root, "folder2", "file3.txt")},
	}

	for _, test := range table {
		actual, err := media.GetPath(test.id)
		if err != nil {
			t.Fatalf("with id %s, got error %v, expected nil-error", test.id, err)
		}
		if test.expected != actual {
			t.Errorf("with %s, got path %s, expected path %s", test.id, actual, test.expected)
		}
	}
}

func TestMedia_GetPath_ReturnsErrorOnInvalidId(t *testing.T) {
	media := New()
	root := filepath.Join("..", "..", "test", "server", "media", "testdata")
	_, err := media.List(root)
	if err != nil {
		t.Fatalf("expected nil-error, got %v", err)
	}
	if _, err := media.GetPath("foo.txt"); err == nil {
		t.Errorf("expected non-nil error, got nil error")
	}
}

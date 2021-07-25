package media

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

type Media struct {
	// paths maps IDs to file paths.
	mapping map[string]string
}

// List lists files in a path. The path is expected to have the
// following folder structure:
//    path/
//    |----folder1/
//    |    |----track1.mp3
//    |    |----track2.mp3
//    |----folder2/
//    |    |----track3.mp3
//    |    |----track4.mp3
//    |----folder3/
//	...
func (m *Media) List(path string) (map[string][]string, error) {
	dirMapping := make(map[string][]string)
	directories, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}
	for _, dir := range directories {
		files, err := ioutil.ReadDir(filepath.Join(path, dir.Name()))
		if err != nil {
			return nil, err
		}
		list := make([]string, len(files))
		for i, file := range files {
			list[i] = file.Name()
		}
		dirMapping[dir.Name()] = list
	}
	m.updateIDMapping(path, dirMapping)
	return dirMapping, nil
}

// GetPath returns the file path from an id.
func (m *Media) GetPath(id string) (string, error) {
	if path, exists := m.mapping[id]; exists {
		return path, nil
	} else {
		return "", fmt.Errorf("cannot find file with id %s", id)
	}
}

func (m *Media) updateIDMapping(root string, dirMapping map[string][]string) {
	mapping := make(map[string]string)
	for dir, files := range dirMapping {
		for _, file := range files {
			id := file
			path := filepath.Join(root, dir, file)
			mapping[id] = path
		}
	}
	m.mapping = mapping
}

func New() *Media {
	return &Media{
		mapping: make(map[string]string),
	}
}

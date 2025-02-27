package composefile

import (
	"fmt"
	"os"
	"slices"

	"github.com/compose-spec/compose-go/loader"
	"gopkg.in/yaml.v3"
)

type File struct {
	path string
	data map[string]interface{}
}

func New(path string) (*File, error) {
	composeFileBytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	data, err := loader.ParseYAML(composeFileBytes)
	if err != nil {
		return nil, err
	}

	f := File{
		path: path,
		data: data,
	}

	return &f, nil
}

func (f *File) IncludeAdd(path, envFile string) error {
	var exists bool

	include, ok := f.data["include"].([]interface{})
	if !ok {
		return fmt.Errorf("include field not found or not in expected format")
	}

	for _, item := range include {
		if includeMap, ok := item.(map[string]interface{}); ok {
			for key, value := range includeMap {
				if key == "path" && value == path {
					exists = true
					break
				}
			}
		}
	}

	if !exists {
		newItem := map[string]interface{}{
			"path": path,
		}

		if len(envFile) > 0 {
			newItem["env_file"] = envFile
		}

		include = append(include, newItem)

		f.data["include"] = include
	}

	return nil
}

func (f *File) IncludeRemove(path string) error {
	idx := -1

	include, ok := f.data["include"].([]interface{})
	if !ok {
		return fmt.Errorf("include field not found or not in expected format")
	}

	for i, item := range include {
		if includeMap, ok := item.(map[string]interface{}); ok {
			for key, value := range includeMap {
				if key == "path" && value == path {
					idx = i
					break
				}
			}
		}
	}

	include = slices.Delete(include, idx, idx+1)

	f.data["include"] = include

	return nil
}

func (f *File) Save() error {
	out, err := yaml.Marshal(f.data)
	if err != nil {
		return err
	}

	if err := os.WriteFile(f.path, out, 0640); err != nil {
		return err
	}

	return nil
}

package db

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func (d *Driver) Write(collection, resource string, v interface{}) error {
	if collection == "" || resource == "" {
		return fmt.Errorf("collection and resource required")
	}

	m := d.getOrCreateMutex(collection)
	m.Lock()
	defer m.Unlock()

	dir := filepath.Join(d.dir, collection)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	path := filepath.Join(dir, resource+".json")
	tmp := path + ".tmp"

	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return err
	}

	if err := os.WriteFile(tmp, b, 0644); err != nil {
		return err
	}

	return os.Rename(tmp, path)
}

func (d *Driver) ReadAll(collection string) ([]string, error) {
	if collection == "" {
		return nil, fmt.Errorf("missing collection")
	}

	dir := filepath.Join(d.dir, collection)

	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var records []string
	for _, f := range files {
		b, err := os.ReadFile(filepath.Join(dir, f.Name()))
		if err != nil {
			return nil, err
		}
		records = append(records, string(b))
	}

	return records, nil
}

func (d *Driver) Read(collection, resource string, v interface{}) error {
	if collection == "" {
		return fmt.Errorf("missing collection")
	}
	if resource == "" {
		return fmt.Errorf("missing resource")
	}

	record := filepath.Join(d.dir, collection, resource)

	if _, err := stat(record); os.IsNotExist(err) {
		return fmt.Errorf("record '%s/%s' does not exist", collection, resource)
	}

	b, err := os.ReadFile(record + ".json")
	if err != nil {
		return err
	}

	return json.Unmarshal(b, v)
}

func (d *Driver) Delete(collection, resource string) error {
	if collection == "" {
		return fmt.Errorf("missing collection")
	}
	if resource == "" {
		return fmt.Errorf("missing resource")
	}

	m := d.getOrCreateMutex(collection)
	m.Lock()
	defer m.Unlock()

	path := filepath.Join(d.dir, collection, resource)

	switch fi, err := stat(path); {
	case fi == nil || err != nil:
		return fmt.Errorf("record not found")

	case fi.Mode().IsDir():
		return os.RemoveAll(path)

	case fi.Mode().IsRegular():
		return os.Remove(path + ".json")
	}

	return nil
}

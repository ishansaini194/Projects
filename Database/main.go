package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

const Version = "1.0.0"

type (
	Logger interface {
		Fatal(string, ...interface{})
		Error(string, ...interface{})
		Warn(string, ...interface{})
		Info(string, ...interface{})
		Debug(string, ...interface{})
		Trace(string, ...interface{})
	}

	Driver struct {
		mutex   sync.Mutex
		mutexes map[string]*sync.Mutex
		dir     string
		log     Logger
	}
)

type Options struct {
	Logger
}

type User struct {
	Name    string
	Age     json.Number
	Contact string
	Company string
	Address Address
}

type Address struct {
	City    string
	State   string
	Country string
	Pincode json.Number
}

func New(dir string, options *Options) (*Driver, error) {
	dir = filepath.Clean(dir)

	opts := Options{}
	if options != nil {
		opts = *options
	}

	driver := Driver{
		dir:     dir,
		log:     opts.Logger,
		mutexes: make(map[string]*sync.Mutex),
	}

	if _, err := os.Stat(dir); err == nil {
		if opts.Logger != nil {
			opts.Logger.Debug("Using '%s' (database already exists)\n", dir)
		}
		return &driver, nil
	}

	if opts.Logger != nil {
		opts.Logger.Debug("Creating the database at '%s'...\n", dir)
	}

	return &driver, os.MkdirAll(dir, 0755)
}

func (d *Driver) Write(collection, resource string, v interface{}) error {
	if collection == "" {
		return fmt.Errorf("missing collection")
	}
	if resource == "" {
		return fmt.Errorf("missing resource")
	}

	mutex := d.getOrCreateMutex(collection)
	mutex.Lock()
	defer mutex.Unlock()

	dir := filepath.Join(d.dir, collection)
	finalPath := filepath.Join(dir, resource+".json")
	tmpPath := finalPath + ".tmp"

	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return err
	}
	b = append(b, '\n')

	if err := os.WriteFile(tmpPath, b, 0644); err != nil {
		return err
	}

	return os.Rename(tmpPath, finalPath)
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

func (d *Driver) ReadAll(collection string) ([]string, error) {
	if collection == "" {
		return nil, fmt.Errorf("missing collection")
	}

	dir := filepath.Join(d.dir, collection)

	if _, err := os.Stat(dir); err != nil {
		return nil, err
	}

	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	records := []string{}
	for _, file := range files {
		b, err := os.ReadFile(filepath.Join(dir, file.Name()))
		if err != nil {
			return nil, err
		}
		records = append(records, string(b))
	}

	return records, nil
}

func (d *Driver) Delete(collection, resource string) error {
	if collection == "" {
		return fmt.Errorf("missing collection")
	}
	if resource == "" {
		return fmt.Errorf("missing resource")
	}

	mutex := d.getOrCreateMutex(collection)
	mutex.Lock()
	defer mutex.Unlock()

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

func (d *Driver) getOrCreateMutex(collection string) *sync.Mutex {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	if m, ok := d.mutexes[collection]; ok {
		return m
	}

	m := &sync.Mutex{}
	d.mutexes[collection] = m
	return m
}

func stat(path string) (os.FileInfo, error) {
	if fi, err := os.Stat(path); err == nil {
		return fi, nil
	}
	return os.Stat(path + ".json")
}

func main() {
	dir := "./data"

	db, err := New(dir, nil)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	employees := []User{
		{"Ishan", "23", "8968678456", "notGood", Address{"Bhatoya", "Punjab", "India", "143534"}},
		{"Tushar", "20", "8968678456", "Gov Hospital", Address{"Faridkot", "Punjab", "India", "146001"}},
		{"Kunal", "17", "8968678456", "Gov School", Address{"Dinanagar", "Punjab", "India", "143531"}},
		{"Anubhav", "23", "8968678456", "Pvt School", Address{"Dinanagar", "Punjab", "India", "143531"}},
		{"Uday", "20", "8968678456", "Pvt Uni", Address{"Patiala", "Punjab", "India", "147001"}},
	}

	for _, e := range employees {
		if err := db.Write("user", e.Name, e); err != nil {
			fmt.Println("Write error:", err)
		}
	}

	records, err := db.ReadAll("user")
	if err != nil {
		fmt.Println("ReadAll error:", err)
		return
	}

	allUsers := []User{}
	for _, r := range records {
		u := User{}
		if err := json.Unmarshal([]byte(r), &u); err != nil {
			fmt.Println("Unmarshal error:", err)
		}
		allUsers = append(allUsers, u)
	}

	b, err := json.MarshalIndent(allUsers, "", "  ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(string(b))

	if err := db.Delete("user", "Uday"); err != nil {
		fmt.Println("Delete error:", err)
	}
}

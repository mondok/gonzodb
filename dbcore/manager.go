package dbcore

import (
	"encoding/csv"
	"io/ioutil"
	"os"
)

type Manager struct {
	Schemas   []*Schema
	reader    *csv.Reader
	writer    *csv.Writer
	schemaDir string
}

type Schema struct {
	Name    string
	Columns []*Column
	Rows    []*Row
}

type Row struct {
	Data []interface{}
}

type Column struct {
	Name string
	Type string
}

// NewStore Creates a new DB core object
func NewManager(schemaDir string) *Manager {
	m := &Manager{
		schemaDir: schemaDir,
	}
	return m
}

func (m *Manager) Close() (err error) {
	return
}

func (m *Manager) init() (err error) {
	m.Schemas = []*Schema{}
	if ok, e := exists(m.schemaDir); ok {
		err = m.loadTables()
	} else {
		err = e
	}
	return
}

func (m *Manager) loadTables() error {
	var fis []os.FileInfo
	var err error
	if fis, err = ioutil.ReadDir(m.schemaDir); err != nil {
		return err
	}
	for _, fi := range fis {
		if r, e := os.Open(fi.Name()); e != nil {
			return e
		} else {
			csvReader := csv.NewReader(r)
			defer r.Close()
			csvData, e := csvReader.ReadAll()
			if e != nil {
				return e
			}
			s := &Schema{}
			s.Name = fi.Name()
			for idx, each := range csvData {
				if idx == 0 {
					for _, name := range each {
						col := &Column{}
						col.Name = name
						s.Columns = append(s.Columns, col)
					}
				}
				if idx == 1 {
					for x, typ := range each {
						s.Columns[x].Type = typ
					}
				}
				if idx > 1 {
					row := &Row{}
					l := len(each)
					objs := make([]interface{}, l)
					for _, typ := range each {
						objs = append(objs, typ)
					}
					row.Data = objs
					s.Rows = append(s.Rows, row)
				}
			}
		}
	}
	return nil
}

func exists(path string) (exists bool, err error) {
	exists = false
	_, err = os.Stat(path)
	if err == nil {
		exists = true
	}
	return
}

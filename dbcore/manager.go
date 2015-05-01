package dbcore

import "github.com/mondok/gonzodb/util"

type Manager struct {
	Schemas []*Shema
}

type Shema struct {
	Name    string
	Columns []*Column
	Rows    []*Row
}

type Row struct {
	Data []interface{}
}

type Column struct {
	Name *string
	Type *string
}

// NewStore Creates a new DB core object
func NewManager() *Manager {
	m := &Manager{}
	return m
}

func (m *Manager) init() (err error) {
	m.Schemas = []*Shema{}

	return
}

func (m *Manager) loadTables(err error) {
	return
}

func (m *Manager) initSystemTable() (err error) {
	sysTable := m.table("system")
	if sysTable != nil {
		// create sys schema
	} else {
		// load sys schema
	}
	return
}

func (m *Manager) table(name string) (table *Shema) {
	t, found, _ := util.From(m.Schemas).Where(func(s util.T) (bool, error) {
		return s.(*Shema).Name == name, nil
	}).First()
	if found {
		table = t.(*Shema)
	}
	return
}

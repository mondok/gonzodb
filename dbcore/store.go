package dbcore

import . "github.com/mondok/gonzodb/util"

// Core is the primary file store
// for managing data I/O
type Core struct {
	Tables []*Table
}

type Table struct {
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

const dataDir = "./data"

// NewStore Creates a new DB core object
func NewStore() (c *Core, err error) {
	c = &Core{}
	err = c.init()
	return
}

func (c *Core) init() (err error) {
	// TODO: load from disk
	c.Tables = []*Table{}
	return
}

func (c *Core) loadTables(err error) {
	return
}

func (c *Core) initSystemTable() (err error) {
	sysTable := c.table("system")
	if sysTable != nil {
		// create sys table
	} else {
		// load sys table
	}
	return
}

func (c *Core) table(name string) (table *Table) {
	t, found, _ := From(c.Tables).Where(func(s T) (bool, error) {
		return s.(*Table).Name == name, nil
	}).First()
	if found {
		table = t.(*Table)
	}
	return
}

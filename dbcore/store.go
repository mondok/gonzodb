package dbcore

import (
	"fmt"
	"io"
	"os"

	"github.com/mondok/gonzodb/util"
)

// Core is the primary file store
// for managing data I/O
type Core struct {
	Schemas    []*Shema
	fileHandle io.Writer
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

const dataDir = "./data"
const dataFile = "db.gnz"

// NewStore Creates a new DB core object
func NewStore() (c *Core, err error) {
	c = &Core{}
	err = c.init()
	return
}

func (c *Core) init() (err error) {
	c.Schemas = []*Shema{}
	return
}

func (c *Core) loadTables(err error) {
	fname := fmt.Sprintf("%s/%s", dataDir, dataFile)
	if !util.FileExists(fname) {
		c.fileHandle, _ = os.Create(fname)
	}
	return
}

func (c *Core) initSystemTable() (err error) {
	sysTable := c.table("system")
	if sysTable != nil {
		// create sys schema
	} else {
		// load sys schema
	}
	return
}

func (c *Core) table(name string) (table *Shema) {
	t, found, _ := util.From(c.Schemas).Where(func(s util.T) (bool, error) {
		return s.(*Shema).Name == name, nil
	}).First()
	if found {
		table = t.(*Shema)
	}
	return
}

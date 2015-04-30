package dbcore

// Core is the primary file store
// for managing data I/O
type Core struct {
	Tables []*Table
}

type Table struct {
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
	err = c.initialize()
	return
}

func (c *Core) initialize() (err error) {
	c.Tables = []*Table{}
	return
}

package dbcore

// Core is the primary file store
// for managing data I/O
type Core struct{}

const dataDir = "./data"

// NewStore Creates a new DB core object
func NewStore() (c *Core, err error) {
	c = &Core{}
	err = c.initialize()
	return
}

func (c *Core) initialize() (err error) {
	return
}

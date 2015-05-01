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
	manager    *Manager
	fileHandle io.Writer
}

const dataDir = "./data"
const dataFile = "db.gnz"

// NewStore Creates a new DB core object
func NewStore() (c *Core, err error) {
	c = &Core{}
	c.manager = &Manager{}
	err = c.init()
	return
}

func (c *Core) init() (err error) {
	return
}

func (c *Core) initFromFile(err error) {
	fname := fmt.Sprintf("%s/%s", dataDir, dataFile)
	if !util.FileExists(fname) {
		c.fileHandle, _ = os.Create(fname)
	}
	return
}

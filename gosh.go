package gosh

import (
	"fmt"

	"github.com/cmj0121/argparse"
	"github.com/cmj0121/logger"
)

type Gosh struct {
	argparse.Help
	Version bool `short:"v" help:"show version info" callback:"Ver"`

	// the internal logger
	*logger.Logger `-`
	LogLevel       string `name:"log" choices:"warn info debug verbose" help:"log level"`
}

func New() (gosh *Gosh) {
	gosh = &Gosh{
		Logger: logger.New(PROJ_NAME),
	}
	return
}

func (gosh *Gosh) Run() (err error) {
	parser := argparse.MustNew(gosh)
	err = parser.Run()

	// override the log level
	gosh.Logger.SetLevel(gosh.LogLevel)
	gosh.Info("start run gosh")

	return
}

func (gosh Gosh) Ver(parser *argparse.ArgParse) (exit bool) {
	fmt.Printf("%v (%d.%d.%d)\n", PROJ_NAME, MAJOR, MINOR, MACRO)
	exit = true
	return
}

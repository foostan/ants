package command

import (
	"fmt"
	"github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/foostan/ants/ants"
	"strings"
)

var RunFlags = []cli.Flag{
	cli.StringFlag{
		Name:  "log-level",
		Value: "info",
		Usage: "Log level",
	},
	cli.StringFlag{
		Name:  "config-file",
		Value: "config.json",
		Usage: "configuration file to load",
	},
}

func RunCommand(c *cli.Context) {
	// setup logger
	logLvStr := c.String("log-level")
	logLevel, err := logrus.ParseLevel(logLvStr)
	if err != nil {
		fmt.Errorf("Error: %v", err)
	}
	logger := logrus.New()
	logger.Level = logLevel

	// read config
	config_file := c.String("config-file")
	if config_file == "" {
		logger.Fatalf("Error missing flag 'config-file'")
	}
	cnf, err := ants.Read(config_file)
	if err != nil {
		logger.Fatalf("Error reading '%s': %s", config_file, err)
	}

	// read command
	cmd := strings.Join(c.Args(), " ")
	if cmd == "" {
		logger.Fatalf("Missing a command")
	}

	// create ant
	ant, err := ants.Create(cnf, logger)
	if err != nil {
		logger.Fatalf("Error creating ant: %s", err)
	}

	err = ant.Run(cmd)
	if err != nil {
		logger.Fatalf("Error running ant: %s", err)
	}
}

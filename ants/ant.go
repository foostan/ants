package ants

import (
	"bytes"
	"code.google.com/p/go.crypto/ssh"
	"strconv"

	log "github.com/Sirupsen/logrus"
)

type Ant struct {
	config *Config
	Logger *log.Logger
}

func Create(config *Config, logger *log.Logger) (*Ant, error) {
	return &Ant{
		config: config,
		Logger: logger,
	}, nil
}

func (a *Ant) Run(cmd string) error {
	a.Logger.Info("Running ant")
	a.Ssh(cmd)

	return nil
}

func (a *Ant) Ssh(cmd string) {
	for _, node := range a.config.Nodes {
		config := &ssh.ClientConfig{
			User: node.SshUser,
			Auth: []ssh.AuthMethod{
				ssh.Password(node.SshPassword),
			},
		}

		addr := node.Ip + ":" + strconv.Itoa(node.SshPort)
		a.Logger.Infof("Connect to %s", addr)
		conn, err := ssh.Dial("tcp", addr, config)
		if err != nil {
			a.Logger.Fatalf("Unable to connect: %s", err)
		}
		defer conn.Close()

		session, err := conn.NewSession()
		if err != nil {
			a.Logger.Fatalf("Unable to create session: %s", err)
		}
		defer session.Close()

		var stdoutBuf bytes.Buffer
		session.Stdout = &stdoutBuf
		err = session.Run(cmd)
		if err != nil {
			a.Logger.Fatalf("Execute command is failure: %s", err)
		}

		a.Logger.Info(stdoutBuf.String())
	}
}

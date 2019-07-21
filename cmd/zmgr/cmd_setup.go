package main

import (
	"errors"

	"github.com/Urethramancer/signor/log"
	"github.com/Urethramancer/signor/opt"
)

type CmdSetup struct {
	opt.DefaultHelp
	Host     string `placeholder:"HOST" help:"Domain of the Zimbra LDAP server to configure for."`
	Port     string `placeholder:"PORT" help:"Port for the LDAP server. Use 389 if unsure."`
	Password string `placeholder:"PASSWORD" help:"Password for the LDAP root user (use Zimbra's zmldappasswd tool to set it)."`
}

func (cmd *CmdSetup) Run(in []string) error {
	if cmd.Help || cmd.Password == "" {
		return errors.New(opt.ErrorUsage)
	}

	cfg := Config{
		Host:     cmd.Host,
		Port:     cmd.Port,
		Password: cmd.Password,
	}

	err := saveConfig(&cfg)
	if err != nil {
		return err
	}

	log.Default.Msg("Settings saved to %s", configName)
	return nil
}

package main

import (
	"errors"

	"github.com/Urethramancer/signor/log"
	"github.com/Urethramancer/signor/opt"
	"github.com/Urethramancer/zimbra"
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

	zc, err := zimbra.Connect(cfg.Host, cfg.Port, cfg.Password, "")
	if err != nil {
		return err
	}

	defer zc.Close()
	m := log.Default.Msg
	m("Login OK.")

	p, err := zc.GetLMTPPort()
	if err != nil {
		return err
	}

	cfg.LMTPPort = p
	err = saveConfig(&cfg)
	if err != nil {
		return err
	}

	m("Settings saved to %s", configName)
	return nil
}

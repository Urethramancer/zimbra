package main

import (
	"errors"

	"github.com/Urethramancer/signor/log"
	"github.com/Urethramancer/signor/opt"
	"github.com/Urethramancer/zimbra"
)

// CmdUser subcommands.
type CmdUser struct {
	opt.DefaultHelp
	List CmdUserList `command:"list" aliases:"ls" help:"List user accounts."`
	Add  CmdUserAdd  `command:"add" help:"Add a new user account."`
	Del  CmdUserDel  `command:"delete" aliases:"del,rm" help:"Delete a user account."`
}

func (cmd *CmdUser) Run(in []string) error {
	return errors.New(opt.ErrorUsage)
}

/*
 * Listing users.
 */

// CmdUserList subcommands.
type CmdUserList struct {
	opt.DefaultHelp
	Domain string `placeholder:"DOMAIN" help:"Optional domain to limit search."`
}

// Run the list command.
func (cmd *CmdUserList) Run(in []string) error {
	if cmd.Help {
		return errors.New(opt.ErrorUsage)
	}

	cfg := loadConfig()
	zc, err := zimbra.Connect(cfg.Host, cfg.Port, cfg.Password, cfg.LMTPPort)
	m := log.Default.Msg
	if err != nil {
		m("Login failed. (%s)", err.Error())
		return nil
	}

	defer zc.Close()
	var list []string
	if cmd.Domain == "" {
		list, err = zc.GetUsers()
	} else {
		list, err = zc.GetUsersInDomain(cmd.Domain)
	}
	if err != nil {
		return err
	}

	for _, x := range list {
		m("%s", x)
	}
	return nil
}

/*
 * Adding users.
 */

// CmdUserAdd options.
type CmdUserAdd struct {
	opt.DefaultHelp
	Email     string `placeholder:"EMAIL" help:"E-mail address to create."`
	GivenName string `placeholder:"GIVENNAME" help:"Optional first name of user."`
	SurName   string `placeholder:"SURNAME" help:"Optional last name of user."`
}

// Run the add command.
func (cmd *CmdUserAdd) Run(in []string) error {
	if cmd.Help || cmd.Email == "" {
		return errors.New(opt.ErrorUsage)
	}

	cfg := loadConfig()
	zc, err := zimbra.Connect(cfg.Host, cfg.Port, cfg.Password, cfg.LMTPPort)
	m := log.Default.Msg
	if err != nil {
		m("Login failed. (%s)", err.Error())
		return nil
	}

	defer zc.Close()
	pw, err := zc.AddUser(cmd.Email, cmd.GivenName, cmd.SurName)
	if err != nil {
		return err
	}

	m("Password: %s", pw)
	return nil
}

// CmdUserDel options.
type CmdUserDel struct {
	opt.DefaultHelp
	Email string `placeholder:"EMAIL" help:"E-mail address to delete."`
}

// Run the delete command.
func (cmd *CmdUserDel) Run(in []string) error {
	if cmd.Help || cmd.Email == "" {
		return errors.New(opt.ErrorUsage)
	}

	cfg := loadConfig()
	zc, err := zimbra.Connect(cfg.Host, cfg.Port, cfg.Password, cfg.LMTPPort)
	m := log.Default.Msg
	if err != nil {
		m("Login failed. (%s)", err.Error())
		return nil
	}

	defer zc.Close()
	return zc.DelUser(cmd.Email)
}

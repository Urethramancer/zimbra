package zimbra

import (
	"fmt"
	"strings"

	"github.com/Urethramancer/signor/uuid"
	"gopkg.in/ldap.v3"
)

func (zc *ZimbraLDAP) getAccounts(scope string) ([]string, error) {
	req := ldap.NewSearchRequest(scope, ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		"(&(objectClass=zimbraAccount))", []string{"dn", "cn"}, nil)
	res, err := zc.conn.Search(req)
	if err != nil {
		return nil, err
	}

	list := []string{}
	for _, e := range res.Entries {
		list = append(list, UserToEmail(e.DN))
	}

	return list, nil
}

// GetUsers returns a list of all users.
func (zc *ZimbraLDAP) GetUsers() ([]string, error) {
	return zc.getAccounts("")
}

// GetUsers returns a list of all users in a domain.
func (zc *ZimbraLDAP) GetUsersInDomain(domain string) ([]string, error) {
	d, err := NewDomain(domain)
	if err != nil {
		return nil, err
	}

	return zc.getAccounts(d.BindDN)
}

// AddUser with one e-mail address and return the password.
func (zc *ZimbraLDAP) AddUser(email, gn, sn string) (string, error) {
	u, err := NewUser(email)
	if err != nil {
		return "", err
	}

	req := ldap.NewAddRequest(u.BindDN, nil)
	req.Attribute("uid", []string{u.Name})
	req.Attribute("mail", []string{email})
	if gn != "" {
		req.Attribute("givenName", []string{gn})
	}
	if sn != "" {
		req.Attribute("sn", []string{sn})
		n := fmt.Sprintf("%s %s", gn, sn)
		req.Attribute("cn", []string{n})
		req.Attribute("displayName", []string{n})
	}
	req.Attribute("objectClass", []string{"zimbraAccount", "amavisAccount", "inetOrgPerson"})
	req.Attribute("zimbraAccountStatus", []string{"active"})
	req.Attribute("zimbraId", []string{uuid.NewGenerator().Generate()})
	a := strings.Split(zc.Address, ":")
	tr := fmt.Sprintf("lmtp:%s:7025", a[0])
	req.Attribute("zimbraMailTransport", []string{tr})
	req.Attribute("zimbraMailHost", []string{a[0]})
	err = zc.conn.Add(req)
	if err != nil {
		return "", err
	}

	pw := GenString(16)
	pwreq := ldap.NewPasswordModifyRequest(u.BindDN, "", pw)
	_, err = zc.conn.PasswordModify(pwreq)
	return pw, err
}

// DelUser based on e-mail address.
func (zc *ZimbraLDAP) DelUser(email string) error {
	u, err := NewUser(email)
	if err != nil {
		return err
	}

	req := ldap.NewDelRequest(u.BindDN, nil)
	return zc.conn.Del(req)
}

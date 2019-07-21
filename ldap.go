package zimbra

import (
	"crypto/tls"
	"net"

	"gopkg.in/ldap.v2"
)

// ZimbraLDAP holds an active connection to a Zimbra LDAP server.
type ZimbraLDAP struct {
	// User with administrator privileges.
	User *User
	// Address string from host and port.
	Address string

	conn *ldap.Conn
}

// Connect sets up a secure TCP+TLS connection to the LDAP server and tries to authenticate as the admin user.
// If successful, a ZimbraLDAP struct is returned.
func Connect(host, port, username, password string) (*ZimbraLDAP, error) {
	var err error
	u, err := NewUser(username)
	if err != nil {
		return nil, err
	}

	zc := ZimbraLDAP{
		User:    u,
		Address: net.JoinHostPort(host, port),
	}

	zc.conn, err = ldap.Dial("tcp", zc.Address)
	if err != nil {
		return nil, err
	}

	cfg := &tls.Config{
		InsecureSkipVerify: false,
		ServerName:         host,
	}
	err = zc.conn.StartTLS(cfg)
	if err != nil {
		return nil, err
	}

	err = zc.bind(password)
	if err != nil {
		return nil, err
	}

	u.conn = &zc
	return &zc, nil
}

func (zc *ZimbraLDAP) bind(password string) error {
	return zc.conn.Bind(zc.User.BindDN, password)
}

func (zc *ZimbraLDAP) Close() {
	zc.conn.Close()
}

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

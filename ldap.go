package zimbra

import (
	"crypto/tls"
	"net"

	"gopkg.in/ldap.v3"
)

// ZimbraLDAP holds an active connection to a Zimbra LDAP server.
type ZimbraLDAP struct {
	// Address string from host and port.
	Address string

	conn *ldap.Conn
}

// Connect sets up a secure TCP+TLS connection to the LDAP server and tries to authenticate as the admin user.
// If successful, a ZimbraLDAP struct is returned.
func Connect(host, port, password string) (*ZimbraLDAP, error) {
	zc := ZimbraLDAP{
		Address: net.JoinHostPort(host, port),
	}

	var err error
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

	return &zc, nil
}

func (zc *ZimbraLDAP) bind(password string) error {
	return zc.conn.Bind("cn=config", password)
}

func (zc *ZimbraLDAP) Close() {
	zc.conn.Close()
}

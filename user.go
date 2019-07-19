package zimbra

import (
	"strings"

	"github.com/Urethramancer/signor/stringer"
)

// User structure for easy handling with LDAP.
type User struct {
	// Name is the uid portion of the DN.
	Name string
	// Domain contains the remaining (DC) parts with all subdomains and the TLD.
	Domain *Domain
	// BindDN is used to authenticate this user.
	BindDN string

	conn *ZimbraLDAP
}

// NewUser splits an e-mail account string into LDAP-friendly parts.
func NewUser(name string) (*User, error) {
	u := User{}
	a := strings.Split(name, "@")
	if len(a) < 2 {
		return nil, ErrInvalidEmail
	}

	var err error
	u.Name = a[0]
	u.Domain, err = NewDomain(a[1])
	if err != nil {
		return nil, err
	}

	s := stringer.New()
	s.WriteStrings("uid=", u.Name, ",ou=people,", u.Domain.BindDN)
	u.BindDN = s.String()
	return &u, nil
}

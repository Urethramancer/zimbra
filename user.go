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

// UserToEmail converts LDAP result format to an e-mail address.
func UserToEmail(user string) string {
	a := strings.Split(user, ",")
	s := stringer.New()
	dot := false
	for _, x := range a {
		t := strings.Split(x, "=")
		if len(t) != 2 {
			continue
		}
		switch t[0] {
		case "uid":
			s.WriteStrings(t[1], "@")
		case "dc":
			if dot {
				s.WriteString(".")
			} else {
				dot = true
			}
			s.WriteString(t[1])
		}
	}
	return s.String()
}

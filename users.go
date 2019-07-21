package zimbra

import "gopkg.in/ldap.v3"

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

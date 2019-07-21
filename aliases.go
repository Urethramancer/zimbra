package zimbra

import "gopkg.in/ldap.v3"

func (zc *ZimbraLDAP) GetAliases(email string) ([]string, error) {
	u, err := NewUser(email)
	if err != nil {
		return nil, err
	}

	req := ldap.NewSearchRequest(u.Domain.BindDN, ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		"(&(objectClass=zimbraAccount)(uid="+u.Name+"))", []string{"mail"}, nil)
	res, err := zc.conn.Search(req)
	if err != nil {
		return nil, err
	}

	var list []string
	for _, e := range res.Entries {
		x := e.GetAttributeValues("mail")
		list = append(list, x...)
	}
	return list, nil
}

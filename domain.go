package zimbra

import (
	"strings"

	"github.com/Urethramancer/signor/stringer"
)

type Domain struct {
	Parts  []string
	BindDN string
}

func NewDomain(domain string) (*Domain, error) {
	d := Domain{}
	d.Parts = strings.Split(domain, ".")
	if len(d.Parts) == 0 {
		return nil, ErrInvalidDomain
	}

	s := stringer.New()
	s.WriteStrings("dc=", d.Parts[0])
	for i := 1; i < len(d.Parts); i++ {
		s.WriteStrings(",dc=", d.Parts[i])
	}
	d.BindDN = s.String()
	return &d, nil
}

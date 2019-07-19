package zimbra

import "errors"

var (
	ErrInvalidEmail  = errors.New("invalid e-mail address")
	ErrInvalidDomain = errors.New("invalid domain")
)

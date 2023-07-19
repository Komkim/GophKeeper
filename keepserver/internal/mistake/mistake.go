package mistake

import "errors"

var (
	ErrNotAuthenticated = errors.New("user not authenticated")
	ErrLoginIsTaken     = errors.New("login is taken")
	ErrDBID             = errors.New("uuid is not correct")
)

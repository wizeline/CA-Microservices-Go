package service

func validateUserCreate(u UserCreateArgs) error {
	if u == (UserCreateArgs{}) {
		return ErrEmptyArgs
	}
	if u.FirstName == "" {
		return &InvalidInputErr{Field: "FirstName", Err: ErrEmptyValue}
	}
	if u.LastName == "" {
		return &InvalidInputErr{Field: "LastName", Err: ErrEmptyValue}
	}
	if err := validateEmail(u.Email); err != nil {
		return &InvalidInputErr{Field: "Email", Err: err}
	}
	if u.Username == "" {
		return &InvalidInputErr{Field: "Username", Err: ErrEmptyValue}
	}
	if err := validateUserPasswd(u.Passwd); err != nil {
		return &InvalidInputErr{Field: "Passwd", Err: err}
	}
	return nil
}

func validateUserUpdate(user UserUpdateArgs) error {
	if user == (UserUpdateArgs{}) {
		return ErrEmptyArgs
	}
	if user.ID == 0 {
		return &InvalidInputErr{Field: "ID", Err: ErrZeroValue}
	}
	return nil
}

func validateUserPasswd(pwd string) error {
	if pwd == "" || len(pwd) < 6 {
		return ErrInvalidPasswd
	}
	return nil
}

func validateUserFilter(filter string) error {
	if filter == "" {
		return &InvalidInputErr{Field: "filter", Err: ErrEmptyValue}
	}
	switch filter {
	case "FirstName", "LastName", "Email", "Username":
		return nil
	default:
		return &InvalidFilterErr{Filter: filter, Err: ErrNotSupported}
	}
}

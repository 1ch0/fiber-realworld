package bcode

var (
	// ErrUnsupportedEmailModification is the error of unsupported email modification
	ErrUnsupportedEmailModification = NewBcode(20001, "the user already has an email address and cannot modify it again")
	// ErrUserAlreadyDisabled is the error of user already disabled
	ErrUserAlreadyDisabled = NewBcode(20002, "the user is already disabled")
	// ErrUserAlreadyEnabled is the error of user already enabled
	ErrUserAlreadyEnabled = NewBcode(20003, "the user is already enabled")
	// ErrUserCannotModified is the error of user cannot modified
	ErrUserCannotModified = NewBcode(20004, "the user cannot be modified in dex login mode")
	// ErrUserInvalidPassword is the error of user invalid password
	ErrUserInvalidPassword = NewBcode(20005, "the password is invalid")
	// ErrUserInconsistentPassword is the error of user inconsistent password
	ErrUserInconsistentPassword = NewBcode(20007, "the password is inconsistent with the user")
	// ErrUsernameNotExist is the error of username not exist
	ErrUsernameNotExist = NewBcode(20008, "the username is not exist")
	// ErrDexNotFound is the error of dex not found
	// ErrEmptyAdminEmail is the error of empty admin email
	ErrEmptyAdminEmail = NewBcode(20010, "the admin email is empty, please set the admin email before using sso login")
	// ErrNoAdminUser is the error of no admin user
	ErrNoAdminUser = NewBcode(20011, "the admin user is not found, please init the platform first")
)

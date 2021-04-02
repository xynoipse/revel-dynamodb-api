package user

const (
	Entity = "USER"

	// Password Config
	PasswordTime    = 1
	PasswordMemory  = 64 * 1024
	PasswordThreads = 4
	PasswordKeyLen  = 32

	// Services
	CreateService int = iota

	// Responses
	RegisterSuccess    = "Registered successfully"
	LogInSuccess       = "Logged in successfully"
	EmailAlreadyTaken  = "The email has already been taken."
	InvalidCredentials = "These credentials do not match our records."
)

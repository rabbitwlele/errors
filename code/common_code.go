package code

var (
	OK         = New(0, "success")
	ServerErr  = New(500, "server error")
	UnknownErr = New(900, "unknown error")
)

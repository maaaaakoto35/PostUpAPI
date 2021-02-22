package controllers

// Context this interface is connecting to echo.Context.
type Context interface {
	Param(string) string
	Bind(interface{}) error
	JSON(int, interface{}) error
}

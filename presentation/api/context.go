package api

type Context interface {
	Param(param string) string
	JSON(status int, data interface{})
}

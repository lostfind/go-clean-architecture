package api

type Context interface {
	Param(string) string
	JSON(int, interface{})
}

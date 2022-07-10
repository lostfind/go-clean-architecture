package server

type Server interface {
	Run(addr ...string) (err error)
}

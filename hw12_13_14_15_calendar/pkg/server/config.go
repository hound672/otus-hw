package server

type Config struct {
	UseReflection bool

	PortGRPC int
	PortHTTP int

	HTTPReadTimeout       int
	HTTPWriteTimeout      int
	HTTPIdleTimeout       int
	HTTPReadHeaderTimeout int
}

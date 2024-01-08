package server

type Config struct {
	UseReflection bool `mapstructure:"use_reflection"`

	PortGRPC int `mapstructure:"port_grpc"`
	PortHTTP int `mapstructure:"port_http"`

	HTTPReadTimeout       int `mapstructure:"http_read_timeout"`
	HTTPWriteTimeout      int `mapstructure:"http_write_timeout"`
	HTTPIdleTimeout       int `mapstructure:"http_idle_timeout"`
	HTTPReadHeaderTimeout int `mapstructure:"http_read_header_timeout"`
}

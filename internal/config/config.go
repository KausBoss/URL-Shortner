package config

type Server struct {
	Address string
}
type Config struct {
	Server Server
}

func Load() *Config {
	cfg := &Config{
		Server: Server{Address: ":8080"},
	}

	//cfg.Server.Address = os.Getenv("SERVER_ADDRESS")
	return cfg
}

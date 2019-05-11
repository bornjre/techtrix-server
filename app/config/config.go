package config

var conf *Config

type Config struct {
	Dev        bool
	DBpath     string
	StatDBpath string
}

func init() {
	// todo get that from arg/Env/config file
	conf = &Config{
		Dev:    true,
		DBpath: "mybolt.db",
	}
}

func GetConfig() *Config {
	return conf
}

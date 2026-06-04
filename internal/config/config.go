package config

type Config struct {
	Config ServerConfig `json:"config"`
	Gates  []Gate       `json:"gates"`
}

type ServerConfig struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

type Gate struct {
	Name     string `json:"name"`
	Mnemonic string `json:"mnemonic"`
}
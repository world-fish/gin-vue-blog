package config

type System struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
	Evn  string `yaml:"evn"`
}

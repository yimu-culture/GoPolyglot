package configs

type Config struct {
	Server    Server    `yaml:"server"`
	Log       Log       `yaml:"log"`
	Databases Databases `yaml:"databases"`
}
type Server struct {
	Env     string `yaml:"env"`
	AppName string `yaml:"appName"`
	Mode    string `yaml:"mode"`
	Address string `yaml:"address"`
	Port    string `yaml:"port"`
}
type Log struct {
	AccessKeyID     string `yaml:"accessKeyID"`
	AccessKeySecret string `yaml:"accessKeySecret"`
	LogStore        string `yaml:"logStore"`
	Endpoint        string `yaml:"endpoint"`
	Project         string `yaml:"project"`
	LookAddr        string `yaml:"lookAddr"`
	Storage         int    `yaml:"storage"`
}
type Redis struct {
	Address      string `yaml:"address"`
	Port         int    `yaml:"port"`
	Db           int    `yaml:"db"`
	Username     string `yaml:"username"`
	Password     string `yaml:"password"`
	PoolSize     int    `yaml:"poolSize"`
	MinIdleConns int    `yaml:"minIdleConns"`
	DialTimeout  int    `yaml:"dialTimeout"`
	Asname       string `yaml:"asname"`
}
type Mysql struct {
	Username     string `yaml:"username"`
	Password     string `yaml:"password"`
	Address      string `yaml:"address"`
	Dbname       string `yaml:"dbname"`
	Asname       string `yaml:"asname"`
	MaxOpenConns int    `yaml:"maxOpenConns"`
}
type Databases struct {
	Redis []Redis `yaml:"redis"`
	Mysql []Mysql `yaml:"mysql"`
}

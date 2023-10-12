package config

type NacosConfig struct {
	Host      string `mapstructure:"host"`
	Port      uint64 `mapstructure:"port"`
	Namespace string `mapstructure:"namespace"`
	User      string `mapstructure:"user"`
	Password  string `mapstructure:"password"`
	DataId    string `mapstructure:"dataId"`
	Group     string `mapstructure:"group"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host" json:"host"`
	Port     int    `mapstructure:"port" json:"port"`
	Password string `mapstructure:"password" json:"password"`
	DB       int    `mapstructure:"db" json:"db"`
	PoolSize int    `mapstructure:"db" json:"pool_size"`
}

type ConsulConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
}

type MySQLConfig struct {
	Host     string `mapstructure:"host" json:"host"`
	Port     int    `mapstructure:"port" json:"port"`
	User     string `mapstructure:"user" json:"user"`
	Password string `mapstructure:"password" json:"password"`
	Dbname   string `mapstructure:"dbname" json:"dbname"`
}

type DouyinConfig struct {
	Name       string       `mapstructure:"name" json:"name"`
	Host       string       `mapstructure:"host" json:"host"`
	Tags       []string     `mapstructure:"tags" json:"tags"`
	Port       int          `mapstructure:"port" json:"port"`
	Rdb        RedisConfig  `json:"redis" mapstructure:"redis"`
	Db         MySQLConfig  `json:"mysql" mapstructure:"mysql"`
	ConsulInfo ConsulConfig `json:"consul" mapstructure:"consul"`
}

package conf

func init() {
	RuntimeConf = new(Conf)
}

type Conf struct {
	MysqlConf `yaml:"mysql"`
}

type MysqlConf struct {
	Host        string `yaml:"host"`
	Port        int    `yaml:"port"`
	Db          string `yaml:"db"`
	Username    string `yaml:"username"`
	Password    string `yaml:"password"`
	MaxOpenConn int    `yaml:"max_open_conn"`
	MaxIdleConn int    `yaml:"max_idle_conn"`
}

var RuntimeConf *Conf

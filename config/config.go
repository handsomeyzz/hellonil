package config

const (
	JwtExpireTime = 24          //jwt过期时间
	JwtSecretID   = "hello-nil" //jwt加密ID
)

type Redis struct {
	Addr     string
	Password string
	Port     int
}

type Mysql struct {
	Addr     string
	Password string
	User     string
	Port     int
	DB       string
}

type Cos struct {
	Addr      string
	Secredid  string
	Secredkey string
}

type LogConfig struct {
	Level      string
	Filename   string
	MaxSize    int
	MaxAge     int
	MaxBackups int
}

// redis配置文件
func RedisX() Redis {
	return Redis{
		Addr:     "", //ip地址
		Password: "",      //密码
		Port:     ,            //端口号
	}
}

// mysql配置文件
func MysqlX() Mysql {
	return Mysql{
		Addr:     "", //ip地址
		Port:     0,                                //端口号
		User:     "root",                               //用户
		Password: "",                      //密码
		DB:       "",                           //数据库名字
	}
}

// cos存储配置文件
func CosX() Cos {
	return Cos{
		Addr:      "", //ip地址
		Secredid:  "",               //私密id
		Secredkey: "",                   //私密key

	}
}

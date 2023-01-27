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
		Addr:     "101.35.43.151", //ip地址
		Password: "hellonil",      //密码
		Port:     6379,            //端口号
	}
}

// mysql配置文件
func MysqlX() Mysql {
	return Mysql{
		Addr:     "gz-cdb-0r0kvcx1.sql.tencentcdb.com", //ip地址
		Port:     63930,                                //端口号
		User:     "root",                               //用户
		Password: "HELLOnil123++",                      //密码
		DB:       "hellonil",                           //数据库名字
	}
}

// cos存储配置文件
func CosX() Cos {
	return Cos{
		Addr:      "https://nil-1316622710.cos.ap-nanjing.myqcloud.com", //ip地址
		Secredid:  "AKIDgATYN5WsPiy71oQovsd7PVO1QessdCcr",               //私密id
		Secredkey: "1oPvHalKgmRu9AwUs0wgDUVGn58yc5MQ",                   //私密key

	}
}

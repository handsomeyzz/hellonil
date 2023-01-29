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

// redis配置文件
func RedisX() Redis {
	return Redis{
		Addr:     "",   //ip地址
		Password: "",   //密码
		Port:     6379, //端口号
	}
}

// mysql配置文件
func MysqlX() Mysql {
	return Mysql{
		Addr:     "",     //ip地址
		Port:     63930,  //端口号
		User:     "root", //用户
		Password: "",     //密码
		DB:       "",     //数据库名字
	}
}

// Cos存储配置，txt文件下，自己填一下
var CosUrl = ""
var CosSecretID = ""
var CosSecretkey = ""

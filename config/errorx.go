package config

const (
	JwtInvalid = iota
	JwtExpire

	UserNoExist
	PwError
)

var codeTag = map[int]string{
	JwtInvalid:  "invalid token",
	JwtExpire:   "jwt expired",
	UserNoExist: "user No exist",
	PwError:     "password  error",
}

func Text(code int) string {
	return codeTag[code]
}

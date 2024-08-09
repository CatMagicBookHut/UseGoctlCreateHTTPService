package errmsg

const (
	SUCCESS = 200
	ERROR   = 500
)

var codeMsg = map[int]string{
	SUCCESS: "ok",
	ERROR:   "error",
}

func GetErrMsg(code int) string {
	return codeMsg[code]
}

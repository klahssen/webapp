package context

type contextItem string

//constants
const (
	JwtToken contextItem = "jwttoken"
	ReqTime  contextItem = "reqtime"
)

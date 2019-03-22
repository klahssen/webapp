package http

type contextItem string

//constants
const (
	JwtTokenInCtx contextItem = "jwttoken"
	ReqTimeInCtx  contextItem = "reqtime"
)

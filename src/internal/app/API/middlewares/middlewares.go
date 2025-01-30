package middlewares

type IMiddleware interface{}

type Middleware struct {
	IMiddleware
}

func NewMiddleware() *Middleware {
	return &Middleware{}
}

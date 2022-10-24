package router

type RouterGroup struct {
	BaseRouter
	InitRouter
	UserRouter
	ApiRouter
}

var RouterGroupApp = new(RouterGroup)

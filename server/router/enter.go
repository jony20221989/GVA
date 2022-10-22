package router

type RouterGroup struct {
	BaseRouter
	InitRouter
	UserRouter
}

var RouterGroupApp = new(RouterGroup)

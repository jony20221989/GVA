package router

type RouterGroup struct {
	BaseRouter
	InitRouter
}

var RouterGroupApp = new(RouterGroup)

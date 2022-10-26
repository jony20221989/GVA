package router

type RouterGroup struct {
	BaseRouter
	InitRouter
	UserRouter
	ApiRouter
	MenuRouter
}

var RouterGroupApp = new(RouterGroup)

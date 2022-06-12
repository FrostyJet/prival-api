package http

type Router interface {
	InitRoutes()
	Serve()
}

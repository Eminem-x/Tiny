package router

import (
	"example/handler"
	"tinyGin"
)

func RegisterRouter(r *tinyGin.Engine) {
	// 1. simply http handler
	r.GET("/ping", handler.Ping)

	// 2. context based on req and resp
	r.GET("/context", handler.Context)

	// 3. router
	r.GET("/login", handler.Login)

	// 4. group router
	{
		{
			admin := r.Group("/admin")
			admin.GET("/ping", handler.Ping)
		}

		{
			user := r.Group("/user")
			user.GET("/ping", handler.Ping)
		}
	}

	// 5. middleware
	r.Use(handler.Logger)

	// 6. template

	// 7. recover
	r.GET("/panic", handler.Panic)
}

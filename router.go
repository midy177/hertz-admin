/*
 * Copyright 2024 HertzAdmin Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 */

package main

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"hertz-admin/biz/handler/middleware"
	"hertz-admin/configs"
	"hertz-admin/data"
)

// customizeRegister registers customize routers.
func customizedRegister(r *server.Hertz) {
	// your code ...
	// login and refresh_token power by jwt Auth middleware
	r.POST("/api/login", middleware.GetJWTMiddleware(configs.Data(), data.Default(), data.CasbinEnforcer()).LoginHandler)
	r.POST("/api/logout", middleware.GetJWTMiddleware(configs.Data(), data.Default(), data.CasbinEnforcer()).LogoutHandler)
	r.POST("/api/refresh_token", middleware.GetJWTMiddleware(configs.Data(), data.Default(), data.CasbinEnforcer()).RefreshHandler)
}

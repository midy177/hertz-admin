/*
 * Copyright 2024 HertzAdmin Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 */

package main

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/cors"
	_ "github.com/mattn/go-sqlite3"
	"hertz-admin/configs"
	"hertz-admin/data"
	"time"

	"github.com/cloudwego/hertz/pkg/app/server"
)

func main() {
	// init
	data.InitDataConfig()

	// start server
	c := configs.Data()
	h := server.Default(
		server.WithHostPorts(fmt.Sprintf("%s:%d", c.Host, c.Port)))
	//h.LoadHTMLGlob("assets/index.html")
	h.StaticFS("/", &app.FS{
		Root:                 "./assets",
		IndexNames:           []string{"index.html"},
		Compress:             true,
		CompressedFileSuffix: ".hertz",
		AcceptByteRange:      true,
		PathNotFound: func(c context.Context, ctx *app.RequestContext) {
			app.ServeFile(ctx, "./assets/index.html")
		},
	})
	h.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		MaxAge: 12 * time.Hour,
	}))
	register(h)
	h.Spin()
}

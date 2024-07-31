/*
 * Copyright 2023 FormulaGo Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 */

package main

import (
	"context"
	"fmt"
	"formulago/configs"
	"formulago/data"
	"github.com/cloudwego/hertz/pkg/app"
	_ "github.com/mattn/go-sqlite3"

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

	register(h)
	h.Spin()
}

/*
 * Copyright 2024 HertzAdmin Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 */

package data

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	entAdapter "github.com/casbin/ent-adapter"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"hertz-admin/configs"
)

var casbinEnforcer *casbin.Enforcer

func initCasbin() {
	var err error
	casbinEnforcer, err = newCasbin(configs.Data())
	if err != nil {
		hlog.Fatal(err)
	}
}

// CasbinEnforcer Get a default casbin enforcer instance
func CasbinEnforcer() *casbin.Enforcer {
	return casbinEnforcer
}

func newCasbin(config configs.Config) (enforcer *casbin.Enforcer, err error) {
	var dsnStr = "./admin.db"
	switch config.Database.Type {
	case "mysql", "postgres":
		dsnStr = fmt.Sprintf(mySqlDsn,
			config.Database.Username, config.Database.Password, config.Database.Host, config.Database.Port, config.Database.DBName)
	case "sqlite3":
		if config.Database.Host == "" {
			config.Database.Host = "./admin.db"
		}
		dsnStr = fmt.Sprintf("%s%s", config.Database.Host, "?cache=shared&mode=rwc&_fk=1")
	default:
		hlog.Error(fmt.Errorf("unsupported database type: %s", config.Database.Type))
		return
	}
	adapter, err := entAdapter.NewAdapter(config.Database.Type, dsnStr)
	if err != nil {
		hlog.Error(err)
		return
	}
	//adapter, err := entAdapter.NewAdapterWithClient(data.DBClient)

	var text string
	if config.Casbin.ModelText == "" {
		text = `
		[request_definition]
		r = sub, obj, act
		
		[policy_definition]
		p = sub, obj, act
		
		[role_definition]
		g = _, _
		
		[policy_effect]
		e = some(where (p.eft == allow))
		
		[matchers]
		m = r.sub == p.sub && keyMatch2(r.obj,p.obj) && r.act == p.act
		`
	} else {
		text = config.Casbin.ModelText
	}

	m, err := model.NewModelFromString(text)
	if err != nil {
		hlog.Error(err)
		return
	}

	enforcer, err = casbin.NewEnforcer(m, adapter)
	if err != nil {
		hlog.Error(err)
		return
	}

	err = enforcer.LoadPolicy()
	if err != nil {
		hlog.Error(err)
		return
	}

	return
}

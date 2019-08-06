// Copyright 2017 Emir Ribic. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

// go_blog - Go(lang) restful starter kit
//
// API Docs for go_blog v1
//
// 	 Terms Of Service:  N/A
//     Schemes: http
//     Version: 2.0.0
//     License: MIT http://opensource.org/licenses/MIT
//     Contact: Emir Ribic <vasarostik@gmail.com> https://vasarostik.ba
//     Host: localhost:8080
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Security:
//     - bearer: []
//
//     SecurityDefinitions:
//     bearer:
//          type: apiKey
//          name: Authorization
//          in: header
//
// swagger:meta
package api

import (
	"crypto/sha1"

	"github.com/vasarostik/go_blog/pkg/utl/zlog"

	"github.com/vasarostik/go_blog/pkg/api/auth"
	al "github.com/vasarostik/go_blog/pkg/api/auth/logging"
	at "github.com/vasarostik/go_blog/pkg/api/auth/transport"
	"github.com/vasarostik/go_blog/pkg/api/password"
	pl "github.com/vasarostik/go_blog/pkg/api/password/logging"
	pt "github.com/vasarostik/go_blog/pkg/api/password/transport"
	"github.com/vasarostik/go_blog/pkg/api/user"
	ul "github.com/vasarostik/go_blog/pkg/api/user/logging"
	ut "github.com/vasarostik/go_blog/pkg/api/user/transport"
	"github.com/vasarostik/go_blog/pkg/utl/config"
	"github.com/vasarostik/go_blog/pkg/utl/middleware/jwt"
	"github.com/vasarostik/go_blog/pkg/utl/postgres"
	"github.com/vasarostik/go_blog/pkg/utl/rbac"
	"github.com/vasarostik/go_blog/pkg/utl/secure"
	"github.com/vasarostik/go_blog/pkg/utl/server"
	"github.com/vasarostik/go_blog/pkg/api/post"
	psl "github.com/vasarostik/go_blog/pkg/api/post/logging"
	pst "github.com/vasarostik/go_blog/pkg/api/post/transport"
)

// Start starts the API service
func Start(cfg *config.Configuration) error {
	db, err := postgres.New(cfg.DB.PSN, cfg.DB.Timeout, cfg.DB.LogQueries)
	if err != nil {
		return err
	}

	sec := secure.New(cfg.App.MinPasswordStr, sha1.New())
	rbac := rbac.New()
	jwt := jwt.New(cfg.JWT.Secret, cfg.JWT.SigningAlgorithm, cfg.JWT.Duration)
	log := zlog.New()
	e := server.New()

	e.Static("/swaggerui", cfg.App.SwaggerUIPath)

	at.NewHTTP(al.New(auth.Initialize(db, jwt, sec, rbac), log), e, jwt.MWFunc())

	v1 := e.Group("/v1")
	v1.Use(jwt.MWFunc())

	ut.NewHTTP(ul.New(user.Initialize(db, rbac, sec), log), v1, e)
	pt.NewHTTP(pl.New(password.Initialize(db, rbac, sec), log), v1)
	pst.NewHTTP(psl.New(post.Initialize(db, rbac, sec), log), v1)

	server.Start(e, &server.Config{
		Port:                cfg.Server.Port,
		ReadTimeoutSeconds:  cfg.Server.ReadTimeout,
		WriteTimeoutSeconds: cfg.Server.WriteTimeout,
		Debug:               cfg.Server.Debug,
	})

	return nil
}

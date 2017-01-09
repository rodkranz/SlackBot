// Copyright 2017 Kranz. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package main

import (
	"net/http"
	"log"
	"os"
	"fmt"
	"io/ioutil"

	"gopkg.in/macaron.v1"

	"github.com/rodkranz/webhook/module/github"
	"github.com/rodkranz/webhook/module/hydrate"
	"github.com/rodkranz/webhook/module/slack"
	"github.com/rodkranz/webhook/module/conf"
)

func main() {
	m := macaron.Classic()
	m.Any("/", webHook)

	ServerAddress := fmt.Sprintf("%s:%d", conf.HTTPHost, conf.HTTPPort)
	log.Printf("Error to run server [%v]", ServerAddress)
	err := http.ListenAndServe(ServerAddress, m)
	if err != nil {
		log.Fatalf("Error to run server: %v", err)
		os.Exit(1)
	}
	os.Exit(0)
}

func webHook(ctx *macaron.Context) string {
	b, err := ctx.Req.Body().Bytes()
	if err != nil {
		fmt.Printf("Error: %v", err.Error())
		ctx.Error(http.StatusBadRequest, err.Error())
	}

	gitHubWebHook, err := github.NewPayload(b)
	if err != nil {
		fmt.Printf("Error: %v", err.Error())
		ctx.Error(http.StatusBadRequest, err.Error())
	}

	slackPayload := hydrate.NewSlackFromGitHub(gitHubWebHook)
	res, err := slack.Send(slackPayload)
	if err != nil {
		fmt.Printf("Error: %v", err.Error())
		ctx.Error(http.StatusBadRequest, err.Error())
	}

	b, err = ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Error: %v", err.Error())
		ctx.Error(http.StatusBadRequest, err.Error())
	}
	defer res.Body.Close()

	return string(b)
}

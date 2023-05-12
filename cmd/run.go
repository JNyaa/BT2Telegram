package cmd

/*
  Copyright 2023 Malonan & JNyaa

  Licensed under the Apache License, Version 2.0 (the "License");
  you may not use this file except in compliance with the License.
  You may obtain a copy of the License at

      http://www.apache.org/licenses/LICENSE-2.0

  Unless required by applicable law or agreed to in writing, software
  distributed under the License is distributed on an "AS IS" BASIS,
  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
  See the License for the specific language governing permissions and
  limitations under the License.
*/

import (
	"fmt"
	"time"

	tele "github.com/3JoB/telebot"
	tb "github.com/3JoB/telebot-utils/bot"
	tbmw "github.com/3JoB/telebot-utils/middleware"
	"github.com/3JoB/telebot/middleware"

	"btg/bot"
	"btg/lib/config"
	"btg/lib/log"
	"btg/smtp"
)

var F = config.Get()

func T() string {
	return "[Runtime/" + time.Now().UTC().Format("2006-01-02 15:04:05") + "]"
}

func Start() {
	go smtp.Start()
	fmt.Println(T() + " Ready to start.....")
	fmt.Println(T() + " Configure robot information...")

	t := tb.New().
		SetError(func(err error, ctx tele.Context) { log.Use().Error(err.Error()) }).
		SetKey(F.String("bot.key"))
	if F.String("bot.api_server") != "" {
		t = t.SetAPIServer(F.String("bot.api_server"))
	}

	fmt.Println(T() + " Registering...")
	b := t.CreateBot()
	b.RemoveWebhook()
	b.Middleware(middleware.Recover())
	// b.Middleware(middleware.AutoRespond())
	b.Middleware(tbmw.Logger(&tbmw.LogSettings{Path: "./log/", FileName: "sticker-log"}))
	bot.Handle(b.B)
	fmt.Println(T() + " The robot is running on @" + b.Me().Username)
	b.Start()
}

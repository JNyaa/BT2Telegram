package smtp

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
	"bytes"
	"errors"
	"io"
	"log"
	"net/mail"
	"time"

	"github.com/emersion/go-smtp"
	"github.com/golang-module/dongle"

	"btg/lib/channel"
	"btg/lib/config"
)

var c = config.Get()  // Config
var ch = channel.Do() // Channel

// The Backend implements SMTP server methods.
type Backend struct{}

func (bkd *Backend) NewSession(_ *smtp.Conn) (smtp.Session, error) {
	return &Session{}, nil
}

// A Session is returned after EHLO.
type Session struct{}

func (s *Session) AuthPlain(username, password string) error {
	if !c.Bool("smtp.auth.status") {
		return nil
	}
	if username != c.String("smtp.auth.user") || password != c.String("smtp.auth.pass") {
		return errors.New("invalid username or password")
	}
	return nil
}

func (s *Session) Mail(from string, opts *smtp.MailOptions) error {
	return nil
}

func (s *Session) Rcpt(to string) error {
	return nil
}

func (s *Session) Data(r io.Reader) error {
	b, err := io.ReadAll(r)
	if err != nil {
		return err
	}
	message, err := mail.ReadMessage(bytes.NewReader(b))
	if err != nil {
		return err
	}
	body, err := io.ReadAll(message.Body)
	if err != nil {
		return err
	}
	bd := dongle.Decode.FromBytes(body).ByBase64().ToString()
	ch <- bd
	return nil
}

func (s *Session) Reset() {}

func (s *Session) Logout() error {
	return nil
}

func Start() {
	s := smtp.NewServer(&Backend{})

	s.Addr = ":" + c.String("smtp.port")
	s.Domain = "0.0.0.0"
	s.ReadTimeout = 10 * time.Second
	s.WriteTimeout = 10 * time.Second
	s.MaxMessageBytes = 1024 * 1024
	s.MaxRecipients = 50
	s.AllowInsecureAuth = true

	log.Println("SMTP Service: Starting server at", s.Addr)
	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

package bot

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
	tele "github.com/3JoB/telebot"

	"btg/lib/channel"
)

var ch = channel.Do()
var ech = channel.Err()

func ListenSMTP(t *tele.Bot) {
	for {
		data := <-ch
		_, err := t.Send(tele.ChatID(Chats), data)
		ech <- err
	}
}

func Handle(t *tele.Bot) {
	t.Handle("/my", CommandMy)
	go ListenSMTP(t)
}

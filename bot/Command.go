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
	"fmt"

	tele "github.com/3JoB/telebot"
)

func CommandMy(c tele.Context) error {
	msg := fmt.Sprintf("**UserID**: `%v`\n**UserName**: `%v`\n**ChatID**: `%v`\n**ChatType**: `%v`", c.Sender().ID, c.Sender().Username, c.Chat().ID, c.Chat().Type)
	c.Reply(msg, &tele.SendOptions{ParseMode: tele.ModeMarkdownV2})
	return nil
}

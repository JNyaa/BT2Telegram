package channel

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

type ch chan string

type ech chan error

// 消息传出
var csh ch

// 错误回传
var errf ech

func init() {
	csh = make(chan string)
	errf = make(chan error)
}

func Do() ch {
	return csh
}

func Err() ech {
	return errf
}

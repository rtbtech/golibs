// Copyright (c) 2020-2021, Oleg Romanenko (oleg@romanenko.ro)
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package json

import "github.com/valyala/fastjson"

func String(value *fastjson.Value) string {
	if str, err := value.StringBytes(); err == nil && len(str) > 0 {
		return string(str)
	}

	return ""
}

func Int(value *fastjson.Value) int {
	if d, err := value.Int(); err == nil {
		return d
	}

	return 0
}

func Int64(value *fastjson.Value) int64 {
	if d, err := value.Int64(); err == nil {
		return d
	}

	return 0
}

func GetBytes(value *fastjson.Value, key string) []byte {
	if b := value.GetStringBytes(key); b != nil && len(b) > 0 {
		return b
	}

	return nil
}

func GetString(value *fastjson.Value, key string) string {
	if b := value.GetStringBytes(key); b != nil && len(b) > 0 {
		return string(b)
	}

	return ""
}

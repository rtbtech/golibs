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

package bson

import (
	"go.mongodb.org/mongo-driver/bson"
)

func GetDoc(raw bson.Raw, key string) (bson.Raw, bool) {
	if rv, err := raw.LookupErr(key); err == nil {
		if v, ok := rv.DocumentOK(); ok {
			return v, true
		}
	}

	return nil, false
}

func GetArray(raw bson.Raw, key string) (bson.Raw, bool) {
	if rv, err := raw.LookupErr(key); err == nil {
		if v, ok := rv.ArrayOK(); ok {
			return v, true
		}
	}

	return nil, false
}

func GetArrayValues(raw bson.Raw, key string) []bson.RawValue {
	if rv, err := raw.LookupErr(key); err == nil {
		if v, ok := rv.ArrayOK(); ok {
			arrv, _ := v.Values()
			return arrv
		}
	}

	return nil
}

func GetInt64(raw bson.Raw, key string) (int64, bool) {
	if rv, err := raw.LookupErr(key); err == nil {
		if v, ok := rv.Int64OK(); ok {
			return v, true
		}

		if v, ok := rv.Int32OK(); ok {
			return int64(v), true
		}
	}

	return 0, false
}

func GetInt64Quietly(raw bson.Raw, key string, out *int64) {
	if rv, err := raw.LookupErr(key); err == nil && out != nil {
		if v, ok := rv.Int64OK(); ok {
			*out = v
		}

		if v, ok := rv.Int32OK(); ok {
			*out = int64(v)
		}
	}
}

func GetInt32(raw bson.Raw, key string) (int32, bool) {
	if rv, err := raw.LookupErr(key); err == nil {
		if v, ok := rv.Int32OK(); ok {
			return v, true
		}

		if v, ok := rv.Int64OK(); ok {
			return int32(v), true
		}
	}

	return 0, false
}

func GetInt32Quietly(raw bson.Raw, key string, out *int32) {
	if rv, err := raw.LookupErr(key); err == nil && out != nil {
		if v, ok := rv.Int32OK(); ok {
			*out = v
		}

		if v, ok := rv.Int64OK(); ok {
			*out = int32(v)
		}
	}
}

func GetInt(raw bson.Raw, key string) (int64, bool) {
	if rv, err := raw.LookupErr(key); err == nil {
		if v, ok := rv.Int64OK(); ok {
			return v, true
		}

		if v, ok := rv.Int32OK(); ok {
			return int64(v), true
		}
	}

	return 0, false
}

func GetIntQuietly(raw bson.Raw, key string, out *int64) {
	if rv, err := raw.LookupErr(key); err == nil && out != nil {
		if v, ok := rv.Int64OK(); ok {
			*out = v
			return
		}

		if v, ok := rv.Int32OK(); ok {
			*out = int64(v)
			return
		}
	}
}

func GetDouble(raw bson.Raw, key string) (float64, bool) {
	if rv, err := raw.LookupErr(key); err == nil {
		if v, ok := rv.DoubleOK(); ok {
			return v, true
		}
	}

	return 0, false
}

func GetDoubleQuietly(raw bson.Raw, key string, out *float64) {
	if rv, err := raw.LookupErr(key); err == nil && out != nil {
		if v, ok := rv.DoubleOK(); ok {
			*out = v
		}
	}
}

func GetBool(raw bson.Raw, key string) (bool, bool) {
	if rv, err := raw.LookupErr(key); err == nil {
		return rv.BooleanOK()
	}

	return false, false
}

func GetBoolQuietly(raw bson.Raw, key string, out *bool) {
	if rv, err := raw.LookupErr(key); err == nil && out != nil {
		if v, ok := rv.BooleanOK(); ok {
			*out = v
		}
	}
}

func GetString(raw bson.Raw, key string) (string, bool) {
	if rv, err := raw.LookupErr(key); err == nil {
		if v, ok := rv.StringValueOK(); ok {
			return v, true
		}
	}

	return "", false
}

func GetStringDefault(raw bson.Raw, key string, def string) string {
	if rv, err := raw.LookupErr(key); err == nil {
		if v, ok := rv.StringValueOK(); ok {
			return v
		}
	}

	return def
}

func GetStringQuietly(raw bson.Raw, key string, out *string) {
	if rv, err := raw.LookupErr(key); err == nil && out != nil {
		if v, ok := rv.StringValueOK(); ok {
			*out = v
		}
	}
}

func GetStringArrayQuietly(raw bson.Raw, key string, out *[]string) {
	for _, rv := range GetArrayValues(raw, key) {
		if v, ok := rv.StringValueOK(); ok && v != "" {
			*out = append(*out, v)
		}
	}
}

func AppendStringQuietly(raw bson.Raw, key string, out *[]string) {
	if rv, err := raw.LookupErr(key); err == nil && out != nil {
		if v, ok := rv.StringValueOK(); ok {
			*out = append(*out, v)
		}
	}
}

func GetStringArrayQuietlyF(raw bson.Raw, key string, out *[]string, allow func(e string) bool) {
	for _, rv := range GetArrayValues(raw, key) {
		if v, ok := rv.StringValueOK(); ok && v != "" && allow(v) {
			*out = append(*out, v)
		}
	}
}

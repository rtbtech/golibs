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

package xml

import (
	"fmt"
	"html"
	"io"
)

const xmlheader = `<?xml version="1.0" encoding="UTF-8" ?>` + "\r\n"

type xmlFprinter struct {
	dest           io.Writer
	inStartElement bool
	stack          []string
}

func (x *xmlFprinter) push(v string) { x.stack = append(x.stack, v) }

func (x *xmlFprinter) pop() (v string) {
	if x.stack != nil && len(x.stack) > 0 {
		l := len(x.stack) - 1
		v, x.stack = x.stack[l], x.stack[:l]
	}
	return
}

func (x *xmlFprinter) Flush() error {
	if x.inStartElement {
		x.inStartElement = false
		_, e := fmt.Fprintf(x.dest, `>`)
		return e
	}
	return nil
}

func (x *xmlFprinter) Tag(name string) *xmlFprinter {
	if e := x.Flush(); e != nil {
		return x
	}
	x.push(name)
	_, _ = fmt.Fprintf(x.dest, `<%s`, name)
	x.inStartElement = true
	return x
}

func (x *xmlFprinter) End() *xmlFprinter {
	if x.inStartElement {
		x.inStartElement = false
		x.pop()
		_, _ = fmt.Fprintf(x.dest, `/>`)
		return x
	}
	if e := x.Flush(); e != nil {
		return x
	}
	_, _ = fmt.Fprintf(x.dest, `</%s>`, x.pop())
	return x
}

func (x *xmlFprinter) Str(text string) *xmlFprinter {
	if e := x.Flush(); e != nil {
		return x
	}
	_, _ = fmt.Fprintf(x.dest, `%s`, html.EscapeString(text))
	return x
}

func (x *xmlFprinter) CData(text string) *xmlFprinter {
	if e := x.Flush(); e != nil {
		return x
	}
	_, _ = fmt.Fprintf(x.dest, `<![CDATA[%s]]>`, text)
	return x
}

func (x *xmlFprinter) Int32(num int32) *xmlFprinter {
	if x.Flush() != nil {
		return x
	}
	_, _ = fmt.Fprintf(x.dest, `%d`, num)
	return x
}

func (x *xmlFprinter) Int64(num int64) *xmlFprinter {
	if x.Flush() != nil {
		return x
	}
	_, _ = fmt.Fprintf(x.dest, `%d`, num)
	return x
}

func (x *xmlFprinter) Uint32(num uint32) *xmlFprinter {
	if x.Flush() != nil {
		return x
	}
	_, _ = fmt.Fprintf(x.dest, `%d`, num)
	return x
}

func (x *xmlFprinter) Uint64(num uint64) *xmlFprinter {
	if x.Flush() != nil {
		return x
	}
	_, _ = fmt.Fprintf(x.dest, `%d`, num)
	return x
}

func (x *xmlFprinter) Attr(name, value string) *xmlFprinter {
	if !x.inStartElement {
		return x
	}
	_, _ = fmt.Fprintf(x.dest, ` %s="%s"`, name, html.EscapeString(value))
	return x
}

func (x *xmlFprinter) AttrInt32(name string, num int32) *xmlFprinter {
	if !x.inStartElement {
		return x
	}
	_, _ = fmt.Fprintf(x.dest, ` %s="%d"`, name, num)
	return x
}

func (x *xmlFprinter) AttrInt64(name string, num int64) *xmlFprinter {
	if !x.inStartElement {
		return x
	}
	_, _ = fmt.Fprintf(x.dest, ` %s="%d"`, name, num)
	return x
}

func (x *xmlFprinter) AttrUint32(name string, num uint32) *xmlFprinter {
	if !x.inStartElement {
		return x
	}
	_, _ = fmt.Fprintf(x.dest, ` %s="%d"`, name, num)
	return x
}

func (x *xmlFprinter) AttrUint64(name string, num uint64) *xmlFprinter {
	if !x.inStartElement {
		return x
	}
	_, _ = fmt.Fprintf(x.dest, ` %s="%d"`, name, num)
	return x
}

func NewXmlWriter(w io.Writer) *xmlFprinter {
	// xmlheader
	_, _ = fmt.Fprintf(w, xmlheader)
	return &xmlFprinter{dest: w}
}

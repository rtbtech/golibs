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

package domain_test

import (
	"github.com/rtbtech/golibs/strings/domain"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestName(t *testing.T) {
	require.Equal(t, domain.Name("http://www.yandex.ru/aaa/nnn"), "yandex.ru")
	require.Equal(t, domain.Name("//www.yandex.ru/aaa/nnn"), "yandex.ru")
	require.Equal(t, domain.Name("yandex.ru/aaa/nnn"), "yandex.ru")
	require.Equal(t, domain.Name("yandex.ru"), "yandex.ru")
	require.Equal(t, domain.Name("mail.yandex.ru"), "mail.yandex.ru")
}

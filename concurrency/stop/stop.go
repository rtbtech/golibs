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

package stop

import (
	"os"
	"os/signal"
	"sync"
	"sync/atomic"
	"syscall"
)

type info struct {
	v int32
	c chan int32
}

var instance *info
var once sync.Once

func getInstance() *info {
	once.Do(func() {
		instance = &info{v: 0, c: make(chan int32)}
	})
	return instance
}

func Stop() {
	atomic.StoreInt32(&getInstance().v, 1)
	close(getInstance().c)
}

func IsStopped() bool {
	return atomic.LoadInt32(&getInstance().v) > 0
}

func IsStoppedCh() chan int32 {
	return getInstance().c
}

func Wait() {
	osSignals := make(chan os.Signal, 2)
	signal.Notify(osSignals, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-osSignals:
		break
	case <-getInstance().c:
		break
	}

	Stop()
}

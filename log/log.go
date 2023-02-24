// Copyright 2023 panelbase Author. All Rights Reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//      http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package log

import (
	"log"
)

type Logger interface {
	Info(format string, args ...any)
	Debug(format string, args ...any)
	Error(format string, args ...any)
}

type defLog struct {
	lo *log.Logger
}

func (d *defLog) Info(format string, args ...any) { d.lo.Printf(format, args...) }

func (d *defLog) Debug(format string, args ...any) { d.Info(format, args...) }

func (d *defLog) Error(format string, args ...any) { d.Info(format, args...) }

var Default = &defLog{}

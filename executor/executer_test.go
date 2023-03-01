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

package executor

import (
	"github.com/whrwsoftware/panelbase/cmd"
	"reflect"
	"testing"
	"time"
)

func TestExecutor_Add(t *testing.T) {
	etr := NewExecutor(false)
	etr.Add(NewStep(0, "goecho", "hello", "world"))
	if !(etr.steps != nil &&
		len(etr.steps) == 1 &&
		etr.steps[0].id == 0 &&
		etr.steps[0].name == "goecho" &&
		reflect.DeepEqual(etr.steps[0].args, []string{"hello", "world"})) {
		t.Error("test failed!")
	}
}

func TestExecutor_Clear(t *testing.T) {
	etr := NewExecutor(false)
	etr.Add(NewStep(0, "goecho", "hello", "world"))
	etr.Clear()
	if len(etr.steps) != 0 {
		t.Error("test failed!")
	}
}

func TestExecutor_Create(t *testing.T) {
	etr := NewExecutor(false)
	etr.Create("goecho", "hello", "world")
	if len(etr.steps) != 1 {
		t.Error("test failed!")
	}
}

func TestExecutor_Kill(t *testing.T) {
	etr := NewExecutor(false)
	etr.Create("goscan")
	go func() {
		time.Sleep(time.Second)
		etr.Kill()
	}()
	_, _ = etr.Exec()
	t.Log("test ok!")
}

func TestExecutor_NextId(t *testing.T) {
	etr := NewExecutor(false)
	if etr.NextId() != 1 {
		t.Error("test failed!")
	}
	etr.Create("")
	if etr.NextId() != 2 {
		t.Error("test failed!")
	}
}

func TestExecutor_Exec(t *testing.T) {
	{
		etr := NewExecutor(false)
		etr.Create("goscanx")
		if _, err := etr.Exec(); err == nil {
			t.Error("test failed!")
		}
	}
	{
		etr := NewExecutor(false)
		etr.Create("goscan")
		ch := make(chan bool, 1)
		go func() {
			<-ch
			t.Log("test ok!")
			etr.Kill()
		}()
		go func() {
			time.Sleep(time.Second)
			ch <- true
		}()
		_, _ = etr.Exec()
	}
}

func TestExecutor_kill(t *testing.T) {
	{
		etr := NewExecutor(false)
		etr.Create("goscan")
		go func() {
			time.Sleep(time.Second)
			etr.Kill()
			t.Log("test ok!")
		}()
		_, _ = etr.Exec()
	}
}

func TestExecutor_StartSync(t *testing.T) {
	var (
		outC = make(chan Message, 1)
		errC = make(chan Message, 1)
	)
	etr := NewExecutor(false)
	etr.Create("goecho", "hello")
	etr.Create("goecho", "world")
	etr.Create("goecho", "haha")
	etr.Create("goecho", "apple")
	etr.Create("goecho", "coco")
	etr.Chan(outC, errC)
	go func() {
	loop:
		for {
			select {
			case oo, ok := <-outC:
				if !ok {
					break loop
				}
				t.Log("outC:", oo.Step.id, oo.Text, oo.Error)
			case ee, ok := <-errC:
				if !ok {
					break loop
				}
				t.Log("errC:", ee.Step.id, ee.Text, ee.Error)
			}
		}
	}()
	_, _ = etr.Exec()
	t.Log("test ok!")
}

func TestExecutor_StartAsync(t *testing.T) {
	var (
		outC = make(chan Message, 1)
		errC = make(chan Message, 1)
	)
	etr := NewExecutor(true)
	etr.Create("goecho", "hello")
	etr.Create("gomkdir", "tmp")
	etr.Create("gomkdir", "tmp1")
	etr.Create("gomkdir", "tmp2")
	etr.Create("gomkdir", "tmp3")
	etr.Create("goecho", "world")
	etr.Chan(outC, errC)
	go func() {
	loop:
		for {
			select {
			case oo, ok := <-outC:
				if !ok {
					break loop
				}
				t.Log("outC:", oo.Step.id, oo.Text, oo.Error)
			case ee, ok := <-errC:
				if !ok {
					break loop
				}
				t.Log("errC:", ee.Step.id, ee.Text, ee.Error)
			}
		}
	}()
	_, err := etr.Exec()
	t.Log(err)
	t.Log("test ok!")
	{
		_, _, _, _ = cmd.Run("gorm", "tmp")
		_, _, _, _ = cmd.Run("gorm", "tmp1")
		_, _, _, _ = cmd.Run("gorm", "tmp2")
		_, _, _, _ = cmd.Run("gorm", "tmp3")
	}
}

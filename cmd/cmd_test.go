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

package cmd

import (
	"reflect"
	"testing"
	"time"
)

func TestRun(t *testing.T) {
	type args struct {
		name string
		args []string
	}
	tests := []struct {
		name     string
		args     args
		wantOuts []byte
		wantErrs []byte
		wantErr  bool
	}{
		{"test_echo_empty", args{"goecho", []string{""}}, []byte("\n"), []byte{}, false},
		{"test_echo_warp", args{"goecho", []string{"\n"}}, []byte("\n\n"), []byte{}, false},
		{"test_echo_args", args{"goecho", []string{"123"}}, []byte("123\n"), []byte{}, false},
		{"test_echo_args_list", args{"goecho", []string{"123", "456"}}, []byte("123 456\n"), []byte{}, false},
		{"test_echo_args_list2", args{"goecho", []string{"123", "456", "789"}}, []byte("123 456 789\n"), []byte{}, false},
		{"test_echo_err", args{"goechox", []string{}}, nil, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOuts, gotErrs, err := Run(tt.args.name, tt.args.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Run() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotOuts, tt.wantOuts) {
				t.Errorf("Run() gotOuts = %v, want %v", gotOuts, tt.wantOuts)
			}
			if !reflect.DeepEqual(gotErrs, tt.wantErrs) {
				t.Errorf("Run() gotErrs = %v, want %v", gotErrs, tt.wantErrs)
			}
		})
	}
}

func TestStart(t *testing.T) {
	var err error
	var outC = make(chan string, 1)
	var inC = make(chan string, 1)
	go func() {
		for {
			if outStr, ok := <-outC; ok {
				t.Log("out:", outStr)
			}
		}
	}()
	go func() {
		for {
			nn := time.Now().Format(time.RFC3339)
			inC <- nn
			t.Log("in:", nn)
			time.Sleep(time.Second)
		}
	}()
	go func() {
		<-time.After(time.Second * 5)
		inC <- "qqq"
	}()
	if err = Start("goscan", []string{}, outC, nil, inC, nil); err != nil {
		t.Error(err)
	}
}

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

package main

import (
	"fmt"
	"github.com/whrwsoftware/panelbase/apps/svc/postfix"
	"os"
)

func main() {
	var (
		outC = make(chan string, 1)
		errC = make(chan string, 1)
	)
	var app = postfix.GetApp(outC, errC)
	if app == nil {
		fmt.Println("postfix: not support current os")
		return
	}

	go func() {
		defer func() { _ = recover() }()
		defer close(outC)
		for {
			if cc, ok := <-outC; !ok {
				break
			} else {
				fmt.Println(cc)
			}
		}
	}()

	go func() {
		defer func() { _ = recover() }()
		defer close(errC)
		for {
			if cc, ok := <-errC; !ok {
				break
			} else {
				fmt.Println(cc)
			}
		}
	}()

	fmt.Println("postfix app control panel")
	fmt.Println("---")
	fmt.Println("0. check")
	fmt.Println("1. install")
	fmt.Println("2. uninstall")
	fmt.Println("3. start")
	fmt.Println("4. stop")
	fmt.Println("5. restart")
	fmt.Println("6. version")
	fmt.Println("7. configure")
	fmt.Println("q. exit")
	fmt.Println("---")
	var s string
	fmt.Printf(":")
	fmt.Scanf("%s")

	switch s {
	case "1":
	case "2":
	case "3":
	case "4":
	case "5":
	case "6":
	case "7":
	case "q":
		os.Exit(0)
	}

	select {}
}

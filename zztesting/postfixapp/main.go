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
	appConf "github.com/whrwsoftware/panelbase/appconf/postfix"
	"github.com/whrwsoftware/panelbase/appmanager/managers"
	"github.com/whrwsoftware/panelbase/apps"
	"github.com/whrwsoftware/panelbase/apps/postfix"
	"github.com/whrwsoftware/panelbase/appver"
	"github.com/whrwsoftware/panelbase/zvars/oss"
	"os"
)

func main() {
	manager := managers.SqliteManager("/duckcp/apps/manager/data.db")
	currentOS := oss.CurrentOS()
	version := appver.PostfixVersion().Version
	fmt.Println("postfix control")
	fmt.Println("------")
	fmt.Println()
	fmt.Println("current os =>", currentOS)
	fmt.Println("current version =>", version)
	fmt.Println()
	var app = apps.GetApp(postfix.Template(), manager, currentOS)
	if app == nil {
		fmt.Println("postfix: not support current os")
		return
	}

	for {
		fmt.Println("---install control---")
		fmt.Println("10.check 11.install 12.uninstall 13.reinstall")
		fmt.Println()

		fmt.Println("---service control---")
		fmt.Println("20.enable 21.disable 22.start 23.stop 24.restart 25.status")
		fmt.Println()

		fmt.Println("---configuration control---")
		fmt.Println("30.configure 31.load 32.clean")
		fmt.Println()

		fmt.Println("---log control---")
		fmt.Println("90.cat 91.truncate 92.pwd")
		fmt.Println()

		fmt.Println("------")
		fmt.Println("q.exit")
		fmt.Println()

		fmt.Println("---")
		var s string
		fmt.Printf(":")
		fmt.Scanf("%s", &s)

		switch s {

		case "10":
			fmt.Println("start check")
			fmt.Println(app.Check(manager))
			fmt.Println("end check")
			break
		case "11":
			fmt.Println("start install")
			if _, err := app.Check(manager); err != nil {
				fmt.Println(err)
				break
			}
			if _, err := app.Install(version); err != nil {
				fmt.Println(err)
				break
			}
			fmt.Println("end install")
			fmt.Println("postfix installed")
			break
		case "12":
			fmt.Println("start uninstall")
			if _, err := app.Uninstall(version); err != nil {
				fmt.Println(err)
				break
			}
			fmt.Println("end uninstall")
			fmt.Println("postfix uninstalled")
			break
		case "13":
			fmt.Println("start reinstall")
			if _, err := app.Reinstall(version); err != nil {
				fmt.Println(err)
				break
			}
			fmt.Println("end reinstall")
			fmt.Println("postfix reinstalled")
			break

		case "20":
			fmt.Println("start enable")
			if _, err := app.Enable(); err != nil {
				fmt.Println(err)
				break
			}
			fmt.Println("end enable")
			fmt.Println("postfix enable")
			break
		case "21":
			fmt.Println("start disable")
			if _, err := app.Disable(); err != nil {
				fmt.Println(err)
				break
			}
			fmt.Println("end disable")
			fmt.Println("postfix disabled")
			break
		case "22":
			fmt.Println("start start")
			if _, err := app.Start(); err != nil {
				fmt.Println(err)
				break
			}
			fmt.Println("end start")
			fmt.Println("postfix started")
			break
		case "23":
			fmt.Println("start stop")
			if _, err := app.Stop(); err != nil {
				fmt.Println(err)
				break
			}
			fmt.Println("end stop")
			fmt.Println("postfix stopped")
			break
		case "24":
			fmt.Println("start restart")
			if _, err := app.Restart(); err != nil {
				fmt.Println(err)
				break
			}
			fmt.Println("end restart")
			fmt.Println("postfix restarted")
			break
		case "25":
			fmt.Println("start status")
			st, _, _ := app.Status()
			fmt.Println(st)
			fmt.Println("end status")
			fmt.Println("postfix status")
			break

		case "30":
			fmt.Println("start configure")
			var myHostname string
			var myDomain string
			fmt.Println("Please input my hostname for config:")
			fmt.Scanf("%s", &myHostname)
			fmt.Println("Please input my domain for config:")
			fmt.Scanf("%s", &myDomain)
			fmt.Println("your configuration")
			fmt.Println("---")
			fmt.Println("myhostname=", myHostname)
			fmt.Println("mydomain=", myDomain)
			fmt.Println("Are you sure start configure?")
			var yesOrNo string
			fmt.Scanf("%s", yesOrNo)
			if yesOrNo == "no" {
				break
			}
			if err := app.Configure(appConf.ConfBindMap(
				appConf.OptMainCf{MyHostname: myHostname, MyDomain: myDomain, Version: version},
				appConf.OptMasterCf{},
			)); err != nil {
				fmt.Println(err)
				break
			}
			fmt.Println("end configure")
			fmt.Println("postfix configured")
			break
		case "31":
			fmt.Println("start load")
			if val, err := app.Load(appConf.DistMainCf); err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(val)
			}
			fmt.Println("end load")
			fmt.Println("postfix loaded")
			break
		case "32":
			fmt.Println("start clean")
			if err := app.Clean(); err != nil {
				fmt.Println(err)
			}
			fmt.Println("end clean")
			fmt.Println("postfix cleaned")
			break

		case "90":
			fmt.Println("start cat")
			if logBuf, err := os.ReadFile(app.File()); err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(string(logBuf))
			}
			fmt.Println("end cat")
			fmt.Println("postfix cut")
			break
		case "91":
			fmt.Println("start truncate")
			if err := app.Truncate(); err != nil {
				fmt.Println(err)
			}
			fmt.Println("end truncate")
			fmt.Println("postfix truncated")
			break
		case "92":
			fmt.Println("start pwd")
			fmt.Println(app.File())
			fmt.Println("end pwd")
			fmt.Println("postfix pwd")
			break

		case "q":
			fmt.Println("bye")
			os.Exit(0)
		}
		fmt.Print("\033[H\033[2J")
	}
}

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

package appmanager

type Type uint

const (
	Other       Type = iota // 其他
	Environment             // 运行环境
	SystemTool              // 系统工具
	Plugin                  // 插件
	ThirdParty              // 第三方
	Suite                   // 套件
)

type App struct {
	AppName     string            // Nginx 1.22.1
	Name        string            // nginx
	Version     string            // 1.22.1
	Pkg         string            // nginx@1.22.1
	Icon        string            // base64;
	Provider    string            // 官方
	Description string            // 轻量级，占有内存少，并发能力强
	Workdir     string            // /duckcp/apps/$Name/$Version
	Type        Type              // 分类
	Tag         map[string]string // Tag

	Installed bool
}

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

package appver

var nginxMaxVersionId = 1000

func nginxNextVersionId() (vi int) {
	vi = nginxMaxVersionId
	nginxMaxVersionId--
	return
}

var nginxVer = []*ver{
	Ver("1.23.3", "1.23.3", nginxNextVersionId(), ""),
	Ver("1.22.1", "1.22.1", nginxNextVersionId(), ""),
	Ver("1.20.2", "1.20.2", nginxNextVersionId(), ""),
	Ver("1.18.0", "1.18.0", nginxNextVersionId(), ""),
	Ver("1.16.1", "1.16.1", nginxNextVersionId(), ""),
	Ver("1.14.2", "1.14.2", nginxNextVersionId(), ""),
	Ver("1.12.2", "1.12.2", nginxNextVersionId(), ""),
	Ver("1.10.3", "1.10.3", nginxNextVersionId(), ""),
	Ver("1.8.1", "1.8.1", nginxNextVersionId(), ""),
}

var Nginx = &struct {
	Name        string
	Provider    string
	Description string
	Ver         []*ver
}{"nginx", "官方", "", nginxVer}

func NginxMinVersion() *ver { return Nginx.Ver[len(Nginx.Ver)-1] }
func NginxMaxVersion() *ver { return Nginx.Ver[0] }

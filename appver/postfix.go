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

var postfixMaxVersionId = 1000

func postfixNextVersionId() (vi int) {
	vi = postfixMaxVersionId
	postfixMaxVersionId--
	return
}

var postfixVer = []*ver{
	Ver("postfix@3.7.4", "3.7.4", postfixNextVersionId(), "/duckcp/logs/duckcp-postfix.log"),
}

var Postfix = &struct {
	Name        string
	Provider    string
	Description string
	Ver         []*ver
}{"postfix", "官方", "", postfixVer}

func PostfixMinVersion() *ver { return Postfix.Ver[len(Postfix.Ver)-1] }
func PostfixMaxVersion() *ver { return Postfix.Ver[0] }
func PostfixVersion() *ver    { return PostfixMaxVersion() }

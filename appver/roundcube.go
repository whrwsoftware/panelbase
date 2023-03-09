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

var roundcubeMaxVersionId = 1000

func roundcubeNextVersionId() (vi int) {
	vi = roundcubeMaxVersionId
	roundcubeMaxVersionId--
	return
}

var roundcubeVer = []*ver{
	Ver("roundcube@1.6.1", "1.6.1", roundcubeNextVersionId(), "/duckcp/logs/roundcube.log"),
}

var Roundcube = &struct {
	Name        string
	Provider    string
	Description string
	Ver         []*ver
}{"roundcube", "官方", "", roundcubeVer}

func RoundcubeMinVersion() *ver { return Roundcube.Ver[len(Roundcube.Ver)-1] }
func RoundcubeMaxVersion() *ver { return Roundcube.Ver[0] }
func RoundcubeVersion() *ver    { return RoundcubeMaxVersion() }

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

var dovecotMaxVersionId = 1000

func dovecotNextVersionId() (vi int) {
	vi = dovecotMaxVersionId
	dovecotMaxVersionId--
	return
}

var dovecotVer = []*ver{
	Ver("dovecot@2.3.17", "2.3.17", dovecotNextVersionId(), "/duckcp/logs/duckcp-dovecot.log"),
}

var Dovecot = &struct {
	Name        string
	Provider    string
	Description string
	Ver         []*ver
}{"dovecot", "官方", "", dovecotVer}

func DovecotMinVersion() *ver { return Dovecot.Ver[len(Dovecot.Ver)-1] }
func DovecotMaxVersion() *ver { return Dovecot.Ver[0] }
func DovecotVersion() *ver    { return DovecotMaxVersion() }

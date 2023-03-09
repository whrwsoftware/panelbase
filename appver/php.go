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

var phpMaxVersionId = 1000

func phpNextVersionId() (vi int) {
	vi = phpMaxVersionId
	phpMaxVersionId--
	return
}

var phpVer = []*ver{
	Ver("php-56", "5.6", phpNextVersionId(), ""),
	Ver("php-70", "7.0", phpNextVersionId(), ""),
	Ver("php-71", "7.1", phpNextVersionId(), ""),
	Ver("php-72", "7.2", phpNextVersionId(), ""),
	Ver("php-73", "7.3", phpNextVersionId(), ""),
	Ver("php-74", "7.4", phpNextVersionId(), ""),
	Ver("php-80", "8.0", phpNextVersionId(), ""),
	Ver("php-81", "8.1", phpNextVersionId(), ""),
	Ver("php-82", "8.2", phpNextVersionId(), ""),
}

var Php = &struct {
	Name        string
	Provider    string
	Description string
	Ver         []*ver
}{"php", "官方", "", phpVer}

func PhpMinVersion() *ver { return Php.Ver[len(Php.Ver)-1] }
func PhpMaxVersion() *ver { return Php.Ver[0] }
func Php73Version() *ver  { return Php.Ver[4] }

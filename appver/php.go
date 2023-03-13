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

var phpVer = []*PhpVer{
	PhpVerV("php82 php82-php-fpm php82-php-pdo php82-php-dg php82-php-mbstring", "php82-php-fpm", "8.2", phpNextVersionId(), 82, "/duckcp/logs/duckcp-php82.log"),
	PhpVerV("php81 php81-php-fpm php81-php-pdo php81-php-dg php81-php-mbstring", "php81-php-fpm", "8.1", phpNextVersionId(), 81, "/duckcp/logs/duckcp-php81.log"),
	PhpVerV("php80 php80-php-fpm php80-php-pdo php80-php-dg php80-php-mbstring", "php80-php-fpm", "8.0", phpNextVersionId(), 80, "/duckcp/logs/duckcp-php80.log"),
	PhpVerV("php74 php74-php-fpm php74-php-pdo php74-php-dg php74-php-mbstring", "php74-php-fpm", "7.4", phpNextVersionId(), 74, "/duckcp/logs/duckcp-php74.log"),
	PhpVerV("php73 php73-php-fpm php73-php-pdo php73-php-dg php73-php-mbstring", "php73-php-fpm", "7.3", phpNextVersionId(), 73, "/duckcp/logs/duckcp-php73.log"),
	PhpVerV("php72 php72-php-fpm php72-php-pdo php72-php-dg php72-php-mbstring", "php72-php-fpm", "7.2", phpNextVersionId(), 72, "/duckcp/logs/duckcp-php72.log"),
	PhpVerV("php71 php71-php-fpm php71-php-pdo php71-php-dg php71-php-mbstring", "php71-php-fpm", "7.1", phpNextVersionId(), 71, "/duckcp/logs/duckcp-php71.log"),
	PhpVerV("php70 php70-php-fpm php70-php-pdo php70-php-dg php70-php-mbstring", "php70-php-fpm", "7.0", phpNextVersionId(), 70, "/duckcp/logs/duckcp-php70.log"),
	PhpVerV("php56 php56-php-fpm php56-php-pdo php56-php-dg php56-php-mbstring", "php56-php-fpm", "5.6", phpNextVersionId(), 56, "/duckcp/logs/duckcp-php56.log"),
	PhpVerV("php55 php55-php-fpm php55-php-pdo php55-php-dg php55-php-mbstring", "php55-php-fpm", "5.5", phpNextVersionId(), 55, "/duckcp/logs/duckcp-php55.log"),
	PhpVerV("php54 php54-php-fpm php54-php-pdo php54-php-dg php54-php-mbstring", "php54-php-fpm", "5.4", phpNextVersionId(), 54, "/duckcp/logs/duckcp-php54.log"),
}

var Php = &struct {
	Name string
	Ver  []*PhpVer
}{"php", phpVer}

func PhpMinVersion() *PhpVer { return Php.Ver[len(Php.Ver)-1] }
func PhpMaxVersion() *PhpVer { return Php.Ver[0] }
func Php73Version() *PhpVer  { return Php.Ver[4] }
func PhpVersion() *PhpVer    { return PhpMaxVersion() }

type PhpVer struct {
	Pkg       string
	Service   string
	Version   string
	VersionId int
	VersionV  int
	LogFile   string
}

func PhpVerV(pkg string, service string, version string, versionId int, versionV int, logFile string) *PhpVer {
	return &PhpVer{Pkg: pkg, Service: service, Version: version, VersionId: versionId, VersionV: versionV, LogFile: logFile}
}

func FindPhpVerByV(v int) (pv *PhpVer) {
	for _, vv := range Php.Ver {
		if vv.VersionV == v {
			pv = vv
			break
		}
	}
	return
}

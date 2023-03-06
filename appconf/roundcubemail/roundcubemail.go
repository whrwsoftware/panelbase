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

package roundcubemail

import (
	"embed"
	"github.com/whrwsoftware/panelbase/appconf"
	"os"
)

var (
	//go:embed config.inc.php
	fs embed.FS
)

const (
	NameConfigIncPhp = "config.inc.php"
)

const (
	DistConfigIncPhp = "/var/www/roundcubemail/config/config.inc.php"
)

var (
	perm os.FileMode = 0644
)

var (
	ConfBindConfigIncPhp = appconf.NewConfBind[any](fs, NameConfigIncPhp, DistConfigIncPhp, perm)
)

var ConfBinds = []*appconf.ConfBind[any]{
	ConfBindConfigIncPhp,
}

type (
	OptConfigIncPhp struct {
		DbFile         string // /tmp/mail.db
		ImapHost       string // localhost:143
		SmtpHost       string // localhost:25
		SupportUrl     string // mail.999dns.xyz
		UsernameDomain string // 999dns.xyz
	}
)
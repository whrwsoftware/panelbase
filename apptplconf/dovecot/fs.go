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

package postfix

import (
	"embed"
	"github.com/whrwsoftware/panelbase/apptplconf"
)

//go:embed dovecot.conf conf.d/10-auth.conf conf.d/10-mail.conf conf.d/10-master.conf
var fs embed.FS

const (
	dovecotConfDist   = "/etc/dovecot/dovecot.conf"
	confD10AuthDist   = "/etc/dovecot/conf.d/10-auth.conf"
	confD10MailDist   = "/etc/dovecot/conf.d/10-mail.conf"
	confD10MasterDist = "/etc/dovecot/conf.d/10-master.conf"
)

type Args struct {
	MyHostname string
	MyDomain   string
	MyOrigin   string
}

func DovecotConf(a Args) (err error) {
	return apptplconf.Gen(fs, "dovecot.conf", dovecotConfDist, a, 0600)
}

func ConfD10Auth() (err error) {
	return apptplconf.Gen(fs, "conf.d/10-auth.conf", confD10AuthDist, nil, 0600)
}

func ConfD10Mail() (err error) {
	return apptplconf.Gen(fs, "conf.d/10-mail.conf", confD10MailDist, nil, 0600)
}

func ConfD10Master() (err error) {
	return apptplconf.Gen(fs, "conf.d/10-master.conf", confD10MasterDist, nil, 0600)
}

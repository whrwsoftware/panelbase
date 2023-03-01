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

package dovecot

import (
	_ "embed"
	"github.com/whrwsoftware/panelbase/appconf"
)

var (
	//go:embed dovecot.conf
	FSDovecotConf string

	//go:embed conf.d/10-auth.conf
	FSConfD10AuthConf string

	//go:embed conf.d/10-mail.conf
	FSConfD10MailConf string

	//go:embed conf.d/10-master.conf
	FSConfD10MasterConf string
)

const (
	DistDovecotConf       = "/etc/dovecot/dovecot.conf"
	DistConfD10AuthConf   = "/etc/dovecot/conf.d/10-auth.conf"
	DistConfD10MailConf   = "/etc/dovecot/conf.d/10-mail.conf"
	DistConfD10MasterConf = "/etc/dovecot/conf.d/10-master.conf"
)

type (
	Opt struct {
		MyHostname string
		MyDomain   string
		MyOrigin   string
	}
	OptConfD10Auth   struct{}
	OptConfD10Mail   struct{}
	OptConfD10Master struct{}
)

var (
	GenDovecotConf       = appconf.Gen[Opt]
	GenConfD10AuthConf   = appconf.Gen[OptConfD10Auth]
	GenConfD10MailConf   = appconf.Gen[OptConfD10Mail]
	GenConfD10MasterConf = appconf.Gen[OptConfD10Master]
)

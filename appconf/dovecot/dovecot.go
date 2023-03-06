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
	"embed"
	"github.com/whrwsoftware/panelbase/appconf"
	"os"
)

var (
	//go:embed dovecot.conf conf.d/10-auth.conf conf.d/10-mail.conf conf.d/10-master.conf conf.d/10-ssl.conf conf.d/10-ssl.conf conf.d/15-mailboxes.conf conf.d/20-imap.conf conf.d/20-lmtp.conf
	fs embed.FS
)

const (
	NameDovecotConf          = "dovecot.conf"
	NameConfD10AuthConf      = "conf.d/10-auth.conf"
	NameConfD10MailConf      = "conf.d/10-mail.conf"
	NameConfD10MasterConf    = "conf.d/10-master.conf"
	NameConfD10SslConf       = "conf.d/10-ssl.conf"
	NameConfD15MailboxesConf = "conf.d/15-mailboxes.conf"
	NameConfD20ImapConf      = "conf.d/20-imap.conf"
	NameConfD20LmtpConf      = "conf.d/20-lmtp.conf"
)

const (
	DistDovecotConf          = "/etc/dovecot/dovecot.conf"
	DistConfD10AuthConf      = "/etc/dovecot/conf.d/10-auth.conf"
	DistConfD10MailConf      = "/etc/dovecot/conf.d/10-mail.conf"
	DistConfD10MasterConf    = "/etc/dovecot/conf.d/10-master.conf"
	DistConfD10SslConf       = "/etc/dovecot/conf.d/10-ssl.conf"
	DistConfD15MailboxesConf = "/etc/dovecot/conf.d/15-mailboxes.conf"
	DistConfD20ImapConf      = "/etc/dovecot/conf.d/20-imap.conf"
	DistConfD20LmtpConf      = "/etc/dovecot/conf.d/20-lmtp.conf"
)

var (
	perm os.FileMode = 0644
)

var (
	ConfBindDovecotConf          = appconf.NewConfBind[any](fs, NameDovecotConf, DistDovecotConf, perm)
	ConfBindConfD10AuthConf      = appconf.NewConfBind[any](fs, NameConfD10AuthConf, DistConfD10AuthConf, perm)
	ConfBindConfD10MailConf      = appconf.NewConfBind[any](fs, NameConfD10MailConf, DistConfD10MailConf, perm)
	ConfBindConfD10MasterConf    = appconf.NewConfBind[any](fs, NameConfD10MasterConf, DistConfD10MasterConf, perm)
	ConfBindConfD10SslConf       = appconf.NewConfBind[any](fs, NameConfD10SslConf, DistConfD10SslConf, perm)
	ConfBindConfD15MailboxesConf = appconf.NewConfBind[any](fs, NameConfD15MailboxesConf, DistConfD15MailboxesConf, perm)
	ConfBindConfD20ImapConf      = appconf.NewConfBind[any](fs, NameConfD20ImapConf, DistConfD20ImapConf, perm)
	ConfBindConfD20LmtpConf      = appconf.NewConfBind[any](fs, NameConfD20LmtpConf, DistConfD20LmtpConf, perm)
)

var ConfBinds = []*appconf.ConfBind[any]{
	ConfBindDovecotConf,
	ConfBindConfD10AuthConf,
	ConfBindConfD10MailConf,
	ConfBindConfD10MasterConf,
	ConfBindConfD10SslConf,
	ConfBindConfD15MailboxesConf,
	ConfBindConfD20ImapConf,
	ConfBindConfD20LmtpConf,
}

type (
	OptDovecotConf     struct{}
	OptConfD10AuthConf struct {
		DisablePlaintextAuth string // yes or no
		AuthDebug            string // yes or no
		AuthDebugPasswords   string // yes or no
	}
	OptConfD10MailConf   struct{}
	OptConfD10MasterConf struct{}
	OptConfD10SslConf    struct {
		SSL                    string // required or yes or no
		SSLCert                string
		SSLKey                 string
		SSLDhParametersLength  int    // default 2048
		SSLPreferServerCiphers string // yes or no
	}
	OptConfD15MailboxesConf struct{}
	OptConfD20ImapConf      struct{}
	OptConfD20LmtpConf      struct{}
)

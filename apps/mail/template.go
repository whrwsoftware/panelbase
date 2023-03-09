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

package mail

import (
	"github.com/whrwsoftware/panelbase/app"
	"github.com/whrwsoftware/panelbase/apps"
	"github.com/whrwsoftware/panelbase/appver"
	"github.com/whrwsoftware/panelbase/zvars/oss"
)

type template struct {
	Postfix   apps.GetAppTemplate
	Dovecot   apps.GetAppTemplate
	Roundcube apps.GetAppTemplate
}

func Template(postfix apps.GetAppTemplate, dovecot apps.GetAppTemplate, roundcube apps.GetAppTemplate) *template {
	return &template{Postfix: postfix, Dovecot: dovecot, Roundcube: roundcube}
}

func (t *template) CentOS7() apps.GetAppFunc {
	return func(outC, errC chan string) app.Applicable {
		return NewApp(
			apps.GetApp(t.Postfix, oss.CentOS7, outC, errC),
			apps.GetApp(t.Dovecot, oss.CentOS7, outC, errC),
			apps.GetApp(t.Roundcube, oss.CentOS7, outC, errC),
			appver.PostfixVersion().Version,
			appver.DovecotVersion().Version,
			appver.RoundcubeMaxVersion().Version,
		)
	}
}

func (t *template) CentOS8() apps.GetAppFunc {
	return func(outC, errC chan string) app.Applicable {
		return NewApp(
			apps.GetApp(t.Postfix, oss.CentOS8, outC, errC),
			apps.GetApp(t.Dovecot, oss.CentOS8, outC, errC),
			apps.GetApp(t.Roundcube, oss.CentOS8, outC, errC),
			appver.PostfixVersion().Version,
			appver.DovecotVersion().Version,
			appver.RoundcubeMaxVersion().Version,
		)
	}
}

func (t *template) Ubuntu() apps.GetAppFunc {
	return func(outC, errC chan string) app.Applicable {
		return NewApp(
			apps.GetApp(t.Postfix, oss.Ubuntu, outC, errC),
			apps.GetApp(t.Dovecot, oss.Ubuntu, outC, errC),
			apps.GetApp(t.Roundcube, oss.Ubuntu, outC, errC),
			appver.PostfixVersion().Version,
			appver.DovecotVersion().Version,
			appver.RoundcubeMaxVersion().Version,
		)
	}
}

func (t *template) Debian() apps.GetAppFunc {
	return func(outC, errC chan string) app.Applicable {
		return NewApp(
			apps.GetApp(t.Postfix, oss.Debian, outC, errC),
			apps.GetApp(t.Dovecot, oss.Debian, outC, errC),
			apps.GetApp(t.Roundcube, oss.Debian, outC, errC),
			appver.PostfixVersion().Version,
			appver.DovecotVersion().Version,
			appver.RoundcubeMaxVersion().Version,
		)
	}
}

func (t *template) Arch() apps.GetAppFunc {
	return func(outC, errC chan string) app.Applicable {
		return NewApp(
			apps.GetApp(t.Postfix, oss.Arch, outC, errC),
			apps.GetApp(t.Dovecot, oss.Arch, outC, errC),
			apps.GetApp(t.Roundcube, oss.Arch, outC, errC),
			appver.PostfixVersion().Version,
			appver.DovecotVersion().Version,
			appver.RoundcubeMaxVersion().Version,
		)
	}
}

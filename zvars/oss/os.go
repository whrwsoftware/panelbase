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

package oss

import (
	sos "os"

	"regexp"
)

type OS string

//NAME="Arch Linux"
//PRETTY_NAME="Arch Linux"
//ID=arch
//BUILD_ID=rolling
//VERSION_ID=TEMPLATE_VERSION_ID
//ANSI_COLOR="38;2;23;147;209"
//HOME_URL="https://archlinux.org/"
//DOCUMENTATION_URL="https://wiki.archlinux.org/"
//SUPPORT_URL="https://bbs.archlinux.org/"
//BUG_REPORT_URL="https://bugs.archlinux.org/"
//PRIVACY_POLICY_URL="https://terms.archlinux.org/docs/privacy-policy/"
//LOGO=archlinux-logo

const (
	Unknown OS = "unknown"
	CentOS6    = "centos6"
	CentOS7    = "centos7"
	CentOS8    = "centos8"
	Debian     = "debian"
	Ubuntu     = "ubuntu"
	Arch       = "arch"
)

const (
	redHatReleasePath = "/etc/redhat-release"
	osReleasePath     = "/etc/os-release"
)

func getOSId(str string) (id, versionId string) {
	{
		m := regexp.MustCompile(`ID="?([a-zA-Z]+)"?`).FindAllStringSubmatch(str, 1)
		if m != nil && len(m) > 0 {
			if mm := m[0]; len(mm) > 1 {
				id = mm[1]
			}
		}
	}
	{
		m := regexp.MustCompile(`VERSION_ID="?([.\d]+)"?`).FindAllStringSubmatch(str, -1)
		if m != nil && len(m) > 0 {
			if mm := m[0]; len(mm) > 1 {
				versionId = mm[1]
			}
		}
	}
	return
}

func CurrentOS() (os OS) {
	if bytes, _ := sos.ReadFile(osReleasePath); bytes != nil && len(bytes) > 0 {
		id, versionId := getOSId(string(bytes))
		if id != "" {
			switch id {
			case "centos":
				switch versionId {
				case "7":
					return CentOS7
				case "8":
					return CentOS8
				}
			case "ubuntu":
				return Ubuntu
			case "debian":
				return Debian
			case "arch":
				return Arch
			}
		}
	} else if bytes, _ = sos.ReadFile(redHatReleasePath); bytes != nil && len(bytes) > 0 {
		// CentOS release 6.10 (Final) -> CentOS release 6.9 (Final)
		if b := regexp.MustCompile(`^CentOS release 6\.\d+\s\(Final\)$`).MatchString(string(bytes)); b {
			return CentOS6
		}
	}
	return Unknown
}

type release struct {
	release   string
	id        string
	versionId string
}

var (
	centos6Release = release{`CentOS release 6.9 (Final)`, "centos", ""}

	centos7Release = release{`NAME="CentOS Linux"
VERSION="7 (Core)"
ID="centos"
ID_LIKE="rhel fedora"
VERSION_ID="7"
PRETTY_NAME="CentOS Linux 7 (Core)"
ANSI_COLOR="0;31"
CPE_NAME="cpe:/o:centos:centos:7"
HOME_URL="https://www.centos.org/"
BUG_REPORT_URL="https://bugs.centos.org/"
CENTOS_MANTISBT_PROJECT="CentOS-7"
CENTOS_MANTISBT_PROJECT_VERSION="7"
REDHAT_SUPPORT_PRODUCT="centos"
REDHAT_SUPPORT_PRODUCT_VERSION="7"`, "centos", "7"}

	centos8Release = release{`NAME="CentOS Linux"
VERSION="8"
ID="centos"
ID_LIKE="rhel fedora"
VERSION_ID="8"
PLATFORM_ID="platform:el8"
PRETTY_NAME="CentOS Linux 8"
ANSI_COLOR="0;31"
CPE_NAME="cpe:/o:centos:centos:8"
HOME_URL="https://centos.org/"
BUG_REPORT_URL="https://bugs.centos.org/"
CENTOS_MANTISBT_PROJECT="CentOS-8"
CENTOS_MANTISBT_PROJECT_VERSION="8"`, "centos", "8"}

	ubuntuRelease = release{`PRETTY_NAME="Ubuntu 22.04.1 LTS"
NAME="Ubuntu"
VERSION_ID="22.04"
VERSION="22.04.1 LTS (Jammy Jellyfish)"
VERSION_CODENAME=jammy
ID=ubuntu
ID_LIKE=debian
HOME_URL="https://www.ubuntu.com/"
SUPPORT_URL="https://help.ubuntu.com/"
BUG_REPORT_URL="https://bugs.launchpad.net/ubuntu/"
PRIVACY_POLICY_URL="https://www.ubuntu.com/legal/terms-and-policies/privacy-policy"
UBUNTU_CODENAME=jammy`, "ubuntu", "22.04"}

	debianRelease = release{`PRETTY_NAME="Debian GNU/Linux 11 (bullseye)"
NAME="Debian GNU/Linux"
VERSION_ID="11"
VERSION="11 (bullseye)"
VERSION_CODENAME=bullseye
ID=debian
HOME_URL="https://www.debian.org/"
SUPPORT_URL="https://www.debian.org/support"
BUG_REPORT_URL="https://bugs.debian.org/`, "debian", "11"}

	archRelease = release{`NAME="Arch Linux"
PRETTY_NAME="Arch Linux"
ID=arch
BUILD_ID=rolling
VERSION_ID=TEMPLATE_VERSION_ID
ANSI_COLOR="38;2;23;147;209"
HOME_URL="https://archlinux.org/"
DOCUMENTATION_URL="https://wiki.archlinux.org/"
SUPPORT_URL="https://bbs.archlinux.org/"
BUG_REPORT_URL="https://bugs.archlinux.org/"
PRIVACY_POLICY_URL="https://terms.archlinux.org/docs/privacy-policy/"
LOGO=archlinux-logo"`, "arch", ""}
)

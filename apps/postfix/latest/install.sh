{{if eq .OS "centos7"}}
yum remove -y postfix postfix3 postfix3-sqlite
yum --nogpg install -y https://mirror.ghettoforge.org/distributions/gf/gf-release-latest.gf.el7.noarch.rpm
yum-config-manager --enable gf gf-plus
yum install -y postfix3 postfix3-sqlite
{{else if eq .OS "centos8" }}
yum remove -y postfix postfix3 postfix3-sqlite
yum --nogpg install -y https://mirror.ghettoforge.org/distributions/gf/gf-release-latest.gf.el7.noarch.rpm
yum-config-manager --enable gf gf-plus
yum install -y postfix3 postfix3-sqlite

{{else if eq .OS "ubuntu" }}

{{else if eq .OS "debian" }}

{{else if eq .OS "arch" }}

{{end}}
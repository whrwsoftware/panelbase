{{if eq .OS "centos7"}}
yum remove -y postfix postfix3 postfix3-sqlite
yum-config-manager --disable gf gf-plus
yum remove -y gf-release
{{else if eq .OS "centos8" }}
yum remove -y postfix postfix3 postfix3-sqlite
yum-config-manager --disable gf gf-plus
yum remove -y gf-release
{{else if eq .OS "ubuntu" }}

{{else if eq .OS "debian" }}

{{else if eq .OS "arch" }}

{{end}}

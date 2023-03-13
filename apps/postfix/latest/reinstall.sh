{{if eq .OS "centos7"}}
yum reinstall -y postfix3 postfix3-sqlite
{{else if eq .OS "centos8" }}
yum reinstall -y postfix3 postfix3-sqlite
{{else if eq .OS "ubuntu" }}

{{else if eq .OS "debian" }}

{{else if eq .OS "arch" }}

{{end}}

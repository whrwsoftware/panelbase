{{if eq .OS "centos7"}}
yum install -y http://rpms.remirepo.net/enterprise/remi-release-7.rpm
yum install -y {{.Pkg}}
{{else if eq .OS "centos8" }}
yum install -y http://rpms.remirepo.net/enterprise/remi-release-8.rpm
yum install -y {{.Pkg}}
{{else if eq .OS "ubuntu" }}

{{else if eq .OS "debian" }}

{{else if eq .OS "arch" }}

{{end}}

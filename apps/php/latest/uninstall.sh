{{if eq .OS "centos7"}}
yum remove -y remi-release
yum remove -y {{.Pkg}}
{{else if eq .OS "centos8" }}
yum remove -y remi-release
yum remove -y {{.Pkg}}
{{else if eq .OS "ubuntu" }}

{{else if eq .OS "debian" }}

{{else if eq .OS "arch" }}

{{end}}
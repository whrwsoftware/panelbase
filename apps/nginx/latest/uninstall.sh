{{if eq .OS "centos7"}}
yum remove -y nginx
rm -rf /etc/yum.repos.d/nginx.repo
yum clean all
{{else if eq .OS "centos8" }}
yum remove -y nginx
rm -rf /etc/yum.repos.d/nginx.repo
yum clean all
{{else if eq .OS "ubuntu" }}

{{else if eq .OS "debian" }}

{{else if eq .OS "arch" }}

{{end}}

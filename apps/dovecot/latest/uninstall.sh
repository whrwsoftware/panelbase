{{if eq .OS "centos7"}}
yum remove -y dovecot
rm -rf /etc/yum.repos.d/dovecot.repo
yum clean all
{{else if eq .OS "centos8" }}
yum remove -y dovecot
rm -rf /etc/yum.repos.d/dovecot.repo
yum clean all
{{else if eq .OS "ubuntu" }}

{{else if eq .OS "debian" }}

{{else if eq .OS "arch" }}

{{end}}
{{if eq .OS "centos7"}}
yum remove -y dovecot
cat>/etc/yum.repos.d/dovecot.repo<<EOF
[dovecotrepo]
name=Dovecot 2.3 CentOS \$releasever - \$basearch
baseurl=http://repo.dovecot.org/ce-2.3.17/centos/\$releasever/RPMS/\$basearch
gpgkey=https://repo.dovecot.org/DOVECOT-REPO-GPG
gpgcheck=1
enabled=1
EOF
yum install -y dovecot
{{else if eq .OS "centos8" }}
yum remove -y dovecot
cat>/etc/yum.repos.d/dovecot.repo<<EOF
[dovecotrepo]
name=Dovecot 2.3 CentOS \$releasever - \$basearch
baseurl=http://repo.dovecot.org/ce-2.3.17/centos/\$releasever/RPMS/\$basearch
gpgkey=https://repo.dovecot.org/DOVECOT-REPO-GPG
gpgcheck=1
enabled=1
EOF
yum install -y dovecot
{{else if eq .OS "ubuntu" }}

{{else if eq .OS "debian" }}

{{else if eq .OS "arch" }}

{{end}}
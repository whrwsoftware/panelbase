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

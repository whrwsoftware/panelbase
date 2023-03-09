yum remove -y postfix postfix3 postfix3-sqlite
yum --nogpg install -y https://mirror.ghettoforge.org/distributions/gf/gf-release-latest.gf.el7.noarch.rpm
yum-config-manager --enable gf gf-plus
yum install -y postfix3 postfix3-sqlite

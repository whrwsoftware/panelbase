rm -rf /duckcp/apps/roundcube
wget -O roundcube.tar.gz https://github.com/roundcube/roundcubemail/releases/download/1.6.1/roundcubemail-1.6.1-complete.tar.gz
tar xvf roundcube.tar.gz
rm -rf roundcube.tar.gz
mv roundcubemail-1.6.1 /duckcp/apps/roundcube
# install mecab

http://packages.groonga.org/centos/latest/x86_64/Packages/

```bash
sudo su -
yum install -y http://packages.groonga.org/centos/latest/x86_64/Packages/mecab-0.996-1.el6.x86_64.rpm
yum install -y http://packages.groonga.org/centos/latest/x86_64/Packages/mecab-devel-0.996-1.el6.x86_64.rpm
yum install -y http://packages.groonga.org/centos/latest/x86_64/Packages/mecab-ipadic-2.7.0.20070801-13.el6.1.x86_64.rpm
```

# install mecab-ipadic-neologd

https://github.com/neologd/mecab-ipadic-neologd/blob/master/README.ja.md

```bash
sudo su -
yum install -y patch
cd /usr/local/src
git clone --depth 1 https://github.com/neologd/mecab-ipadic-neologd.git
cd mecab-ipadic-neologd
./bin/install-mecab-ipadic-neologd -n -y
```

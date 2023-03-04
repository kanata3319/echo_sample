# echo_sample

## GOインストール
* 参考：https://www.youtube.com/watch?v=GRSYl3tPIgk&t=361s

### goenvインストール
* https://github.com/syndbg/goenv
```
git clone https://github.com/syndbg/goenv.git ~/.goenv
# PATHの設定
echo 'export GOENV_ROOT="$HOME/.goenv"' >> ~/.bash_profile
echo 'export PATH="$GOENV_ROOT/bin:$PATH"' >> ~/.bash_profile
echo 'eval "$(goenv init -)"' >> ~/.bash_profile
echo 'export PATH="$GOROOT/bin:$PATH"' >> ~/.bash_profile
echo 'export PATH="$PATH:$GOPATH/bin"' >> ~/.bash_profile
cat .bash_profile
exec $SHELL
```
### GO1.19インストール
```
goenv install -l
goenv install 1.19.6
goenv versions
goenv global 1.19.6
goenv version
```
### Echoインストール
```
# -uオプション: echoのインストールとその依存関係にあるパッケージを強制的に更新
go get -u github.com/labstack/echo
```
# go-helloworld

- 公式ドキュメント
    - https://go.dev/doc/
- 日本語翻訳
    - https://xn--go-hh0g6u.com/doc/

# Ubuntu にGoをインストール
https://github.com/golang/go/wiki/Ubuntu
UbuntuのようなDebianベースのシステムには他にもいくつかのオプションがあります。これらのパッケージはGoプロジェクトによって作成されたものではなく、サポートされていませんが、役立つ場合があります。  
最新のバージョンに更新します。
```
sudo add-apt-repository ppa:longsleep/golang-backports
sudo apt update
```
インストールします。
```
sudo apt install golang
```
バージョン確認。
```
go version
```
# Projectの作成
1. プロジェクトフォルダを作成します。
2. `go mod init "フォルダ名"`を実行します。
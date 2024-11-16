# aconf

## 使用方法
- AWSアカウントの切り替え
```shell
aconf switch <profile>
```
- MFA認証の設定（armとプロファイルの結びつけ）
```shell
aconf add <profile> <arm> 
```
- MFA認証（`code`はワンタイムパスワード）
```shell
aconf <profile> <code>
```

## 初期設定
1. Goファイルをビルドする
```shell
cd cmd/aconf/
go build -o aconf
```
2. 所定の場所に移動する
```shell
sudo mv ./aconf /usr/local/bin/
```
3. シェルに権限付与&移動
プロジェクトルートで以下を行う.
```shell
sudo mv aconf.sh /usr/local/bin/aconf
chmod +x /usr/local/bin/aconf
```
4. シェルエイリアスを使った実行
`.zshrc`や`.bashrc`に以下を追加する
```bash
aconf() {
  local output
  output=$(/usr/local/bin/aconf "$@")
  if [ $? -ne 0 ]; then
    echo "$output"
    return 1
  fi
  eval "$output"
}
```
保存後, 再読み込み
```shell
source ~/.bashrc
```
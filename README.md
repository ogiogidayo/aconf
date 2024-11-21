# aconf
![Version](https://img.shields.io/github/v/tag/ogiogidayo/aconf?label=version&sort=semver)
![License](https://img.shields.io/github/license/ogiogidayo/aconf)
![workflow](https://github.com/ogiogidayo/aconf/actions/workflows/CI.yaml/badge.svg)

## Outline
- AWSアカウントをCLI上でMFA認証するときに以下のコードを簡単なコマンドで実行できるようにするツール
```shell
OUTPUT=$(aws sts get-session-token \
  --serial-number <arm> \
  --profile <profile> --token-code xxxxxx)
export AWS_ACCESS_KEY_ID=$(echo $OUTPUT | jq -r .Credentials.AccessKeyId)
export AWS_SECRET_ACCESS_KEY=$(echo $OUTPUT | jq -r .Credentials.SecretAccessKey)
export AWS_SESSION_TOKEN=$(echo $OUTPUT | jq -r .Credentials.SessionToken)
```

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

3. シェルエイリアスを使った実行
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
source ~/.zshrc
```

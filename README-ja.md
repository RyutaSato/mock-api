# 超超簡易DockerとファイルベースのMockサーバー

## 概要
`/responses`配下にJSONファイル(`*.json`)を置くだけで、JSONを返すMockサーバーが完成します。

## 使い方
リポジトリをクローンし、以下のコマンドでDockerを起動します。
```bash
cd mock-api
PORT=8080
```
`${PORT}`はListenしたいポート番号に置き換えるか、環境変数で指定します。
```bash
docker build -t mock-api .
docker run -d -p ${PORT}:8080 -v $(pwd)/responses:/responses mock-api
```

Windows上でPowerShellを使用している場合
```ps1
# cd mock-api
# $PORT=8080
$PWD="$((pwd).Path.Replace('\', '/'))"
docker build -t mock-api .
docker run -d -p ${PORT}:8080 -v ${PWD}/responses:/responses mock-api
```

`responses`配下はコンテナにマウントされるので、コンテナ起動中でもJSONファイルを変更することで続けてテストができます。

## 例 
- `http://localhost:8080/v2/index`をテストしたい場合  
`/responses`配下に`v2`ディレクトリを作成し、その配下に`index.json`を配置します。
`/responses/v2/index.json`を配置します。  

- `http://localhost:8080/v2`をテストしたい場合  
`/responses/v2.json`を配置します。  

## 指摘事項
- `*.json`ファイルしか認識しません。

# 超超簡易DockerとファイルベースのMockサーバー

## 概要
`/responses`配下にJSONファイル(`*.json`)を置くだけで、JSONを返すMockサーバーが完成します。
また、ファイル名を`*.post.json`などとすることで、GET以外のメソッドのレスポンスにも対応します。

## 制限事項
とってもシンプルな構成のため、以下のケースには現状対応していません。  
`JSON-server`や`postman`を検討してください。または、リポジトリをForkして実装してください。
- 複数パターンのレスポンス
- `200`以外のステータスコード
- BodyやHeaderに応じた条件分岐
- JSON以外のレスポンス


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
```powershell
# cd mock-api
# $PORT=8080
$PWD="$((pwd).Path.Replace('\', '/'))"
docker build -t mock-api .
docker run -d -p ${PORT}:8080 -v ${PWD}/responses:/responses mock-api
```

`responses`配下はコンテナにマウントされるので、コンテナ起動中でもJSONファイルを変更することで続けてテストができます。

## 例 
- `/v2/index`をテストしたい場合  
`/responses`配下に`v2`ディレクトリを作成し、その配下に`index.json`を配置します。
`/responses/v2/index.json`を配置します。  

- `/v2`をテストしたい場合  
実装方法は2つあります。
    - `/responses/v2.json`を配置します。  
    - `/responses/v2`配下に`.json`を配置します。

- `/v2?id=123&name=json`をテストしたい場合
`/responses`配下に`v2_id=123&name=json.json`を配置します。
> クエリパラメータは、`?`を`_`に置き換えて指定できます。(ファイル名に`?`が使用できないため)

- `/v2`をPOSTリクエストでテストしたい場合
`/responses`配下に`v2.post.json`を配置します。
> メソッドは、`(path).(method).json`の形式で指定できます。(`GET`のみ省略できます)


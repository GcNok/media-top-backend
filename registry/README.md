## registry
- DIコンテナ的なやつ
- 全てのオブジェクトの生成処理を行い依存の解決を行う
- `registry`の生成は`main`もしくは`interfaces`の中で行い、各`handler`に渡す
- /domain/repository配下にあるinterfaceを/infrastructure配下にあるrepositoryファイルに紐づける
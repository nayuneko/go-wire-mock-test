# Goでdi＆mockテスト

極力、標準寄りのライブラリを使ってやってみる

## di

Wire: https://github.com/google/wire

- servive,repositoryをinterfaceで定義
- di用に`wire.go`を作成
- `wire wire.go`で`wire_gen.go`を生成

## mockテスト

gomock: https://github.com/golang/mock

- `mockgen`でmockを生成
- テストからmockを呼び出してテストを書く

mockgenで生成される引数,戻り値の型が`interface{}`なので実行してみるまで結果がわからないのが不便…
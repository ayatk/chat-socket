# chat socket

WebSocketを使ってチャットっぽいチャットを実現するシステム

エラー処理も設計もないとりあえず動けばいいや的な考えで作ったクソコードだから許して欲しい

## 使い方
ビルド

    dep ensure

    go build

サーバー

    ./chat-socket server [-port 1234]


クライアント

    ./chat-socket client -host <hostname> [-port 1234] -name <your name>

## LICENSE
どうぞお好きなようにしやがれ ライセンス

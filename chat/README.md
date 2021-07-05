## memo
- 実行したいファイルが複数ある場合、適切にビルドしてバイナリを作成する必要がある
  - `go build -o ファイル名`
  - `go build -o .`
    - `./chap1`

- 他の言語でマルチスレッドや並行処理を行う際には多くのコードが必要になりますが、 Goではわずか2文字のキーワードを追加するだけで同等の機能を利用できます。システム開発者がGoを好む理由の1つがここにあります。

### websocket
```javascript
socket = new WebSocket("ws://localhost:8080/room");
socket.onclose = function() {
    alert("接続が終了しました");
}
socket.onmessage = function(e) {
    messages.append($("<li>").text(e.data));
}
```

### TDD (Test Driven Development)
- APIの設計は、自分がそのAPIのユーザーになることから始まります。
- red-green testing
  - 最初にユニットテストを作成し意図的に失敗させる
  - そのテストを成功させるために必要な最小限のコードを追加する
- 1o.Writerであれば何でも受け付けるということは、ユーザーが出力先を自由に選べるというこ とを意味します。
  - 例えば標準出力、ファイル、ネットワークのソケット、☆bytes.Buffe(テr ストケー スではこれを利用しています)、あるいは自作のオブジェクトでもかまいません。

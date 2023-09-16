# sample-todo-app

ginフレームワークの練習用に作ったTODOリストWebAppの雛形です。
めちゃくちゃしょぼいのでいい感じに改良してあげてください。

## 開発環境

* macOS Ventura 13.4.1
* Go 1.20.5
* Gin 1.9.1
* PostgreSQL 14.7（作者の環境）

## 実行方法

```zsh
go run main.go
```

## ご利用にあたって

### 1.リモートリポジトリは変更禁止

利用される際はgitの管轄から外すorご自身のリポジトリと紐づけてください。

### 2.社用等の公的PCでの利用禁止

万一トラブルが発生した際の保証は一切致しかねますので、必ず私用のPCで利用されるようお願いいたします。

### 3.DBについて

DB環境の構築はご自身で行なっていただくようお願いいたします。
今回はPostgreSQLを使用していますが、MySQLでも動作することを確認しています。

## とりあえずやった方がいいこと

タスクの追加機能と削除機能は最低限ないと辛いかなと思います。削除機能はハリボテのボタンのみ置いてあるのでカスタマイズして実装してください。
他にも進行状況のステータスや、ソート機能、実施期限なんかつけると面白いかもしれないですね。


ご不明な点がございましたら https://twitter.com/negimaboy_1293?s=21&t=NkyOf4Ku3wzBvOhu4KZQJw までご連絡ください。
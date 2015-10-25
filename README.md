# GoPrincessPrecure

## Install
```
$ go get github.com/sue445/go_princess_precure
```

## Usage
via .[gore](https://github.com/motemen/gore)

```go
gore> :import github.com/sue445/go_princess_precure
gore> flora := go_princess_precure.NewCureFlora()
gore> flora.Name()
"春野はるか"
gore> flora.Transform()
プリキュア！プリンセスエンゲージ！
咲き誇る花のプリンセス！キュアフローラ！
強く、やさしく、美しく！
Go!プリンセスプリキュア！
冷たい檻に閉ざされた夢、返していただきますわ！
お覚悟はよろしくて？
gore> flora.Name()
"キュアフローラ"
gore> flora.Exchange(go_princess_precure.TransformFlora)
エクスチェンジ！モードエレガント！
舞え、花よ！プリキュア・フローラル・トルビヨン！
ごきげんよう
```

## Development

```sh
$ go get github.com/mattn/gom
$ gom install

$ export GOPATH="`pwd`/_vendor":$GOPATH
$ _vendor/bin/gore
```

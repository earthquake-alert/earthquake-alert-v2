package main

import "github.com/earthquake-alert/erarthquake-alert-v2/src"

// この変数はビルド時に上書きできるようになっています。
// デフォルトは`local`であるため、上書きをしないでビルドするとローカル環境の設定などが適用されます。
// ビルド時に設定する:
//
//	go build  -ldflags="-X main.mode=prod"
//
// 設定できる項目は`test`, `local`, `prod`です。
var mode = "local"

// 初期設定を行います。
// 設定項目の出し分け、ログの初期化など…
func init() {
	src.Init(mode)
}

func main() {
	src.Server()
}

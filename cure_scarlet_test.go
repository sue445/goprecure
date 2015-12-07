package goprecure

import (
	"fmt"
)

func ExampleCureScarlet() {
	scarlet := NewCureScarlet()
	scarlet.PrintInterval = 0

	// human
	fmt.Println("Name", scarlet.Name())

	scarlet.Transform()

	// precure
	fmt.Println("Name", scarlet.Name())

	fmt.Println(Scarlet.String())
	scarlet.Exchange(Scarlet)

	fmt.Println(Hanabi.String())
	scarlet.Exchange(Hanabi)

	fmt.Println(Phoenix.String())
	scarlet.Exchange(Phoenix)

	fmt.Println(Sun.String())
	scarlet.Exchange(Sun)

	// Output:
	// Name 紅城トワ
	// プリキュア！プリンセスエンゲージ！
	// 深紅の炎のプリンセス！キュアスカーレット！
	// 強く、やさしく、美しく！
	// Go!プリンセスプリキュア！
	// 冷たい檻に閉ざされた夢、返していただきますわ！
	// お覚悟決めなさい！
	// Name キュアスカーレット
	// Scarlet
	// エクスチェンジ！モードエレガント！
	// たぎれ、炎よ！プリキュア・スカーレット・フレイム！
	// ごきげんよう
	// Hanabi
	// 花火！
	// 燃えよ、炎よ！プリキュア・スカーレット・スパーク！
	// ごきげんよう
	// Phoenix
	// スカーレットバイオリン！フェニックス！
	// 羽ばたけ炎の翼！プリキュア・フェニックス・ブレイズ！
	// ごきげんよう
	// Sun
	// モードエレガント！
	// 桜！珊瑚！銀河！サン！
	// ドレスアッププレミアム！
	// はぁッ！
	// 響け全ての力！プリキュアエクラエスポワール！
	// ごきげんよう
}

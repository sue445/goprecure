package goprecure

import (
	"fmt"
)

func ExampleCureMermaid() {
	mermaid := NewCureMermaid()
	mermaid.PrintInterval = 0

	// human
	fmt.Println("Name", mermaid.Name())

	mermaid.Transform()

	// precure
	fmt.Println("Name", mermaid.Name())

	fmt.Println(Mermaid.String())
	mermaid.Exchange(Mermaid)

	fmt.Println(Ice.String())
	mermaid.Exchange(Ice)

	fmt.Println(Bubble.String())
	mermaid.Exchange(Bubble)

	fmt.Println(Sango.String())
	mermaid.Exchange(Sango)

	// Output:
	// Name 海藤みなみ
	// プリキュア！プリンセスエンゲージ！
	// 澄み渡る海のプリンセス！キュアマーメイド！
	// 強く、やさしく、美しく！
	// Go!プリンセスプリキュア！
	// 冷たい檻に閉ざされた夢、返していただきますわ！
	// お覚悟はよろしくて？
	// Name キュアマーメイド
	// Mermaid
	// エクスチェンジ！モードエレガント！
	// 高鳴れ、海よ！プリキュア・マーメイド・リップル！
	// ごきげんよう
	// Ice
	// アイス！
	// 高鳴れ、氷よ！プリキュア・フローズン・リップル！
	// ごきげんよう
	// Bubble
	// バブル！
	// 高まれ、泡よ！プリキュア・バブルリップル！
	// ごきげんよう
	// Sango
	// モードエレガント！
	// 桜！珊瑚！銀河！サン！
	// ドレスアッププレミアム！
	// はぁッ！
	// 響け全ての力！プリキュアエクラエスポワール！
	// ごきげんよう
}

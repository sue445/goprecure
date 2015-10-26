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

	fmt.Println(TransformMermaid.String())
	mermaid.Exchange(TransformMermaid)

	fmt.Println(Ice.String())
	mermaid.Exchange(Ice)

	fmt.Println(Bobble.String())
	mermaid.Exchange(Bobble)

	fmt.Println(Coral.String())
	mermaid.Exchange(Coral)

	// Output:
	// Name 海藤みなみ
	// プリキュア！プリンセスエンゲージ！
	// 澄み渡る海のプリンセス！キュアマーメイド！
	// 強く、やさしく、美しく！
	// Go!プリンセスプリキュア！
	// 冷たい檻に閉ざされた夢、返していただきますわ！
	// お覚悟はよろしくて？
	// Name キュアマーメイド
	// TransformMermaid
	// エクスチェンジ！モードエレガント！
	// 高鳴れ、海よ！プリキュア・マーメイド・リップル！
	// ごきげんよう
	// Ice
	// アイス！
	// 高鳴れ、氷よ！プリキュア・フローズン・リップル！
	// ごきげんよう
	// Bobble
	// バブル！
	// 高まれ、泡よ！プリキュア・バブルリップル！
	// ごきげんよう
	// Coral
	// モードエレガント！
	// 桜！珊瑚！銀河！サン！
	// ドレスアッププレミアム！
	// はぁッ！
	// 響け全ての力！プリキュアエクラエスポワール！
}

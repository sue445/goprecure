package go_princess_precure

import (
	"fmt"
)

type DressupKey int

const (
	TransformFlora DressupKey = iota
	TransformMermaid
	TransformTwinkle
	TransformScarlett
)

func exchangeMessage(k DressupKey) string {
	switch k {
	case TransformFlora:
		return `エクスチェンジ！モードエレガント！
舞え、花よ！プリキュア・フローラル・トルビヨン！
ごきげんよう`

	case TransformMermaid:
		return `エクスチェンジ！モードエレガント！
高鳴れ、海よ！プリキュア・マーメイド・リップル！
ごきげんよう`

	case TransformTwinkle:
		return `エクスチェンジ！モードエレガント！
キラキラ、星よ！プリキュア・トゥインクル・ハミング！
ごきげんよう`

	}
	panic(fmt.Sprintf("Unknown DressUpKey: %d", k))
}

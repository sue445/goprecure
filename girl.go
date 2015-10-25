package go_princess_precure

import (
	"fmt"
	"strings"
	"time"
)

const (
	DefaultIntervalSec = 1
)

type Girl struct {
	girl_name          string
	human_name         string
	precure_name       string
	cast_name          string
	color              string
	created_date       string
	current_state      int
	transform_interval time.Duration
	transform_message  string
	dressup_keys       []DressupKey
}

func (g *Girl) Name() string {
	switch g.current_state {
	case 0:
		return g.human_name
	case 1:
		return g.precure_name
	}
	return g.human_name
}

func (g *Girl) Transform() {
	g.printByLine(g.transform_message)

	g.current_state++

	if g.current_state > 1 {
		g.current_state = 0
	}
}

func (g *Girl) printByLine(message string) {
	for _, line := range strings.Split(message, "\n") {
		line = strings.TrimSpace(line)
		fmt.Println(line)
		time.Sleep(g.transform_interval)
	}
}

func (g *Girl) Humanize() {
	g.current_state = 0
}

func (g *Girl) Exchange(k DressupKey) {
	if g.current_state == 0 {
		panic("Human can not exchange!")
	}

	if !g.CanUseKey(k) {
		panic(g.precure_name + " can not use " + k.String())
	}

	g.current_state = 2

	var message = exchangeMessage(k)
	g.printByLine(message)
}

func (g *Girl) CanUseKey(k DressupKey) bool {
	for _, value := range g.dressup_keys {
		if value == k {
			return true
		}
	}
	return false
}

func newGirl() *Girl {
	g := new(Girl)
	g.current_state = 0
	g.transform_interval = DefaultIntervalSec * time.Second
	return g
}

func NewCureFlora() *Girl {
	g := newGirl()
	g.girl_name = "cure_flora"
	g.human_name = "春野はるか"
	g.precure_name = "キュアフローラ"
	g.cast_name = "嶋村侑"
	g.color = "pink"
	g.created_date = "2015-02-01"
	g.dressup_keys = []DressupKey{TransformFlora, Rose, Lily, Blossom}

	g.transform_message = `プリキュア！プリンセスエンゲージ！
咲き誇る花のプリンセス！キュアフローラ！
強く、やさしく、美しく！
Go!プリンセスプリキュア！
冷たい檻に閉ざされた夢、返していただきますわ！
お覚悟はよろしくて？`

	return g
}

func NewCureMermaid() *Girl {
	g := newGirl()
	g.girl_name = "cure_mermaid"
	g.human_name = "海藤みなみ"
	g.precure_name = "キュアマーメイド"
	g.cast_name = "浅野真澄"
	g.color = "blue"
	g.created_date = "2015-02-08"
	g.dressup_keys = []DressupKey{TransformMermaid, Ice, Bobble, Coral}

	g.transform_message = `プリキュア！プリンセスエンゲージ！
澄み渡る海のプリンセス！キュアマーメイド！
強く、やさしく、美しく！
Go!プリンセスプリキュア！
冷たい檻に閉ざされた夢、返していただきますわ！
お覚悟はよろしくて？`

	return g
}

func NewCureTwinkle() *Girl {
	g := newGirl()
	g.girl_name = "cure_twinkle"
	g.human_name = "天ノ川きらら"
	g.precure_name = "キュアトゥインクル"
	g.cast_name = "山村響"
	g.color = "yellow"
	g.created_date = "2015-02-22"
	g.dressup_keys = []DressupKey{TransformTwinkle, Luna, ShootingStar, Galaxy}

	g.transform_message = `プリキュア！プリンセスエンゲージ！
きらめく星のプリンセス！キュアトゥインクル！
強く、やさしく、美しく！
Go!プリンセスプリキュア！
冷たい檻に閉ざされた夢、返していただきますわ！
お覚悟はよろしくて？`

	return g
}

func NewCureScarlett() *Girl {
	g := newGirl()
	g.girl_name = "cure_scarlett"
	g.human_name = "紅城トワ"
	g.precure_name = "キュアスカーレット"
	g.cast_name = "沢城みゆき"
	g.color = "red"
	g.created_date = "2015-07-05"
	g.dressup_keys = []DressupKey{TransformScarlett, Fireworks, Phoenix, Sun}

	g.transform_message = `プリキュア！プリンセスエンゲージ！
深紅の炎のプリンセス！キュアスカーレット！
強く、やさしく、美しく！
Go!プリンセスプリキュア！
冷たい檻に閉ざされた夢、返していただきますわ！
お覚悟決めなさい！`

	return g
}

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
	transform_messages map[string]string
}

func (g *Girl) name() string {
	switch g.current_state {
	case 0:
		return g.human_name
	case 1:
		return g.precure_name
	}
	return g.human_name
}

func (g *Girl) transform() {
	g.printByLine(g.transform_messages["normal"])

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

func (g *Girl) humanize() {
	g.current_state = 0
}

func newCureFlora() *Girl {
	g := new(Girl)
	g.girl_name = "cure_flora"
	g.human_name = "春野はるか"
	g.precure_name = "キュアフローラ"
	g.cast_name = "嶋村侑"
	g.color = "pink"
	g.created_date = "2015-02-01"
	g.current_state = 0
	g.transform_interval = DefaultIntervalSec * time.Second

	g.transform_messages = map[string]string{
		"normal": `プリキュア！プリンセスエンゲージ！
咲き誇る花のプリンセス！キュアフローラ！
強く、やさしく、美しく！
Go!プリンセスプリキュア！
冷たい檻に閉ざされた夢、返していただきますわ！
お覚悟はよろしくて？`,
		"exchange": `エクスチェンジ！モードエレガント！`,
	}

	return g
}

func newCureMermaid() *Girl {
	g := new(Girl)
	g.girl_name = "cure_mermaid"
	g.human_name = "海藤みなみ"
	g.precure_name = "キュアマーメイド"
	g.cast_name = "浅野真澄"
	g.color = "blue"
	g.created_date = "2015-02-08"
	g.current_state = 0
	g.transform_interval = DefaultIntervalSec * time.Second

	g.transform_messages = map[string]string{
		"normal": `プリキュア！プリンセスエンゲージ！
澄み渡る海のプリンセス！キュアマーメイド！
強く、やさしく、美しく！
Go!プリンセスプリキュア！
冷たい檻に閉ざされた夢、返していただきますわ！
お覚悟はよろしくて？`,

		"exchange": `エクスチェンジ！モードエレガント！`,
	}

	return g
}

func newCureTwinkle() *Girl {
	g := new(Girl)
	g.girl_name = "cure_twinkle"
	g.human_name = "天ノ川きらら"
	g.precure_name = "キュアトゥインクル"
	g.cast_name = "山村響"
	g.color = "yellow"
	g.created_date = "2015-02-22"
	g.current_state = 0
	g.transform_interval = DefaultIntervalSec * time.Second

	g.transform_messages = map[string]string{
		"normal": `プリキュア！プリンセスエンゲージ！
きらめく星のプリンセス！キュアトゥインクル！
強く、やさしく、美しく！
Go!プリンセスプリキュア！
冷たい檻に閉ざされた夢、返していただきますわ！
お覚悟はよろしくて？`,

		"exchange": `エクスチェンジ！モードエレガント！`,
	}

	return g
}

func newCureScarlett() *Girl {
	g := new(Girl)
	g.girl_name = "cure_scarlett"
	g.human_name = "紅城トワ"
	g.precure_name = "キュアスカーレット"
	g.cast_name = "沢城みゆき"
	g.color = "red"
	g.created_date = "2015-07-05"
	g.current_state = 0
	g.transform_interval = DefaultIntervalSec * time.Second

	g.transform_messages = map[string]string{
		"normal": `プリキュア！プリンセスエンゲージ！
深紅の炎のプリンセス！キュアスカーレット！
強く、やさしく、美しく！
Go!プリンセスプリキュア！
冷たい檻に閉ざされた夢、返していただきますわ！
お覚悟決めなさい！`,

		"exchange": `エクスチェンジ！モードエレガント！`,
	}

	return g
}

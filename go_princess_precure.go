package main
import "fmt"

type Girl struct {
	girl_name    string
	human_name   string
	precure_name string
	cast_name    string
	color        string
	created_date string
}

func newCureFlora() *Girl {
	girl := new(Girl)
	girl.girl_name = "cure_flora"
	girl.human_name = "春野はるか"
	girl.precure_name = "キュアフローラ"
	girl.cast_name = "嶋村侑"
	girl.color = "pink"
	girl.created_date = "2015-02-01"
	return girl
}

func main() {
	flora := newCureFlora()
	fmt.Println(flora)
}

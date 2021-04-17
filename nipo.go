package nipo

import "fmt"

var templateMac = `
●本日の業務記録
10:00 ~ 11:00 := (Mac)
11:00 ~ 12:00 := (Mac)
13:00 ~ 14:00 := (Mac)
14:00 ~ 15:00 := (Mac)
15:00 ~ 16:00 := (Mac)
16:00 ~ 17:00 := (Mac)
17:00 ~ 18:00 := (Mac)
18:00 ~ 19:00 := (Mac)
`

var templateVM = `
●本日の業務記録
10:00 ~ 11:00 := (仮想)
11:00 ~ 12:00 := (仮想)
13:00 ~ 14:00 := (仮想)
14:00 ~ 15:00 := (仮想)
15:00 ~ 16:00 := (仮想)
16:00 ~ 17:00 := (仮想)
17:00 ~ 18:00 := (仮想)
18:00 ~ 19:00 := (仮想)
`

// PrintTemplate 日報のテンプレートを出力します
func PrintTemplate(pattern string) {
	switch pattern {
	case "mac":
		fmt.Println(templateMac)
	case "vm":
		fmt.Println(templateVM)
	default:
		fmt.Println("引数に渡せるのは `mac`と`vm`(仮想)だけだよ！")
	}
}


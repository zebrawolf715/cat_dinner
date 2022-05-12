package main

import (
	"fmt"
)

func main() {

	//挨拶文
	fmt.Println("注文の多い料理店へようこそ！")
	totalMoney := 3000 // 初期の手持ちのお金3,000円をセット
	totalMoney2 := totalMoney
	fmt.Printf("あなたは%d円を手元に持っています。\n", totalMoney)
	fmt.Println("手持ちのお金がなくならない程度にレストランで食事をしよう")
	fmt.Printf("手持ちのお金がなくなったらレストランから出られないよ、気をつけて！\n\n")

	//追加
	eatAsk(&totalMoney2)

	//合計を出力
	//fmt.Printf("%d円\n", totalMoney)

	//脱出出来るかどうか
	if totalMoney2 <= 0 {
		println("あなたは注文の多い料理店から出られません")
		println("GAME OVER")
	} else {
		//食べた場合と食べてない場合の分岐
		if totalMoney == totalMoney2 {
			println("あなたは何も食べずに無事帰りました")
		} else {
			totalEat := totalMoney - totalMoney2
			fmt.Printf("あなたは猫達の隙を付いて%d円分食べることができました\n", totalEat)
			println("あなたは無事帰りました")
		}
	}
}

func eatAsk(totalM *int) {

	//メニューセット
	foods := map[string]int{
		"トマトスープ":    800,
		"わかめスープ":    700,
		"ステーキ":      1800,
		"ショートケーキ":   500,
		"チョコレートケーキ": 600,
		"サラダ":       500,
		"魚":         1600,
	}

	eatPrice := 3000

	//食べるか食べないかの分岐
	var eatStart string
	for key, value := range foods {
		//食べるかユーザに聞く
		fmt.Println("あなたは1品ランダムに食べますか？Y/N")
		fmt.Scan(&eatStart)
		if eatStart == "Y" {
			fmt.Printf("あなたは%sを食べました\n", key)
			fmt.Printf("%sは%dでした\n\n", key, value)
			eatPrice = eatPrice - value
		} else if eatStart == "N" {
			fmt.Printf("あなたは帰ることにしました\n\n")
			*totalM = eatPrice
			return
		} else {
			continue
		}
	}

}

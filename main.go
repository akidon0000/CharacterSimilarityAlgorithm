package main

import (
	"fmt"
	"math"
	"strconv"
	"unicode/utf8"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/masatana/go-textdistance"
)

// ===============
const excelPath = "./matching.xlsx"
const excelSheet = "Sheet1"

// ===============

func main() {
	fmt.Println("\n", "\n", "---相性診断開始---\n", "Ver,0.1.2")
	f, err := excelize.OpenFile(excelPath)
	if err != nil {
		fmt.Println(err)
		return
	}

	rows := f.GetRows(excelSheet)
	for a, row := range rows {
		if a == 0 {
			continue
		}
		h, i, j, k, l := algorithm(row[0], row[1], row[2], row[3], row[4], row[5], row[6])

		f.SetCellValue(excelSheet, "H"+strconv.Itoa(a+1), h)
		f.SetCellValue(excelSheet, "I"+strconv.Itoa(a+1), i)
		f.SetCellValue(excelSheet, "J"+strconv.Itoa(a+1), j)
		f.SetCellValue(excelSheet, "K"+strconv.Itoa(a+1), k)
		f.SetCellValue(excelSheet, "L"+strconv.Itoa(a+1), l)
	}
	f.SaveAs(excelPath)
	fmt.Println("\n", "\n", "---終了---\n")
}

func algorithm(my_my string, my_par string, my_qua string, par_my string, par_par string, par_qua string, z string) (float64, float64, float64, float64, float64) {

	rawAffinity := 0.0 // [補正前]相性度
	affinity := 0.0    // [補正後]相性度
	matchLevel := 10   // 補正強度 高いほど相性度が下がる
	correction1 := 0.0 // 補正値1
	correction2 := 0.0 // 補正値2
	correction3 := 0.0 // 補正値3 quadkey距離
	quadLevel, err := strconv.Atoi(z)
	if err != nil {
		fmt.Println("Atoi zoomLevel ERROR")
	}

	// 範囲外なら結果は0
	if my_qua[0:quadLevel] != par_qua[0:quadLevel] {
		return 0, 0, 0, 0, 0
	}

	// 文字数
	a_Len := utf8.RuneCountInString(my_my)
	b_Len := utf8.RuneCountInString(my_par)
	c_Len := utf8.RuneCountInString(par_my)
	d_Len := utf8.RuneCountInString(par_par)
	qua_1Len := utf8.RuneCountInString(my_qua)
	qua_2Len := utf8.RuneCountInString(par_qua)

	// 0-100の値をとる
	rawAffinity += textdistance.JaroWinklerDistance(my_my, par_par) * 100
	rawAffinity += textdistance.JaroWinklerDistance(my_par, par_my) * 100
	rawAffinity /= 2

	// それぞれ文字数が多いほど，文字数が一致してるほど補正値は高い (プラス補正)
	// correction1 += float64((b_Len + c_Len) + (b_Len - c_Len) + (a_Len + d_Len) + (a_Len - d_Len))
	correction1 += float64(( a_Len + b_Len + c_Len + d_Len) / 4)

	// 何文字異なっているか × matchLevel (マイナス補正)
	correction2 -= float64(textdistance.LevenshteinDistance(my_my, par_par) * matchLevel)
	correction2 -= float64(textdistance.LevenshteinDistance(my_par, par_my) * matchLevel)

	// どれほど距離が近いか (プラス補正)
	a, err := strconv.Atoi(my_qua[quadLevel:qua_1Len])
	if err != nil {
		fmt.Println("Atoi a ERROR")
	}
	b, err := strconv.Atoi(par_qua[quadLevel:qua_2Len])
	if err != nil {
		fmt.Println("Atoi b ERROR")
	}
	// Quadkeyの差分
	c := strconv.Itoa(int(math.Abs(float64(a) - float64(b))))
	correction3 -= float64(utf8.RuneCountInString(c)) * 1.5
	// 差分が0の時を差し引く
	if c == "0" {
		correction3 = 0
	}

	affinity = rawAffinity

	affinity += correction1
	affinity += correction2
	if affinity > 100 {
		affinity = 100
	}
	affinity += correction3
	if 0 > affinity {
		affinity = 0
	}

	return affinity, correction1, correction2, correction3, rawAffinity
}

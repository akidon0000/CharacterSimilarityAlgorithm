package main

import (
	// "bufio"
	"fmt"
	// "os"
	"strconv"
	"math"

	"github.com/360EntSecGroup-Skylar/excelize"
)

// ===============
const excelPath = "./testData.xlsx"
const excelSheet = "Sheet1"
// const celNum = 12

// ===============

func main() {
	f, err := excelize.OpenFile(excelPath)
	if err != nil {
		fmt.Println(err)
		return
	}

	rows := f.GetRows(excelSheet)
	for i, row := range rows {
		str := lss(row[0], row[1])



		f.SetCellValue(excelSheet, "C"+strconv.Itoa(i+1), str)
	}
	f.SaveAs(excelPath)
}

func ld(s string,t string) float64{
	if s == ""{
		return float64(len(t))
	}
	if t == ""{
		return float64(len(s))
	}
	if s[0] == t[0]{
		return ld(s[1:], t[1:])
	}
	l1 := ld(s, t[1:])
	l2 := ld(s[1:], t)
	l3 := ld(s[1:], t[1:])
	return 1 + math.Min(l1, math.Min(l2, l3))
}

func lds(s string, t string) float64{
    return ld(s, t) / math.Max(float64(len(s)), float64(len(t)))
}

func lss(s string, t string) float64{
    return -lds(s, t) + 1
}

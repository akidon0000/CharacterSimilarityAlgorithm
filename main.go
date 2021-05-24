package main

import (
	"fmt"
	"strconv"
	"time"
  "math/rand"

	"github.com/360EntSecGroup-Skylar/excelize"

	// "github.com/masatana/go-textdistance"
	// "github.com/toldjuuso/go-jaro-winkler-distance"
)

// ===============
const excelPath = "./testData.xlsx"
const excelSheet = "Sheet1"
const testNum = 500000
// const quadkey1 = 123123123012312312
// const quadkey2 = 123123123012122312

// ===============

func main(){
	f, err := excelize.OpenFile(excelPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	for i := 0; i < testNum; i++ {
		fmt.Println(i+1)
		rand.Seed(time.Now().UnixNano())
		var a int
		var b int
    for i := 0; i < 10; i++ {
			a = int(rand.Intn(8)+1)
			b = int(rand.Intn(8)+1)
    }
		random1, _ := MakeRandomStr(a)
		random2, _ := MakeRandomStr(b)
		f.SetCellValue(excelSheet, "A"+strconv.Itoa(i+1), "0")
		f.SetCellValue(excelSheet, "B"+strconv.Itoa(i+1), "test20210524-" + strconv.Itoa(i+1))
		f.SetCellValue(excelSheet, "C"+strconv.Itoa(i+1), random1)
		f.SetCellValue(excelSheet, "D"+strconv.Itoa(i+1), random2)
		f.SetCellValue(excelSheet, "E"+strconv.Itoa(i+1), "0133002112310210000")
		f.SetCellValue(excelSheet, "F"+strconv.Itoa(i+1), "NULL")
	}
	f.SaveAs(excelPath)
}

func MakeRandomStr(digit int) (string, error) {
    const letters = "あいうえおかきくけこさしすせそたちつてとなにぬねのまみむめもはひふへほやゆよらりるれろわをんんんんあいうえおかきくけこ"
		rand.Seed(time.Now().UnixNano())

		var a int
    var result string
		result = ""
    for i := 0; i < digit; i++ {
			for i := 0; i < 50; i++ {
				a = rand.Intn(50)
			}
			rs := []rune(letters)
			result += string(rs[a:a+1])
    }
    return result, nil
}






// func main() {
// 	fmt.Println("====")
// 	str := "aaaa"
// 	fmt.Println(utf8.RuneCountInString(str))
// 	for i := 0; i < utf8.RuneCountInString(str); i++ {
// 		s := str[:len(str)/utf8.RuneCountInString(str)]
// 		fmt.Println(s)
// 		for j := i; j < utf8.RuneCountInString(str)-1; j++ {
// 			s += "_"
// 		}
// 		fmt.Println("s:")
// 		fmt.Println(s)
// 	}

// 	fmt.Println("==開始==")
// 	f, err := excelize.OpenFile(excelPath)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	rows := f.GetRows(excelSheet)
// 	for i, row := range rows {
// 		fmt.Print(i)
// 		fmt.Println("==============")
// 		fmt.Println(textdistance.LevenshteinDistance(row[0], row[1]))
// 		fmt.Println(textdistance.DamerauLevenshteinDistance(row[0], row[1]))
// 		fmt.Println(textdistance.JaroDistance(row[0], row[1]))
// 		fmt.Println(textdistance.JaroWinklerDistance(row[0], row[1]))
// 		// str := lss(row[0], row[1])
// 		// str := textdistance.JaroWinklerDistance
// 		fmt.Println(jwd.Calculate(row[0], row[1]))
// 		str := jwd.Calculate(row[0], row[1])
// 		fmt.Println(row[0] + " " + row[1] + " : " + strconv.FormatFloat(str, 'f', 2, 64))
// 		num := 10 - utf8.RuneCountInString(strconv.Itoa(quadkey1-quadkey2))
// 		fmt.Println("---")
// 		fmt.Println(num)
// 		f.SetCellValue(excelSheet, "C"+strconv.Itoa(i+1), str)
// 	}
// 	f.SaveAs(excelPath)
// 	fmt.Println("==完了しました==")
// }

// func ld(s string, t string) float64 {
// 	if s == "" {
// 		return float64(utf8.RuneCountInString(t))
// 	}
// 	if t == "" {
// 		return float64(utf8.RuneCountInString(s))
// 	}
// 	if s[0] == t[0] {
// 		return ld(s[1:], t[1:])
// 	}
// 	l1 := ld(s, t[1:])
// 	l2 := ld(s[1:], t)
// 	l3 := ld(s[1:], t[1:])
// 	return 1 + math.Min(l1, math.Min(l2, l3))
// }

// func lds(s string, t string) float64 {
// 	fmt.Println(ld(s, t))
// 	fmt.Println(math.Max(float64(utf8.RuneCountInString(s)), float64(utf8.RuneCountInString(t))))
// 	return ld(s, t) / math.Max(float64(utf8.RuneCountInString(s)), float64(utf8.RuneCountInString(t)))
// 	// return ld(s, t)
// }

// func lss(s string, t string) float64 {
// 	return -lds(s, t) + 1
// }

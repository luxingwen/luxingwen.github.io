package main

import (
	"fmt"
)

var (
	tianGan = []string{"癸", "甲", "乙", "丙", "丁", "戊", "己", "庚", "辛", "壬"}
	dizhi   = []string{"亥", "子", "丑", "寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌"}
	dzMonth = []string{"丑", "寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌", "亥", "子"}
)

type Tgdz struct {
	Year    int
	Month   int
	Day     int
	T       int
	TdYear  TianGanDizhi
	TdMonth TianGanDizhi
	TdDay   TianGanDizhi
	TdT     TianGanDizhi
}

type TianGanDizhi struct {
	TianGanNum int
	DizhiNum   int
	TianGanStr string
	DizhiStr   string
}

func (this *Tgdz) tgdz() {
	this.getYearTianGanDizhi()
	this.getMonthGan()
	this.getDayGan()
	this.getTTianGanDizhi()
}

func (this *Tgdz) getYearTianGanDizhi() {
	year := this.Year
	this.TdYear.TianGanStr = tianGan[(year-3)%10]
	this.TdYear.DizhiStr = dizhi[(year-3)%12]
	this.TdYear.TianGanNum = (year - 3) % 10
	if this.TdYear.TianGanNum == 0 {
		this.TdYear.TianGanNum = 10
	}
	this.TdYear.DizhiNum = (year - 3) % 12
	if this.TdYear.DizhiNum == 0 {
		this.TdYear.DizhiNum = 12
	}
	return
}

func (this *Tgdz) getMonthGan() {
	year := this.Year
	month := this.Month
	yearGan := (year - 3) % 10
	if yearGan == 0 {
		yearGan = 10
	}
	this.TdMonth.TianGanNum = ((yearGan*2 + month) - 1) % 10
	if this.TdMonth.TianGanNum == 0 {
		this.TdMonth.TianGanNum = 10
	}
	oldMonth := month
	this.TdMonth.DizhiNum = (oldMonth - 1) % 12
	if this.TdMonth.DizhiNum == 0 {
		this.TdMonth.DizhiNum = 12
	}

	this.TdMonth.TianGanStr = tianGan[((yearGan*2+month)-1)%10]

	this.TdMonth.DizhiStr = dzMonth[(oldMonth-1)%12]
}

func (this *Tgdz) getDayGan() (tiangan string, dz string) {
	year := this.Year
	month := this.Month
	day := this.Day

	C := year / 100
	y := year % 100
	if month == 1 || month == 2 {
		y -= 1
	}
	M := month
	if month == 1 || month == 2 {
		M += 12
	}
	i := 0
	if month%2 == 0 {
		i = 6
	}
	G := 4*C + C/4 + 5*y + y/4 + 3*(M+1)/5 + day - 3
	G = G % 10
	this.TdDay.TianGanNum = G
	if this.TdDay.TianGanNum == 0 {
		this.TdDay.TianGanNum = 10
	}
	Z := 8*C + C/4 + 5*y + y/4 + 3*(M+1)/5 + day + 7 + i
	Z = Z % 12

	this.TdDay.DizhiNum = Z
	if this.TdDay.DizhiNum == 0 {
		this.TdDay.DizhiNum = 12
	}
	this.TdDay.TianGanStr = tianGan[G]
	this.TdDay.DizhiStr = dizhi[Z]
	return
}

func (this *Tgdz) getTTianGanDizhi() {
	this.TdT.DizhiNum, this.TdT.DizhiStr = getTimeZhi(this.T)
	this.TdT.TianGanNum = (this.TdDay.TianGanNum*2 + this.TdT.DizhiNum - 2) % 10
	this.TdT.TianGanStr = tianGan[this.TdT.TianGanNum]
	if this.TdT.TianGanNum == 0 {
		this.TdT.TianGanNum = 10
	}
}

func (t *Tgdz) Print() {
	fmt.Printf("%d-%d-%d : %d 点 => %s %s 年, %s %s 月, %s %s 日, %s %s 时\n", t.Year, t.Month, t.Day, t.T, t.TdYear.TianGanStr, t.TdYear.DizhiStr,
		t.TdMonth.TianGanStr, t.TdMonth.DizhiStr, t.TdDay.TianGanStr, t.TdDay.DizhiStr, t.TdT.TianGanStr, t.TdT.DizhiStr)
}

func getTimeZhi(t int) (zindex int, dz string) {
	switch t {
	case 23, 0, 24:
		return 1, dizhi[1]
	case 1, 2:
		return 2, dizhi[2]
	case 3, 4:
		return 3, dizhi[3]
	case 5, 6:
		return 4, dizhi[4]
	case 7, 8:
		return 5, dizhi[5]
	case 9, 10:
		return 6, dizhi[6]
	case 11, 12:
		return 7, dizhi[7]
	case 13, 14:
		return 8, dizhi[8]
	case 15, 16:
		return 9, dizhi[9]
	case 17, 18:
		return 10, dizhi[10]
	case 19, 20:
		return 11, dizhi[11]
	case 21, 22:
		return 12, dizhi[0]
	}
	return
}

func main() {

	t := new(Tgdz)
	t.Year = 2019
	t.Month = 9
	for di := 25; di < 30; di++ {
		fmt.Println("--------------------")
		t.Day = di
		for i := 0; i < 24; i++ {
			t.T = i
			t.tgdz()
			t.Print()
		}
		fmt.Println("<<<<<<<<<<<<<<<<<<<<<<<")
	}

}

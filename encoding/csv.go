package main

import (
	"encoding/csv"
	"os"
)

func main() {
	//r := strings.NewReader(`"1","气质美女芦奇性感内衣秀完美身材","65","17/5/2018 20:57:53","1"`)

	//csvR := csv.NewReader(r)
	//records, _ := csvR.Read()
	//fmt.Println(records)

	//f, _ := os.Open("user.csv")

	//csvR = csv.NewReader(f)
	//fmt.Println(csvR.ReadAll())

	w := csv.NewWriter(os.Stdout)
	//w.Write([]string{"6","极品嫩妹考拉丰乳细腰吸引你眼球","45","16/5/2018 23:45:55","1"})
	//w.Flush()

	w.WriteAll([][]string{{"6","极品嫩妹考拉丰乳细腰吸引你眼球","45","16/5/2018 23:45:55","1"},{"7","巨乳女神尤妮丝惹火曲线勾人魂魄","48","15/5/2018 20:30:27","1"}})
	w.Flush()

}
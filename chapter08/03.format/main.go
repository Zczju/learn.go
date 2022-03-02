package main

import (
	"encoding/json"
	"fmt"
	gobmi "github.com/armstrongli/go-bmi"
	"io/ioutil"
	"learn.go/pkg/apis"
	"log"
	"os"
)

func main() {
	input := &inputFromStd{}
	records := NewRecord("C:\\Users\\Administrator\\Desktop\\records.txt")
	c := &Calc{}
	rk := &FatRateRank{}

	for {
		// 录入
		personInfo := input.GetInput()

		// 保存用户信息
		if err := records.savePersonalInformation(personInfo); err != nil {
			log.Fatal("保存失败: ", err)
		}

		// 计算
		fatRate, err := c.FatRate(personInfo)
		if err != nil {
			log.Fatal("计算fr失败: ", err)
		}

		// 注册到rank
		rk.inputRecord(personInfo.Name, fatRate)

		rankResult, _ := rk.getRank(personInfo.Name) // todo handle error
		fmt.Println("排名结果: ", rankResult)
	}

}

func caseStudy() {
	filePath := "C:\\Users\\Administrator\\Desktop\\小强.self.information.json"

	personalInformation := apis.PersonalInformation{
		Name:   `"小"强""`,
		Sex:    "男",
		Tall:   1.70,
		Weight: 71,
		Age:    35,
	}

	fmt.Printf("%+v\n", personalInformation)
	data, err := json.Marshal(personalInformation)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("marshal的结果是(原生): ", data)
	fmt.Println("marshal的结果是(string): ", string(data))

	writeFile(filePath, data)

	readFile(filePath)
}

func readFile(filePath string) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("读取文件失败: ", err)
		return
	}

	fmt.Println("读取出来的内容是: ", string(data))

	personalInformation := apis.PersonalInformation{}
	json.Unmarshal(data, &personalInformation) // todo handle error

	fmt.Println("开始计算体脂信息: ", personalInformation)
	bmi, _ := gobmi.BMI(personalInformation.Weight, personalInformation.Tall) // todo handle error
	fmt.Printf("%s的bmi是: %v \n", personalInformation.Name, bmi)
	fatRate := gobmi.CalcFatRate(bmi, personalInformation.Age, personalInformation.Sex)
	fmt.Printf("%s的fatRate是: %v \n", personalInformation.Name, fatRate)

}

func writeFile(filePath string, data []byte) {
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("无法打开文件", filePath, "错误信息是：", err)
		os.Exit(1) // 异常退出，并输出错误代码
	}
	defer file.Close()

	_, err = file.Write(data)
	fmt.Println(err)
}

package apis

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"testing"
)

func TestMarshalJson(t *testing.T) {
	personalInformation := PersonalInformation{
		Name:   `"小"强""`, // 只有public的字段才可以编码到json格式中
		Sex:    "男",
		Tall:   1.70,
		Weight: 71,
		Age:    35,
	}

	fmt.Printf("%+v \n", personalInformation)
	data, err := json.Marshal(personalInformation)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("marshal的结果是(原生): ", data)
	fmt.Println("marshal的结果是(string): ", string(data))

}

func TestUnmarshalJson(t *testing.T) {
	data := `{"name":"\"小\"强\"\"","sex":"男","tall":1.7,"weight":71,"age":35}`
	personalInformation := PersonalInformation{}

	err := json.Unmarshal([]byte(data), &personalInformation)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v \n ", personalInformation)
}

func TestMarshalYaml(t *testing.T) {
	personalInformation := PersonalInformation{
		Name:   `"小"强""`,
		Sex:    "男",
		Tall:   1.70,
		Weight: 71,
		Age:    35,
	}

	fmt.Printf("%+v \n", personalInformation)
	data, err := yaml.Marshal(personalInformation)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("marshal的结果是(原生): ", data)
	fmt.Println("marshal的结果是(string): ", string(data))

}

func TestUnmarshalYaml(t *testing.T) {
	data := `name: '"小"强""'
sex: 男
tall: 1.7
weight: 71
age: 35`
	personalInformation := PersonalInformation{}

	err := yaml.Unmarshal([]byte(data), &personalInformation)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v \n", personalInformation)
}

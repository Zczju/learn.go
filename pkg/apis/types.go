package apis

type PersonalInformation struct {
	// name   string  `json:"name"` // 注意, 私有成员变量在序列化、反序列化时会被忽略
	Name   string  `json:"name"` //`yaml:"a,omitempty"` // 当两个注解名冲突的时候，很危险；选其中一个 // 如果yaml的marshal不支持json的注解，需要换包或者换版本
	Sex    string  `json:"sex"`
	Tall   float64 `json:"tall"`
	Weight float64 `json:"weight"`
	Age    int     `json:"age"`
}

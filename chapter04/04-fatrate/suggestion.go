package main

func GetFatRateSuggestion() *fatRateSuggestions {
	// 自己的初始化，我们自己定义一个函数去获取它
	return &fatRateSuggestions{
		suggArr: [][][]int{
			{
				// 第一个元素,表示男 // 用数组的下标表示逻辑意义
				{
					// 18-39
					0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 3, 3, 3, 3, 3, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4,
				},
				{
					// 40-59
				},
				{
					// 60+
				},
			},
			{
				// 第二个元素，表示女
			},
		},
	}
}

type fatRateSuggestions struct {
	suggArr [][][]int
}

func (s *fatRateSuggestions) GetSuggestions(person *Person) string {
	sexIdx := s.getIndexOfSex(person.sex)
	ageIdx := s.getIndexOfAge(person.age)
	maxFRSupported := len(s.suggArr[sexIdx][ageIdx]) - 1
	frIdx := int(person.fatRate * 100)
	if frIdx > maxFRSupported {
		frIdx = maxFRSupported
	}
	suggIdx := s.suggArr[sexIdx][ageIdx][frIdx]
	return s.translateResult(suggIdx)
}

func (s *fatRateSuggestions) getIndexOfSex(sex string) int {
	if sex == "男" {
		return 0
	}
	return 1
}
func (s *fatRateSuggestions) getIndexOfAge(age int) int {
	switch {
	case age >= 18 && age <= 39:
		return 0
	case age >= 40 && age <= 59:
		return 1
	case age >= 60:
		return 2
	default:
		return -1
	}
}
func (s *fatRateSuggestions) translateResult(idx int) string {
	switch idx {
	case 0:
		return "偏瘦"
	case 1:
		return "标准"
	case 2:
		return "偏重"
	case 3:
		return "肥胖"
	case 4:
		return "严重肥胖"
	}
	return "未知"
}

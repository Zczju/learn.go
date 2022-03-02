package main

import "testing"

func Test_fatRateSuggestions_GetSuggestions(t *testing.T) {
	sugg := GetFatRateSuggestion()
	tests := []Person{
		{
			sex:     "男",
			age:     35,
			fatRate: 0.24,
		},
	}
	if got := sugg.GetSuggestions(&tests[0]); got != "肥胖" {
		t.Fail()
	}
}

func Test_fatRateSuggestions_GetSuggestions1(t *testing.T) {
	sugg := GetFatRateSuggestion()
	type args struct {
		person *Person
	}
	tests := []struct { // 临时对象，仅在方法体内有效
		name string
		args args
		want string
	}{
		{
			name: "35-0.24",
			args: args{
				person: &Person{
					sex:     "男",
					age:     35,
					fatRate: 0.24,
				},
			},
			want: "肥胖",
		},
		{name: "35", args: args{person: &Person{sex: "男", age: 35, fatRate: 0}}, want: "偏瘦"},
		{name: "35", args: args{person: &Person{sex: "男", age: 35, fatRate: 0.01}}, want: "偏瘦"},
		{name: "35", args: args{person: &Person{sex: "男", age: 35, fatRate: 0.02}}, want: "偏瘦"},
		{name: "35", args: args{person: &Person{sex: "男", age: 35, fatRate: 0.03}}, want: "偏瘦"},
		{name: "35", args: args{person: &Person{sex: "男", age: 35, fatRate: 0.04}}, want: "偏瘦"},
		{name: "35", args: args{person: &Person{sex: "男", age: 35, fatRate: 0.05}}, want: "偏瘦"},
		{name: "35", args: args{person: &Person{sex: "男", age: 35, fatRate: 0.06}}, want: "偏瘦"},
		{name: "35", args: args{person: &Person{sex: "男", age: 35, fatRate: 0.07}}, want: "偏瘦"},
		{name: "35", args: args{person: &Person{sex: "男", age: 35, fatRate: 0.08}}, want: "偏瘦"},
		{name: "35", args: args{person: &Person{sex: "男", age: 35, fatRate: 0.09}}, want: "偏瘦"},
		{name: "35", args: args{person: &Person{sex: "男", age: 35, fatRate: 0.10}}, want: "标准"},
		{name: "35", args: args{person: &Person{sex: "男", age: 35, fatRate: 0.11}}, want: "标准"},
		{name: "35", args: args{person: &Person{sex: "男", age: 35, fatRate: 0.12}}, want: "标准"},
		{name: "35", args: args{person: &Person{sex: "男", age: 35, fatRate: 0.13}}, want: "标准"},
		{name: "35", args: args{person: &Person{sex: "男", age: 35, fatRate: 0.14}}, want: "标准"},
		{name: "35", args: args{person: &Person{sex: "男", age: 35, fatRate: 0.15}}, want: "标准"},
		{name: "35", args: args{person: &Person{sex: "男", age: 35, fatRate: 0.16}}, want: "偏重"},
		{name: "35", args: args{person: &Person{sex: "男", age: 35, fatRate: 0.17}}, want: "偏重"},
		{name: "35", args: args{person: &Person{sex: "男", age: 35, fatRate: 0.18}}, want: "偏重"},
		{name: "35", args: args{person: &Person{sex: "男", age: 35, fatRate: 0.19}}, want: "偏重"},
		{name: "35", args: args{person: &Person{sex: "男", age: 35, fatRate: 0.20}}, want: "偏重"},
		{name: "35", args: args{person: &Person{sex: "男", age: 35, fatRate: 0.21}}, want: "肥胖"},
		{name: "35", args: args{person: &Person{sex: "男", age: 35, fatRate: 0.22}}, want: "肥胖"},
		{name: "35", args: args{person: &Person{sex: "男", age: 35, fatRate: 0.23}}, want: "肥胖"},
		{name: "35", args: args{person: &Person{sex: "男", age: 35, fatRate: 0.24}}, want: "肥胖"},
		{name: "35", args: args{person: &Person{sex: "男", age: 35, fatRate: 0.25}}, want: "肥胖"},
		{name: "35", args: args{person: &Person{sex: "男", age: 35, fatRate: 0.26}}, want: "严重肥胖"},
		{name: "35", args: args{person: &Person{sex: "男", age: 35, fatRate: 0.27}}, want: "严重肥胖"},
		{name: "35", args: args{person: &Person{sex: "男", age: 35, fatRate: 0.28}}, want: "严重肥胖"},
		{name: "35", args: args{person: &Person{sex: "男", age: 35, fatRate: 0.29}}, want: "严重肥胖"},
		{name: "35", args: args{person: &Person{sex: "男", age: 35, fatRate: 0.30}}, want: "严重肥胖"},
		{name: "35", args: args{person: &Person{sex: "男", age: 35, fatRate: 0.31}}, want: "严重肥胖"},
		{name: "35", args: args{person: &Person{sex: "男", age: 35, fatRate: 0.32}}, want: "严重肥胖"},
		{name: "35", args: args{person: &Person{sex: "男", age: 35, fatRate: 0.33}}, want: "严重肥胖"},
		{name: "35", args: args{person: &Person{sex: "男", age: 35, fatRate: 0.34}}, want: "严重肥胖"},
		{name: "35", args: args{person: &Person{sex: "男", age: 35, fatRate: 0.35}}, want: "严重肥胖"},
		{name: "35", args: args{person: &Person{sex: "男", age: 35, fatRate: 0.36}}, want: "严重肥胖"},
		{name: "35", args: args{person: &Person{sex: "男", age: 35, fatRate: 0.37}}, want: "严重肥胖"},
		{name: "35", args: args{person: &Person{sex: "男", age: 35, fatRate: 0.38}}, want: "严重肥胖"},
		{name: "35", args: args{person: &Person{sex: "男", age: 35, fatRate: 0.39}}, want: "严重肥胖"},
		{name: "35", args: args{person: &Person{sex: "男", age: 35, fatRate: 0.40}}, want: "严重肥胖"},
		{name: "35", args: args{person: &Person{sex: "男", age: 35, fatRate: 0.41}}, want: "严重肥胖"},
		{name: "35", args: args{person: &Person{sex: "男", age: 35, fatRate: 0.42}}, want: "严重肥胖"},
		{name: "35", args: args{person: &Person{sex: "男", age: 35, fatRate: 0.43}}, want: "严重肥胖"},
		{name: "35", args: args{person: &Person{sex: "男", age: 35, fatRate: 0.44}}, want: "严重肥胖"},
		{name: "35", args: args{person: &Person{sex: "男", age: 35, fatRate: 0.94}}, want: "严重肥胖"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := sugg
			if got := s.GetSuggestions(tt.args.person); got != tt.want {
				t.Errorf("GetSuggestions() = %v, want %v", got, tt.want)
			}
		})
	}
}

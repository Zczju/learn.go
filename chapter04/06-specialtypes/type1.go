package main

import (
	"fmt"
	"math/rand"
	"sort"
)

type Math = int // 类型重命名
type English = int
type Chinese = int

func main() {
	var mathScore int = 100
	var mathScore2 Math = 100

	mathScore2 = mathScore
	fmt.Println(mathScore2)
	getScoreOfStudent("")

	vg := &voteGame{students: []*student{
		&student{name: fmt.Sprintf("%d", 1)},
		&student{name: fmt.Sprintf("%d", 2)},
		&student{name: fmt.Sprintf("%d", 3)},
		&student{name: fmt.Sprintf("%d", 4)},
		&student{name: fmt.Sprintf("%d", 5)},
	}}
	leader := vg.goRun()
	fmt.Println(leader)
	leader.Distribute()

	var stdXQ = &student{name: "小强"}
	var LdXQ Leader = Leader(*stdXQ)
	LdXQ.Distribute()

	// bytesTest1 := []byte{}
	// var str1 string = string(bytesTest1)
}

func getScoreOfStudent(name string) (Math, Chinese, English) {
	return 0, 0, 0
}

type voteGame struct {
	students []*student
}

func (g *voteGame) goRun() *Leader {
	// todo...
	for _, item := range g.students {
		randInt := rand.Int()
		if rand.Int()%2 == 0 {
			item.voteA(g.students[randInt%len(g.students)])
		} else {
			item.voteD(g.students[randInt%len(g.students)])
		}

	}
	maxScore := -1
	maxScoreIndex := -1
	for i, item := range g.students {
		if maxScore < item.agree {
			maxScore = item.agree
			maxScoreIndex = i
		}
	}
	if maxScoreIndex >= 0 { // 判断是否大于0，因为如果没有学生，那么index就是默认值-1
		return (*Leader)(g.students[maxScoreIndex])
	}
	return nil
}

// 使用嵌套对象定义（继承）的方式来定义班长
// type Leader struct {
// 	student
// }

// Leader 使用对象重定义
type Leader student

func (*Leader) Distribute() {
	fmt.Println("发作业啦~~")
}

type FoooTestFuncRedefine []string // 类型可以是任意的

func (f FoooTestFuncRedefine) Sort() {
	sort.Strings(f)
}

type student struct {
	name     string
	agree    int
	disagree int
}

// type Leader = student 类型重命名。中间有等号！！

func (std *student) voteA(target *student) {
	target.agree++
}
func (std *student) voteD(target *student) {
	target.disagree++
}

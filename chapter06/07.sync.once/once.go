package main

import (
	"fmt"
	"sync"
)

var counter int = 0
var counterOnce sync.Once

type rank struct {
	standard []string
}

var globalRank = &rank{}

// var globalRankInitialized bool = false
// var globalRankInitializedLock sync.Mutex

var once sync.Once = sync.Once{}

//func init() {
//	globalRank.standard = []string{"Asia"}
//	// init 函数无法做交互，无法传入参数，内容比较固定，每次需要重写
//}

//func initGlobalRankStandard(standard []string) { // 常用的 锁类型的只做一次
//	globalRankInitializedLock.Lock()
//	defer globalRankInitializedLock.Unlock()
//	if globalRankInitialized {
//		return
//	}
//	globalRank.standard = standard
//	globalRankInitialized = true
//}

func initGlobalRankStandard(standard []string) { // 用once比较方便，可省略很多行代码
	once.Do(func() {
		globalRank.standard = standard
	})
}

var facStore = &dbFactoryStore{}

type dbFactoryStore struct {
	store map[string]DBFactory
}

type Conn struct{}

type DBFactory interface {
	GetConnection() *Conn
}

func initMysqlFac(connStr string) DBFactory {
	return &MySqlDBFactory{}
}

type MySqlDBFactory struct {
	once sync.Once
}

func (m MySqlDBFactory) GetConnection() *Conn {
	m.once.Do(func() {
		initMysqlFac("")
	})
	// todo
	return nil
}

func main() {
	//standard := []string{"Asia"}
	//for i := 0; i < 10; i++ {
	//	go func() {
	//		initGlobalRankStandard(standard)
	//	}()
	//}

	for i := 0; i < 10; i++ {
		fmt.Printf("第%d次\n", i)
		counterOnce.Do(func() {
			fmt.Println("初始化")
			counter++
		})
	}
	fmt.Println("最终结果:", counter)

	// connStr := "xxxxxxxx"

}

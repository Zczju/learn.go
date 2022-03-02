package tmp

import "testing"

type Rank interface {
	UpdateFR(name string, fatRate float64)
	GetRank(name string) int
}

type Client interface {
	UpdateFR(name string, fatRate float64)
	GetRank(name string) int
}

func TestHomework(t *testing.T) {
	// var rank Rank // 实例化自己完成

	var clients []Client // 实例化自己完成

	for i := 0; i < len(clients); i++ {
		go func(idx int) {
			// todo add context to control exit
			go func() {
				for {
					clients[idx].UpdateFR("xiaoqiang", 0.23) // 0.23 to be replaced with base +- delta
				}
			}()

			go func() {
				for {
					clients[idx].GetRank("xiaoqiang")
				}
			}()
		}(i)
	}

}

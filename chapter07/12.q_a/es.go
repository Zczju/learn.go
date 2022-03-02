package main

// batch
//一次性提交一批请求, 但又不能太多

var sharedCli = &esClientWithBuffer{}

type esClientWithBuffer struct {
	batchBuffer [][]interface{} // 很多组的数据
	messageCh   chan interface{}
	shortBuffer []interface{}
}

func (cli *esClientWithBuffer) pushBatch() {
	// todo 队列操作
}

func (cli *esClientWithBuffer) prepareBatch() {
	// 由一个单独的routine来跑
	for msg := range cli.messageCh {
		if len(cli.shortBuffer) == batchSize {
			cli.batchBuffer = append(cli.batchBuffer, cli.shortBuffer)
			cli.shortBuffer = []interface{}{} // 将shortBuffer清空，再去装数据
		}
		cli.shortBuffer = append(cli.shortBuffer, msg)
	}
}

func pushToElasticSearchService(data interface{}) {
	// .... 延时push
	sharedCli.messageCh <- data

}

var batchSize int = 20

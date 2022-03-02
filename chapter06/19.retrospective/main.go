package main

import (
	"fmt"
	"time"
)

type Downloader struct {
	fileNameCh chan string
	// some kind of download worker
	finishedCh chan string
}

func (d *Downloader) Download() {
	for fileName := range d.fileNameCh {
		fmt.Println("开始下载文件：", fileName)
		time.Sleep(1 * time.Second)
		fmt.Println("开始处理文件：", fileName)
		time.Sleep(10 * time.Millisecond)
		fmt.Println("保存文件：", fileName)
		d.finishedCh <- fileName
		// 在此处 finishedCh 在协程中往里进的时候，主程序还没走到for range取的部分，所以死锁
	}

}

func main() {
	fileNameCh := make(chan string, 100)
	finishedCh := make(chan string, 100)
	// 如果这个finished chan 没有buffer，则在download完成后无法将结果放进该channel，从而导致死锁。
	// 没有buffer的channel必须有消费者才可以向channel内装数据。

	workerCount := 50
	for i := 0; i < workerCount; i++ {
		go func() {
			downloader := &Downloader{
				fileNameCh: fileNameCh,
				finishedCh: finishedCh,
			}
			downloader.Download()
		}()
	} // 启动worker启动完毕，但还未开始工作，因为还没向chan发信号

	//fileNames, fileNumber := produceFile() // 生成文件
	fileNumber := 100
	fileNames := make([]string, 0, fileNumber)
	for i := 0; i < fileNumber; i++ {
		fileNames = append(fileNames, fmt.Sprintf("file_%d.txt", i))
	}

	finishedFileCount := 0
	go func() {
		for finishedFile := range finishedCh {
			fmt.Println("文件：", finishedFile, "处理完毕")
			finishedFileCount++
			if finishedFileCount == fileNumber {
				close(finishedCh)
			}
		}
	}()

	for _, fileItem := range fileNames {
		fileNameCh <- fileItem
	}
	close(fileNameCh)

	fmt.Println("所有文件处理完毕，结束")
}

//func produceFile() ([]string, int) {
//	fileNumber := 100
//	fileNames := make([]string, 0, fileNumber)
//	for i := 0; i < fileNumber; i++ {
//		fileNames = append(fileNames, fmt.Sprintf("file_%d.txt", i))
//	}
//	return fileNames, fileNumber
//}

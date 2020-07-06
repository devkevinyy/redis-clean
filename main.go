package main

import (
	"encoding/csv"
	"fmt"
	"github.com/chujieyang/redis-clean/utils"
	"io"
	"io/ioutil"
	"os"
	"time"
)

func main() {
	startTime := time.Now()
	db8Count := 0
	db50Count := 0
	//csvPath := "/Users/yangchujie/GoProjects/src/github.com/chujieyang/redis-clean/csv/data.csv"

	csvDirPath := "/home/fuluops/csv/"
	rd, err := ioutil.ReadDir(csvDirPath)
	if err != nil {
		fmt.Println("read dir fail:", err)
		return
	}
	for _, fi := range rd {
		csvPath := fmt.Sprintf("%s%s", csvDirPath, fi.Name())
		fmt.Println(csvPath)
		file, err := os.Open(csvPath)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		defer func() {
			if err := file.Close(); err != nil {
				fmt.Println(err)
			}
		}()
		reader := csv.NewReader(file)
		batchCount := 0
		fileEnd := false
		var db8Keys []interface{}
		var db50Keys []interface{}
		for {
			record, err := reader.Read()
			if err == io.EOF {
				fmt.Println("csv read finish")
				fileEnd = true
			} else if err != nil {
				fmt.Println("Error:", err)
				return
			}
			if fileEnd == false {
				batchCount += 1
				db8Keys = append(db8Keys, fmt.Sprintf("CustomerOrderNo_Counter_%s_%s",
					record[3], record[2]))
				db50Keys = append(db50Keys, fmt.Sprintf("GenerateOrderId_%s", record[1]),
					fmt.Sprintf("GenerateSubOrderId_%s", record[0]))
			}
			if batchCount == 300 || fileEnd == true {  // 批量删除
				fmt.Println(db8Keys)
				fmt.Println(db50Keys)
				if err := utils.RemoveRedisKeys(8, db8Keys, &db8Count); err != nil {
					fmt.Println(err)
				}
				if err := utils.RemoveRedisKeys(50, db50Keys, &db50Count); err != nil {
					fmt.Println(err)
				}
				time.Sleep(300*time.Microsecond)
				batchCount = 0
				db8Keys, db50Keys = []interface{}{}, []interface{}{}
				if fileEnd == true {
					break
				}
			}
		}

	}

	fmt.Println(fmt.Sprintf("执行完成，共删除db8数量: %d, db50数量: %d, 耗时：%s", db8Count, db50Count, time.Since(startTime).String()))
}

package main

import (
	"encoding/csv"
	"fmt"
	"github.com/chujieyang/redis-clean/utils"
	"io"
	"os"
	"time"
)

func main() {
	startTime := time.Now()
	db8Count := 0
	db50Count := 0
	//csvPath := "/Users/yangchujie/GoProjects/src/github.com/chujieyang/redis-clean/csv/data.csv"
	csvPath := "/home/fuluops/csv/data.csv"
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
	var db8Keys []interface{}
	var db50Keys []interface{}
	for {
		record, err := reader.Read()
		if err == io.EOF {
			fmt.Println("csv read finish")
			break
		} else if err != nil {
			fmt.Println("Error:", err)
			return
		}
		batchCount += 1
		db8Keys = append(db8Keys, fmt.Sprintf("CustomerOrderNo_Counter_%s_%s",
			record[0], record[1]))
		db50Keys = append(db50Keys, fmt.Sprintf("GenerateOrderId_%s", record[2]),
			fmt.Sprintf("GenerateSubOrderId_%s", record[3]))
		if batchCount == 500 {  // 批量删除
			fmt.Println(db8Keys)
			fmt.Println(db50Keys)
			if err := utils.RemoveRedisKeys(8, db8Keys, &db8Count); err != nil {
				fmt.Println(err)
			}
			if err := utils.RemoveRedisKeys(50, db50Keys, &db50Count); err != nil {
				fmt.Println(err)
			}
			time.Sleep(100*time.Microsecond)
			batchCount = 0
			db8Keys, db50Keys = []interface{}{}, []interface{}{}
		}
	}
	//db8Count := 0
	//db50Count := 0
	//startTime := time.Now()
	//pageSize := 200
	//for page := 1; page < 10; page++ {
	//	sql := fmt.Sprintf("SELECT * FROM (SELECT Id, OrderId, CustomerOrderNo, MemberId, " +
	//		"ROW_NUMBER() OVER(ORDER BY Id) AS rowindex FROM dbo.SubOrder_202004) t WHERE t.rowindex " +
	//		"BETWEEN %d AND %d", (page-1)*pageSize, page*pageSize)
	//	dataList := utils.QueryData(sql)
	//	var db8Keys []interface{}
	//	var db50Keys []interface{}
	//	for _, item := range dataList {
	//		db8Keys = append(db8Keys, fmt.Sprintf("CustomerOrderNo_Counter_%s_%s",
	//			item.MemberId, item.CustomerOrderNo))
	//		db50Keys = append(db50Keys,
	//			fmt.Sprintf("GenerateOrderId_%s", item.OrderId),
	//			fmt.Sprintf("GenerateSubOrderId_%s", item.Id))
	//	}
	//	fmt.Println(db8Keys)
	//	fmt.Println(db50Keys)
	//	if err := utils.RemoveRedisKeys(8, db8Keys, &db8Count); err != nil {
	//		fmt.Println(err)
	//	}
	//	if err := utils.RemoveRedisKeys(50, db50Keys, &db50Count); err != nil {
	//		fmt.Println(err)
	//	}
	//}
	fmt.Println(fmt.Sprintf("执行完成，共删除db8数量: %d, db50数量: %d, 耗时：%s", db8Count, db50Count, time.Since(startTime).String()))
}

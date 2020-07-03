package utils

import (
	"fmt"
	"github.com/chujieyang/redis-clean/conf"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
)

var dbClient *gorm.DB

type SubOrder struct {
	Id string `gorm:"column:Id"`
	OrderId string `gorm:"column:OrderId"`
	MemberId string `gorm:"column:MemberId"`
	CustomerOrderNo string `gorm:"column:CustomerOrderNo"`
}

func (u SubOrder)TableName() string  {
	return "SubOrder_202004"
}

func init() {
	var err error
	fmt.Println("SqlServer连接池初始化...")
	connString := fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s",
		conf.SqlServerUser, conf.SqlServerPwd, conf.SqlServerHost, conf.SqlServerPort, conf.SqlServerDB)
	dbClient, err = gorm.Open("mssql", connString)
	if err != nil {
		panic(err)
	}
	fmt.Println("SqlServer连接初始化: ", dbClient.DB().Stats().OpenConnections)
}

func QueryData(sql string) (orders []SubOrder) {
	err := dbClient.Raw(sql).Scan(&orders).Error
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

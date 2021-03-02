package main

import (
	"boot/common"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/withlin/canal-go/client"
	protocol "github.com/withlin/canal-go/protocol"
)

func main() {

	// 192.168.199.17 替换成你的canal server的地址
	// example 替换成-e canal.destinations=example 你自己定义的名字
	connector := client.NewSimpleCanalConnector("127.0.0.1", 8086, "", "", "example", 60000, 60*60*1000)
	err := connector.Connect()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	// https://github.com/alibaba/canal/wiki/AdminGuide
	//mysql 数据解析关注的表，Perl正则表达式.
	//
	//多个正则之间以逗号(,)分隔，转义符需要双斜杠(\\)
	//
	//常见例子：
	//
	//  1.  所有表：.*   or  .*\\..*
	//	2.  canal schema下所有表： canal\\..*
	//	3.  canal下的以canal打头的表：canal\\.canal.*
	//	4.  canal schema下的一张表：canal\\.test1
	//  5.  多个规则组合使用：canal\\..*,mysql.test1,mysql.test2 (逗号分隔)

	err = connector.Subscribe(".*\\..*")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	//  连接nats
	err = common.ConnectNats()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	for {

		message, err := connector.Get(100, nil, nil)
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
		batchId := message.Id
		if batchId == -1 || len(message.Entries) <= 0 {
			time.Sleep(300 * time.Millisecond)
			//fmt.Println("===没有数据了===")
			continue
		}
		if batchId != -1 || len(message.Entries) >= 0 {
			time.Sleep(300 * time.Millisecond)
			printEntry(message.Entries)

		}
		fmt.Println(message)

	}
}

func printEntry(entrys []protocol.Entry) {

	for _, entry := range entrys {
		if entry.GetEntryType() == protocol.EntryType_TRANSACTIONBEGIN || entry.GetEntryType() == protocol.EntryType_TRANSACTIONEND {
			continue
		}
		rowChange := new(protocol.RowChange)

		err := proto.Unmarshal(entry.GetStoreValue(), rowChange)
		checkError(err)
		if nil != rowChange {
			eventType := rowChange.GetEventType()
			header := entry.GetHeader()
			fmt.Println(fmt.Sprintf("================> binlog[%s : %d],name[%s,%s], eventType: %s", header.GetLogfileName(), header.GetLogfileOffset(), header.GetSchemaName(), header.GetTableName(), header.GetEventType()))

			if rowChange.GetSql() != "" {
				fmt.Println("-----sql不为空--> ")
				fmt.Println("-------> sql")
				fmt.Printf("---sql语句---->:%s ", rowChange.GetSql())
				fmt.Println()
				// 构建传输对象
				b, _ := json.Marshal(rowChange.GetSql())
				// 通过nat发布窗口评价信息
				if err := common.NatsCli.Publish("one_no_notice", b); err != nil {
					fmt.Println("-------> 发送日志失败")
				}

			}

			for _, rowData := range rowChange.GetRowDatas() {
				if eventType == protocol.EventType_DELETE {
					printColumn(rowData.GetBeforeColumns())
				} else if eventType == protocol.EventType_INSERT {
					fmt.Printf("进入Insert：%s \n", "hello")
					columns := rowData.GetAfterColumns()
					printColumn(columns)

					// 发数据库的某行变化
					columnMap := make(map[string]string)
					for _, col := range columns {
						fmt.Printf("col的mysqlType值：%s  %s  sql type: %v %v\n", col.MysqlType, col.GetMysqlType(), col.SqlType, col.GetSqlType())
						fmt.Printf("col的index值：%d  是否key %v  string(): %v  %v\n", col.Index, col.IsKey, col.String(), col.XXX_unrecognized)
						columnMap[col.GetName()] = col.GetValue()
					}

					mapB, _ := json.Marshal(columnMap)
					fmt.Printf("map的值：%s \n", mapB)

					// 通过nat发布列信息
					if err := common.NatsCli.Publish("columnChange", mapB); err != nil {
						fmt.Println("-------> 发送change日志失败")
					}
				} else {
					fmt.Println("-------> before")
					printColumn(rowData.GetBeforeColumns())
					fmt.Println("-------> after")
					printColumn(rowData.GetAfterColumns())

				}
			}
		}
	}
}

func printColumn(columns []*protocol.Column) {
	for _, col := range columns {
		fmt.Println(fmt.Sprintf("%s : %s  update= %t", col.GetName(), col.GetValue(), col.GetUpdated()))
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

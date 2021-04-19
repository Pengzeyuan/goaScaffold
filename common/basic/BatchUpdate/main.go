package batchUpdate

import (
	"bytes"
	"fmt"
	"math"
	"reflect"
	"strconv"
	"strings"

	"go.uber.org/zap"
)

// tableName表的名字，itemList你定义的数组类型的结构体，[]*Demo
func BuildBatchUpdateSQLArray(tableName string, dataList interface{}) []string {

	fieldValue := reflect.ValueOf(dataList)
	fmt.Printf("fieldValue:%#v \n", fieldValue)
	// 队列的一个元素 指针的实体类型
	fieldType := reflect.TypeOf(dataList).Elem().Elem()
	fmt.Printf("fieldType:%#v \n", fieldType)
	sliceLength := fieldValue.Len()
	// fieldCount
	fieldNum := fieldType.NumField()
	fmt.Printf("fieldNum:%#v \n", fieldNum)

	// 检验结构体标签是否为空和重复
	verifyTagDuplicate := make(map[string]string)
	// 三个字段
	for i := 0; i < fieldNum; i++ {
		fieldTag := fieldType.Field(i).Tag.Get("gorm")

		fieldName := GetFieldName(fieldTag)
		if len(strings.TrimSpace(fieldName)) == 0 {
			zap.L().Panic("the structure attribute should have tag")
			return nil
		}

		//if !strings.HasPrefix(fieldName, "id;") {
		//	zap.L().Panic("the structure attribute should have primary_key")
		//	return nil
		//}

		_, ok := verifyTagDuplicate[fieldName]
		if !ok {
			// fieldName
			verifyTagDuplicate[fieldName] = fieldName
		} else {
			zap.L().Panic("the structure attribute %v tag is not allow duplication")
			return nil
		}

	}

	var IDList []string
	updateMap := make(map[string][]string)
	for i := 0; i < sliceLength; i++ {
		// 得到某一个具体的结构体的实体
		structValue := fieldValue.Index(i).Elem()
		for j := 0; j < fieldNum; j++ {
			elem := structValue.Field(j)

			var temp string
			// 第一个域
			switch elem.Kind() {
			case reflect.Int64:
				// 如果元素是int 格式化它的int值
				temp = strconv.FormatInt(elem.Int(), 10)
			case reflect.String:
				if strings.Contains(elem.String(), "'") {
					temp = fmt.Sprintf("'%v'", strings.ReplaceAll(elem.String(), "'", "\\'"))
				} else {
					temp = fmt.Sprintf("'%v'", elem.String())
				}
			case reflect.Float64:
				temp = strconv.FormatFloat(elem.Float(), 'f', -1, 64)
			case reflect.Bool:
				temp = strconv.FormatBool(elem.Bool())
			default:
				// 返回value 和json key关联的value
				zap.L().Panic(fmt.Sprintf("type conversion error, param is %v", fieldType.Field(j).Tag.Get("json")))
				return nil
			}
			// 类型的第一个域 得到value（string）  被gorm关联的
			gormTag := fieldType.Field(j).Tag.Get("gorm")
			// eg: column
			fieldTag := GetFieldName(gormTag)
			// 以id开头
			if strings.HasPrefix(fieldTag, "id;") {
				// string -> int
				id, err := strconv.ParseInt(temp, 10, 64)
				if err != nil {
					zap.L().Panic("转换id", zap.Error(err))
					return nil
				}
				// id 的合法性校验
				if id < 1 {
					zap.L().Panic("this structure should have a primary key and gt 0")
					return nil
				}
				// id 列表
				IDList = append(IDList, temp)
				continue
			}
			// 值列表    updateMap  string数组Map
			//valuelist 一个string数组
			valueList := append(updateMap[fieldTag], temp)
			// 更新map （ eg： column）代表一个数组
			updateMap[fieldTag] = valueList
		}
	}
	// 多少id 就是多少项数据
	length := len(IDList)
	size := 200
	// eg : 300 个id    /200  = 2  SQLQuantity=2
	SQLQuantity := getSQLQuantity(length, size)
	var SQLArray []string
	k := 0
	// eg:2   批次   就是执行两回这个  构建的sql
	for i := 0; i < SQLQuantity; i++ {
		count := 0
		// 拼接字符
		var record bytes.Buffer
		// 构建sql
		record.WriteString("UPDATE " + tableName + " SET ")
		//  key 类如 fieldName可能colum，set 什么
		for fieldName, fieldValueList := range updateMap {
			// 构建每一个字段
			// 数据库中的字段
			record.WriteString(fieldName)
			record.WriteString(" = CASE " + "id")
			// 分支条件  k=0
			for j := k; j < len(IDList) && j < len(fieldValueList) && j < size+k; j++ {
				// 设置id  eg：  1-》 1    2 -》 2     3-》  33
				record.WriteString(" WHEN " + IDList[j] + " THEN " + fieldValueList[j])
			}
			count++
			if count != fieldNum-1 {
				record.WriteString(" END, ")
			}
		}

		record.WriteString(" END WHERE ")
		record.WriteString("id" + " IN (")
		min := size + k
		if len(IDList) < min {
			min = len(IDList)
		}
		// 在那些id里  按 200 来分组
		record.WriteString(strings.Join(IDList[k:min], ","))
		record.WriteString(");")
		//  游标后移
		k += size
		// 加入这一组sql
		SQLArray = append(SQLArray, record.String())
	}

	return SQLArray
}

func getSQLQuantity(length, size int) int {
	SQLQuantity := int(math.Ceil(float64(length) / float64(size)))
	return SQLQuantity
}

func GetFieldName(fieldTag string) string {
	// 分割字符
	fieldTagArr := strings.Split(fieldTag, ":")
	if len(fieldTagArr) == 0 {
		return ""
	}
	// 拿前面一位
	fieldName := fieldTagArr[len(fieldTagArr)-1]

	return fieldName
}

func main() {

}

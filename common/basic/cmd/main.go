package main

import (
	batchUpdate "basic/BatchUpdate"
	initDB "basic/gormBatchInsert"
	"basic/model"
	"bytes"
	"context"
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"

	"github.com/jinzhu/gorm"

	"go.uber.org/zap"
)

var logg *log.Logger

func main() {

	funcNameBuffer()

	//BatchUpdate()
	//Test8Gorm2()
	//logg = log.New(os.Stdout, "", log.Ltime)
	//timeoutHandler()
	//logg.Printf("end")

	//
	//c := make(chan int)
	//
	//go run(c)
	//
	//fmt.Println("wait")
	//time.Sleep(time.Second * 5)
	//
	//c <- 1
	//<-c
	//
	//fmt.Println("main exited")

	//test8Context()
	//f5()
	//f4()
	//f3()
}

func funcNameBuffer() {
	bufs := bytes.NewBufferString("学swift.")
	fmt.Println(bufs.String())

	//读取第一个rune,赋值给r
	r, z, _ := bufs.ReadRune()
	//打印中文"学",缓冲器头部第一个被拿走
	fmt.Println(bufs.String())
	//打印"学","学"作为utf8储存占3个byte
	fmt.Println("r=", string(r), ",z=", z)
}

func BatchUpdate() {
	initDB.InitDB()
	DB := initDB.DpDB

	var demo []*model.Demo
	demo1 := &model.Demo{1, "nihao", 12.0}
	demo2 := &model.Demo{2, "renzhen", 13.0}
	demo3 := &model.Demo{3, "duidai", 12.0}
	demo4 := &model.Demo{4, "xiexie", 13.0}
	demo5 := &model.Demo{5, "OOP", 12.0}
	demo = append(demo, demo1, demo2, demo3, demo4, demo5)
	sqlStr := batchUpdate.BuildBatchUpdateSQLArray("demo", demo)
	fmt.Printf("%v", sqlStr)

	DB.Exec(sqlStr[0])
}

func test8Context() {
	rand.Seed(time.Now().Unix())

	ctx, _ := context.WithTimeout(context.Background(), time.Second*3)

	var wg sync.WaitGroup
	wg.Add(1)
	go GenUsers(ctx, &wg)
	wg.Wait()

	fmt.Println("生成幸运用户成功")
}

func run(done chan int) {
	for {
		select {
		case <-done:
			fmt.Println("exiting...")
			done <- 1
			break
		default:
		}

		time.Sleep(time.Second * 1)
		fmt.Println("do something")
	}
}
func GenUsers(ctx context.Context, wg *sync.WaitGroup) { //生成用户ID
	fmt.Println("开始生成幸运用户")
	users := make([]int, 0)
guser:
	for {
		select {
		case <-ctx.Done(): //代表父context发起 取消操作

			fmt.Println(users)
			wg.Done()
			break guser
			return
		default:
			users = append(users, getUserID(1000, 100000))
		}
	}

}
func getUserID(min int, max int) int {
	return rand.Intn(max-min) + min
}

func someHandler() {
	//ctx, cancel := context.WithCancel(context.Background())
	//go doStuff(ctx)
	//
	////10秒后取消doStuff
	//time.Sleep(10 * time.Second)
	//cancel()

	// 创建继承Background的子节点Context
	ctx, cancel := context.WithCancel(context.Background())
	go doSth(ctx)

	//模拟程序运行 - Sleep 5秒
	time.Sleep(5 * time.Second)
	cancel()
}
func doTimeOutStuff(ctx context.Context) {
	for {
		time.Sleep(1 * time.Second)

		if deadline, ok := ctx.Deadline(); ok { //设置了deadl
			logg.Printf("deadline set")
			if time.Now().After(deadline) {
				logg.Printf(ctx.Err().Error())
				return
			}

		}

		select {
		case <-ctx.Done():
			logg.Printf("done")
			return
		default:
			logg.Printf("work")
		}
	}
}
func timeoutHandler() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(5*time.Second))
	go doTimeOutStuff(ctx)
	//go doStuff(ctx)

	time.Sleep(10 * time.Second)

	cancel()

}

//每1秒work一下，同时会判断ctx是否被取消，如果是就退出
func doSth(ctx context.Context) {
	var i = 1
	for {
		time.Sleep(1 * time.Second)
		select {
		case <-ctx.Done():
			fmt.Println("done")
			return
		default:
			fmt.Printf("work %d seconds: \n", i)
		}
		i++
	}
}

//每1秒work一下，同时会判断ctx是否被取消了，如果是就退出
func doStuff(ctx context.Context) {
	for {
		time.Sleep(1 * time.Second)
		select {
		case <-ctx.Done():
			logg.Printf("done")
			return
		default:
			logg.Printf("work")
		}
	}
}

func Test1() {
	ctx, cancelFunc := context.WithCancel(context.Background())
	go watch(ctx, "【监控1】")
	go watch(ctx, "【监控2】")
	go watch(ctx, "【监控3】")
	time.Sleep(10 * time.Second)
	fmt.Println("可以了，通知监控停止")
	cancelFunc()
	//为了检测监控过是否停止，如果没有监控输出，就表示停止了
	time.Sleep(5 * time.Second)
}
func watch(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, "监控退出，停止了...")
			return
		default:
			fmt.Println(name, "goroutine监控中...")
			time.Sleep(2 * time.Second)
		}
	}
}

func Test8Gorm2() {
	initDB.InitGorm2()
	//

	DBGorm2 := initDB.DB2

	//
	user := model.User{ID: 16, UserName: "admin", DeletedAt: nil}
	// Omit() 中填入所有此次插入需要忽略的值，Users struct 中的其他值使用给定值或对应零值来更新

	//DBGorm2.Omit("IsFemale", "IsActived").Save(&user)

	//DB.Delete(&user)
	DBGorm2.Model(model.User{}).First(&user)

}

// 关联查询
func f5() {

	//Test6Preload(DB)

	//Test7Related(DB)

	//user := test7Update(DB)

	// Omit() 中填入所有此次插入需要忽略的值，Users struct 中的其他值使用给定值或对应零值来更新

	//user := model.User{ID: 16, UserName: "admin", DeletedAt: nil}
	// Omit() 中填入所有此次插入需要忽略的值，Users struct 中的其他值使用给定值或对应零值来更新

	//DB.Omit("IsFemale", "IsActived").Save(&user)
	//
	////DB.Delete(&user)
	//DB.Model(model.User{}).Where("id = ?", 16).Updates(&user)

}

func test7Update(DB *gorm.DB) model.User {
	article := model.Article{}
	//先查询一条记录, 保存在模型变量food
	//等价于: SELECT * FROM `foods`  WHERE (id = '2') LIMIT 1
	DB.Where("id = ?", 2).Take(&article)
	//修改food模型的值
	article.Title = "java多线程100"

	//等价于: UPDATE `foods` SET `title` = '可乐', `type` = '0', `price` = '100', `stock` = '26', `create_time` = '2018-11-06 11:12:04'  WHERE `foods`.`id` = '2'
	DB.Save(&article)

	//例子1:
	//更新food模型对应的表记录
	//等价于: UPDATE `foods` SET `price` = '25'  WHERE `foods`.`id` = '2'
	DB.Model(&article).Update("title", "java多线程10022")
	//通过food模型的主键id的值作为where条件，更新price字段值。

	DB.Model(model.Article{}).Where("title = ?", "c内存").Update("title", "c内存25")

	//例子1：
	//通过结构体变量设置更新字段
	employee := model.Employee{Id: 2}
	updataFood := model.Employee{
		Age:      121,
		UserName: "柠檬雪碧1",
		Addr:     "青霞路",
	}

	//根据food模型更新数据库记录
	//等价于: UPDATE `foods` SET `price` = '120', `title` = '柠檬雪碧'  WHERE `foods`.`id` = '2'
	//Updates会忽略掉updataFood结构体变量的零值字段, 所以生成的sql语句只有price和title字段。
	//DB.Model(&employee).Updates(&updataFood)

	//设置Where条件，Model参数绑定一个空的模型变量
	//等价于: UPDATE `foods` SET `stock` = '120', `title` = '柠檬雪碧'  WHERE (price > '10')
	DB.Model(model.Employee{}).Where("age > ?", 100).Updates(&updataFood)
	//例子3:
	//如果想更新所有字段值，包括零值，就是不想忽略掉空值字段怎么办？
	//使用map类型，替代上面的结构体变量

	//定义map类型，key为字符串，value为interface{}类型，方便保存任意值
	data := make(map[string]interface{})
	data["addr"] = "" //零值字段
	data["age"] = 35
	//等价于: UPDATE `foods` SET `price` = '35', `stock` = '0'  WHERE (id = '2')
	DB.Model(model.Employee{}).Where("id = ?", 2).Updates(data)

	//等价于: UPDATE `foods` SET `stock` = stock + 1  WHERE `foods`.`id` = '2'
	DB.Model(&employee).Update("age", gorm.Expr("age + 1"))

	user := model.User{ /*Id: 16*/ }
	// 查询用户数据
	//自动生成sql： SELECT * FROM `users`  WHERE (username = 'tizi365') LIMIT 1
	DB.Where("user_name = ?", "tom1").First(&user)
	//DB.Where("id = ?", 2).Take(&article)
	fmt.Println(user)

	var profile model.Profile
	// 通过user关联查询Profile数据, 查询结果保存到profile变量
	DB.Model(&user).Related(&profile) // 属于关系
	// 自动生成sql: SELECT * FROM profiles WHERE user_id = 1 // 1 就是user的 ID，已经自动关联
	return user
}

// 多对多
func Test7Related(DB *gorm.DB) {
	var article model.Article
	DB.Where("title like ?", "%c%").First(&article)
	fmt.Println(article)
	err := DB.Model(&article).
		Related(&article.Category).
		Related(&article.Tag, "tag"). // 这里的foreignKeys  还不能少
		Find(&article).Error
	if err != nil {
		zap.L().Panic("Related fail", zap.Error(err))
	}
	//  查询article队列
	articles := []model.Article{}
	err = DB.Model(articles).
		Where("title like ?", "%c%").
		Preload("Category").
		Preload("Tag").Find(&articles).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return
	}
}

func Test6Preload(DB *gorm.DB) {
	var u model.User
	DB.First(&u)
	//DB.Model(&u).Related(&u.Companies).Find(&u.Companies)
	//var users []model.User
	//// 方法2
	//DB.Where("user_name = @name OR name2 = @name", sql.Named("name", "张飞")).Find(&users)
	//users = []model.User{{ID: 1}, {ID: 36}}

	//DB.Model(&u).Association("Companies").Find(&u.Companies)
	// 三方法
	// 查询单条 user
	DB.Debug().Preload("Companies").First(&u)
	// 对应的 sql 语句
	// SELECT * FROM users LIMIT 1;
	// SELECT * FROM companies WHERE user_id IN (1);

	// 查询所有 user
	list := []model.User{{ID: 16}, {ID: 36}, {ID: 1}}
	DB.Debug().Preload("Companies").Find(&list)
	// 对应的 sql 语句
	// SELECT * FROM users;
	// SELECT * FROM companies WHERE user_id IN (1,2,3...);
}

//根据标签反向关联审核通过的文章 ，这里得到的this是得到的文章

//  手写批量插入
func f4() {
	initDB.InitDB()
	DB := initDB.DpDB

	var ids []int
	var results []model.Employee
	// Get the list of rows that will be affected
	DB.Where("age !=12").Table("employee").Select("id").Find(&results)

	for i := 0; i < len(results); i++ {
		ids = append(ids, results[i].Id)
	}

	query := DB.Where("id IN (?)", ids)

	// Do the update
	query.Model(&model.Employee{}).Updates(model.Employee{Addr: "长安街1号"})

	// Get the updated rows
	query.Find(&results)

	cats := []model.Employee{}

	cat1 := model.Employee{Id: 1, UserName: "blackCap1", Age: 9, Addr: "120号"}
	cat2 := model.Employee{Id: 2, UserName: "blackCap2", Age: 10, Addr: "121号"}
	//cat3:=model.Cat{Id:112,CatName: "blackCap3",CatPrice: "122"}
	cats = append(cats, cat1)
	cats = append(cats, cat2)

	err := initDB.BatchSave(DB, cats)
	if err != nil {
		zap.L().Panic("init config fail", zap.Error(err))
	}
	//cats := []model.Cat{}
	//
	//cat1 := model.Cat{Id: 110, CatName: "blackCap1", CatPrice: "120", CatType: 9}
	//cat2 := model.Cat{Id: 111, CatName: "blackCap2", CatPrice: "121", CatType: 10}
	////cat3:=model.Cat{Id:112,CatName: "blackCap3",CatPrice: "122"}
	//cats = append(cats, cat1)
	//cats = append(cats, cat2)
	////cats = append(cats, cat3)
	////intCats:=[]interface{}
	////intCats:= append(intCats, cat1)
	//
	//var interfaceSlice []interface{} = make([]interface{}, len(cats))
	//for i, d := range cats {
	//	interfaceSlice[i] = d
	//}
	//initDB.BatchCreateModelsByPage(DB, interfaceSlice, "cat")
}
func f1() {
	var name string
	var name1 = []byte("yangyanxing汉字")
	name = string(name1)
	fmt.Println(&name, name)
	name = "杨彦星"             //这个是可以的
	fmt.Println(&name, name) //name 的地址不会改变 //非法 //
	//name[0] = "a"
}
func f2() {
	var name = "yangyanxing"
	fmt.Println(&name, name)
	names := []byte(name) //字节数组需要使用单引号,双引号是字符串了
	names[0] = 'f'
	fmt.Println(&name, string(names))
	name = "杨彦星" //这个是可以的
	// 如果有汉字的话需要使用字符数组
	namer := []rune(name) //需要使用单引号
	namer[0] = '饭'        //name 的地址不会改变
	fmt.Println(&name, string(namer))
}
func f3() {
	var name = "yangyanxing"
	fmt.Println(&name, name)
	names := []byte(name)
	fmt.Printf("%T", names)
}

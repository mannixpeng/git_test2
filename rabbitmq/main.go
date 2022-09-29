package main

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/robfig/cron/v3"
	"sync"
	"time"
)

type T1 struct {
	A int
	B string
}

type T2 struct {
	A int64
	B string
}
type PerformanceDetailHandler struct {
	sync.Mutex
	exceptionMapping map[int64][]interface{}
	handleCount int64
}

func (l *PerformanceDetailHandler) pushExceptionData(id int64) {
	l.Lock()
	defer l.Unlock()

	v, ok := l.exceptionMapping[l.handleCount]
	if !ok {
		list := make([]interface{}, 0)
		v = list
		l.exceptionMapping[l.handleCount] = v
	}
	v = append(v, id)
}

type UserUpdateResp struct {
	ID            uint   `orm:"id" json:"id"`
	PlatformId    uint   `orm:"platform_id" json:"platform_id" form:"platform_id" gorm:"index:idx_platform_id;column:platform_id;comment:平台ID;"`
	ApplicationId int    `orm:"application_id" json:"application_id" form:"application_id" gorm:"index:idx_application_id;column:application_id;comment:所属应用;"`
	NickName      string `orm:"nick_name" json:"nick_name" form:"nick_name" gorm:"column:nick_name;comment:昵称;type:varchar(50);size:50;"`
	RealName      string `orm:"real_name" json:"real_name" form:"real_name" gorm:"column:real_name;comment:真实姓名;type:varchar(50);size:50;"`
	Mobile        string `orm:"mobile" json:"mobile" form:"mobile" gorm:"column:mobile;comment:手机号;type:varchar(32);size:32;"`
	Email         string `orm:"email" json:"email" form:"email" gorm:"column:email;comment:邮箱;type:varchar(100);size:100;"`
	Sex           int8   `orm:"sex" json:"sex" form:"sex" gorm:"column:sex;comment:性别;1=男,2=女;type:tinyint;"`
	Avatar        string `orm:"avatar" json:"avatar" form:"avatar" gorm:"column:avatar;default:'1001';comment:头像;type:varchar(255);size:255;"`
}

type WebsocketAccountInfo struct {
	UserUpdateResp

	FbUserId     string `orm:"fb_user_id" json:"fb_user_id" form:"fb_user_id" gorm:"index:idx_fb_user_id;column:fb_user_id;comment:FB的UserId;type:varchar(128);size:128"`
	TotalBalance int    `orm:"total_balance" json:"total_balance" form:"total_balance" gorm:"column:total_balance;comment:总余额（包含冻结及保险箱余额）;type:bigint(20);"`
	Rank         int    `orm:"rank" json:"rank" form:"rank" gorm:"column:rank;comment:所在平台设置等级;type:int(10);size:10"`
}


func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1)

	cron := cron.New()
	fmt.Println("startTime: ", time.Now())
	cron.AddFunc("@every 30m", func() {
		fmt.Println("endTime: ", time.Now())
		fmt.Println("hello world")
		wg.Done() })
	cron.Start()
	defer cron.Stop()

	// Cron should fire in 2 seconds. After 1 second, call Entries.
	select {
	case <-time.After(1*time.Second):
		cron.Entries()
	}
	wg.Wait()
	return

	c := make(chan struct{}, 100)
	fmt.Println("len: ", len(c), "cap: ", cap(c))
	return
	info := WebsocketAccountInfo{
		UserUpdateResp: UserUpdateResp{
			ID:            9889195,
			PlatformId:    0,
			ApplicationId: 0,
			NickName:      "",
			RealName:      "",
			Mobile:        "",
			Email:         "",
			Sex:           0,
			Avatar:        "",
		},
		FbUserId:     "",
		TotalBalance: 0,
		Rank:         1,
	}
	content, _ := json.Marshal(info)
	fmt.Printf("%s", content)

	//sub := time.Now().Add(-3 * time.Hour)
	//fmt.Printf("%s", sub)
	return
	handler := PerformanceDetailHandler{
		exceptionMapping: make(map[int64][]interface{}),
	}
	for i := 0; i < 1000; i++ {
		handler.pushExceptionData(int64(i))
	}
	for k, v := range handler.exceptionMapping {
		fmt.Println(k)
		fmt.Printf("%+v\n", v)
	}
	//fmt.Printf("%+v", handler)
	return
	fmt.Println(1<<50)
	fmt.Println(0xF)
	return
	var a int64 = 1001
	fmt.Println(a/10)
	return
	var t1 T1
	var to interface{}
	to = t1
	t2 := T2{1, "2"}
	var f interface{}
	f = t2
	err := copier.Copy(&to, f)
	if err != nil {
		fmt.Println(err)
		return
	}
	t := to.(T1)
	fmt.Printf("%+v\n", t)
}

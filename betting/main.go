package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

type Model struct {
	ID        uint      `orm:"id" json:"id" gorm:"primaryKey"`
	CreatedAt time.Time `orm:"created_at" json:"created_at"`
	UpdatedAt time.Time `orm:"updated_at" json:"updated_at"`
	//DeletedAt gorm.DeletedAt 	 `orm:"deleted_at" json:"-" gorm:"index"` // 初始化
	//DeletedAt time.Time	 `orm:"deleted_at,omitempty" json:"-" gorm:"index"` // 正常应用
}

type BaseTask struct {
	Model
	PlatformId    uint   `orm:"platform_id" json:"platform_id" form:"platform_id" gorm:"column:platform_id;comment:平台id;type:bigint;size:19;"`
	ApplicationId uint   `orm:"application_id" json:"application_id" form:"application_id" gorm:"column:application_id;comment:应用id;type:bigint;size:19;"`
	UserId        uint   `orm:"user_id" json:"user_id" form:"user_id" gorm:"column:user_id;comment:用户id;type:bigint;size:19;"`
	Status        int8   `gorm:"comment:状态 0 未处理 1 已处理;type:smallint(8)"`
	Desc          string `gorm:"column:desc;comment:任务描述"`
}

type RechargeSuccessReport struct {
	BaseTask
	Amount uint `orm:"amount" json:"amount" form:"amount" gorm:"column:amount;comment:充值金额;type:int"`
}

type TaskMqBody struct {
	TaskEventType int    `json:"task_event_type,omitempty"`
	Msg           string `json:"msg" json:"msg,omitempty"`
}

func main() {
	fmt.Println(time.Now().Unix())
	rand.Seed(time.Now().Unix())
	intn := rand.Intn(1)
	time.Sleep(time.Duration(intn+1) * time.Second)

	fmt.Println(time.Now().Unix())
	//date := time.Now().AddDate(0, 0, 18)
	startTimeNum := time.Now().AddDate(0, 0, -18).Unix()
	currentNum := time.Now().Unix()
	remainSeconds := startTimeNum - currentNum
	fmt.Println(remainSeconds)
	return

	//sprintf := fmt.Sprintf(`sss:%s`, "hello")
	//fmt.Println(sprintf)
	//fmt.Printf("%T", sprintf)
	//return

	report := RechargeSuccessReport{
		BaseTask: BaseTask{
			PlatformId:    3,
			ApplicationId: 1001,
			UserId:        9,
			Desc:          "用户充值成功",
		},
		Amount: 1000,
	}
	strs, _ := json.Marshal(report)

	body := TaskMqBody{
		TaskEventType: 3,
		Msg:           string(strs),
	}
	content, _ := json.Marshal(body)
	fmt.Println(string(content))


	return
	fmt.Println("hello world test")
	fmt.Println("hello world tttt")
	bettingStr := `{"agent_id": 1, "application_id": 1001, "bucket": 0, "channel_id": 3,"epoch": 1649823085625,"gamelog": {"account":"xcv1","activity_bonus":0,"agent_id":1,"all_bets":40000,"all_bonus":0,"all_dpay":-40000,"channel_id":3,"client_ip":"172.16.130.112","currency_id":2,"detail":{},"device_id":1,"effect_dama":40000,"end_chips":20080,"epoch":1649823085625,"game_id":3000602,"gt_id":30,"jp_bonus":0,"jp_contri":0,"log_id":9571368719044609,"md5":"","play_id":0,"revenue":0,"room_id":0,"room_type":1,"start_chips":60080,"third_account":"36563941067841537","third_type":1,"top_agent_id":1,"type_id":1,"user_id":36563941067841537,"user_name":"xcv1"},"gtid": 30,"key": "","parent_id": 0,"platform_id": 3}`
	var betting interface{}

	if err := json.Unmarshal([]byte(bettingStr), &betting); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v", betting)
}

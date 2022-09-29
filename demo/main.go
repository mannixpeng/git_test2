package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Model struct {
	ID        uint      `orm:"id" json:"id" gorm:"primaryKey"`
	CreatedAt time.Time `orm:"created_at" json:"created_at"`
	UpdatedAt time.Time `orm:"updated_at" json:"updated_at"`
	//DeletedAt gorm.DeletedAt 	 `orm:"deleted_at" json:"-" gorm:"index"` // 初始化
	//DeletedAt time.Time	 `orm:"deleted_at,omitempty" json:"-" gorm:"index"` // 正常应用
}

type UserTaskRecord struct {
	Model
	UserId                int       `json:"user_id" gorm:""`
	PersonalTaskId        int       `json:"personal_task_id" gorm:"个人任务ID"`
	ApplicationId         int       `json:"application_id" gorm:"comment:应用ID"`
	PlatformId            int       `json:"platform_id" gorm:"comment：平台ID"`
	TaskCategory          int       `json:"task_category" gorm:"comment:任务分类"`
	TaskSubCategory       int       `json:"task_sub_category" gorm:"comment:任务子分类"`
	Sort                  int8      `json:"sort" gorm:"comment: 排序"`
	CanRepeat             int8      `json:"can_repeat" gorm:"comment: 是否重复任务"`
	ParentTaskId          int       `json:"parent_task_id" gorm:"comment: 前置任务"`
	TriggerType           int       `json:"trigger_type" gorm:"comment: 触发类型 1 统计游戏局数 2 统计游戏投注 3 统计游戏输赢 4 分享链接 5 免费次数 6 统计大奖次数 7 免费中累计赢取 8 充值成功 9 邀请成功 10 累计积分"`
	AutoReceive           int8      `json:"auto_receive" gorm:"comment:是否自动领取 1 自动 2 手动"`
	RewardsCollection     string    `json:"rewards_collection" gorm:"comment:奖励集合"`
	ReferenceStep         string    `json:"reference_step" gorm:"comment:关联步骤"`
	Status                int8      `json:"status" gorm:"comment:任务状态"`
	LimitedScope          int       `json:"limited_scope" gorm:"comment: 任务适应范围 0 所有游戏 1 facebook 2 游戏类型 3 其他指特定游戏 "`
	LimitedScopeValue     int       `json:"limited_scope_value" gorm:"comment: 任务范围值 当limited_scope为2则是游戏类型ID 为3时为特定游戏ID"`
	FinishedCondition     int64     `json:"finished_condition" gorm:"comment: 完成条件"`
	CurrentTaskValue      int64     `json:"current_task_value" gorm:"comment: 当前任务进行状态"`
	StatisticalConditions int64     `json:"statistical_conditions" gorm:"comment: 统计门槛"`
	TaskDeadline          time.Time `json:"task_deadline" gorm:"comment:任务截至时间"`
}

func main() {
	//time.FixedZone()
	zone, offset := time.Now().Zone()
	fmt.Println(zone, offset)
	return
	//cron.ParseStandard()

	//s := `{"agent_id": 1,"application_id": 1001, "bucket": 1,"channel_id": 3,"epoch": 1649823085625,"gamelog": "{\"account\":\"xcv1\",\"activity_bonus\":0,\"agent_id\":1,\"all_bets\":40000,\"all_bonus\":0,\"all_dpay\":-40000,\"channel_id\":3,\"client_ip\":\"172.16.130.112\",\"currency_id\":2,\"detail\":{},\"device_id\":1,\"effect_dama\":40000,\"end_chips\":20080,\"epoch\":1649823085625,\"game_id\":3000602,\"gt_id\":30,\"jp_bonus\":0,\"jp_contri\":0,\"log_id\":9571368719044609,\"md5\":\"\",\"play_id\":0,\"revenue\":0,\"room_id\":0,\"room_type\":1,\"start_chips\":60080,\"third_account\":\"36563941067841537\",\"third_type\":1,\"top_agent_id\":1,\"type_id\":1,\"user_id\":36563941067841537,\"user_name\":\"xcv1\"}","gtid": 30,"parent_id": 0, "platform_id": 2}`
	//b := []byte(s)
	record := UserTaskRecord{}
	marshal, _ := json.Marshal(record)
	fmt.Println(string(marshal))
}

package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type PlatformReportDay struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	// 核心关系
	PlatformId    int `orm:"platform_id" json:"platform_id" form:"platform_id" gorm:"index:idx_platform_id;column:platform_id;comment:平台id;type:bigint;size:19;"`
	ApplicationId int `orm:"application_id" json:"application_id" form:"application_id" gorm:"index:idx_application_id;column:application_id;comment:应用id;type:bigint;size:19;"`
	AgentId       int `orm:"agent_id" json:"agent_id" form:"agent_id" gorm:"index:idx_agent_id;column:agent_id;comment:代理id;type:bigint;size:19;"`
	ChannelId     int `orm:"channel_id" json:"channel_id" form:"channel_id" gorm:"index:idx_channel_id;column:channel_id;comment:渠道id;type:bigint;size:19;"`
	Month         int `orm:"month" json:"month" form:"month" gorm:"index:idx_month;column:month;comment:月份;type:int;"`
	Date          int `orm:"date" json:"date" form:"date" gorm:"index:idx_date;column:date;comment:日期;type:int;"`

	// 安装类
	InstallCount int `orm:"install_count" json:"install_count" form:"install_count" gorm:"column:install_count;comment:当日安装总人数;type:int;"`

	// // 用户状态信息类
	// YesterdayBalance int `orm:"yesterday_balance" json:"yesterday_balance" form:"yesterday_balance" gorm:"column:yesterday_balance;comment:昨日总余额;type:bigint(20);"`
	// TodayBalance     int `orm:"today_balance" json:"today_balance" form:"today_balance" gorm:"column:today_balance;comment:当日总余额;type:bigint(20);"`

	// 登陆注册类
	LoginCount            int `orm:"login_count" json:"login_count" form:"login_count" gorm:"column:login_count;comment:当日登录总人数;type:int;"`
	RegisterBindCount     int `orm:"register_bind_count" json:"register_bind_count" form:"register_bind_count" gorm:"column:register_bind_count;comment:当日注册绑定总人数;type:int;"`
	RegisterCount         int `orm:"register_count" json:"register_count" form:"register_count" gorm:"column:register_count;comment:当日注册总人数;type:int;"`
	RegisterRechargeCount int `orm:"register_recharge_count" json:"register_recharge_count" form:"register_recharge_count" gorm:"column:register_recharge_count;comment:当日注册充值总人数;type:int;"`

	// 充值&出款&投注类
	RechargeWithdrawRate    int `orm:"recharge_withdraw_rate" json:"recharge_withdraw_rate" form:"recharge_withdraw_rate" gorm:"column:recharge_withdraw_rate;comment:当日充值提现比;type:bigint(20);"`
	RechargeWithdrawSub     int `orm:"recharge_withdraw_sub" json:"recharge_withdraw_sub" form:"recharge_withdraw_sub" gorm:"column:recharge_withdraw_sub;comment:当日充值提现差;type:bigint(20);"`
	RechargeWithdrawSubRate int `orm:"recharge_withdraw_sub_rate" json:"recharge_withdraw_sub_rate" form:"recharge_withdraw_sub_rate" gorm:"column:recharge_withdraw_sub_rate;comment:当日充值提现差比;type:bigint(20);"`
	RechargeBettingRate     int `orm:"recharge_betting_rate" json:"recharge_betting_rate" form:"recharge_betting_rate" gorm:"column:recharge_betting_rate;comment:当日充值投注比;type:bigint(20);"`

	// 充值类
	RechargeCount               int `orm:"recharge_count" json:"recharge_count" form:"recharge_count" gorm:"column:recharge_count;comment:当日充值总人数;type:int;"`
	RechargeAmount              int `orm:"recharge_amount" json:"recharge_amount" form:"recharge_amount" gorm:"column:recharge_amount;comment:当日充值总金额;type:bigint(20);"`
	Arppu                       int `orm:"arppu" json:"arppu" form:"arppu" gorm:"column:arppu;comment:当日付费总用户平均收益(ARPPU);type:bigint(20);"`
	FirstRechargeCount          int `orm:"first_recharge_count" json:"first_recharge_count" form:"first_recharge_count" gorm:"column:first_recharge_count;comment:当日首充总人数;type:int;"`
	FirstRechargeAmount         int `orm:"first_recharge_amount" json:"first_recharge_amount" form:"first_recharge_amount" gorm:"column:first_recharge_amount;comment:当日首充总金额;type:bigint(20);"`
	MultipleRechargeCount       int `orm:"multiple_recharge_count" json:"multiple_recharge_count" form:"multiple_recharge_count" gorm:"column:multiple_recharge_count;comment:当日重充总人数;type:int;"`
	MultipleRechargeAmount      int `orm:"multiple_recharge_amount" json:"multiple_recharge_amount" form:"multiple_recharge_amount" gorm:"column:multiple_recharge_amount;comment:当日重充总金额;type:bigint(20);"`
	FirstMultipleRechargeCount  int `orm:"first_multiple_recharge_count" json:"first_multiple_recharge_count" form:"first_multiple_recharge_count" gorm:"column:first_multiple_recharge_count;comment:当日首充并重充总人数;type:int;"`
	FirstMultipleRechargeAmount int `orm:"first_multiple_recharge_amount" json:"first_multiple_recharge_amount" form:"first_multiple_recharge_amount" gorm:"column:first_multiple_recharge_amount;comment:当日首充并重充总金额;type:bigint(20);"`
	NewArppu                    int `orm:"new_arppu" json:"new_arppu" form:"new_arppu" gorm:"column:new_arppu;comment:当日付费新用户平均收益(ARPPU);type:bigint(20);"`

	BeforeMultipleRechargeCount  int `orm:"before_multiple_recharge_count" json:"before_multiple_recharge_count" form:"before_multiple_recharge_count" gorm:"column:before_multiple_recharge_count;comment:往日重充总人数;type:int;"`
	BeforeMultipleRechargeAmount int `orm:"before_multiple_recharge_amount" json:"before_multiple_recharge_amount" form:"before_multiple_recharge_amount" gorm:"column:before_multiple_recharge_amount;comment:往日重充总金额;type:bigint(20);"`
	OldArppu                     int `orm:"old_arppu" json:"old_arppu" form:"old_arppu" gorm:"column:old_arppu;comment:当日付费老用户平均收益(ARPPU);default:0;type:bigint(20);"`

	// 出款类
	WithdrawCount  int `orm:"withdraw_count" json:"withdraw_count" form:"withdraw_count" gorm:"column:withdraw_count;comment:当日总出款成功总人数;type:int;"`
	WithdrawAmount int `orm:"withdraw_amount" json:"withdraw_amount" form:"withdraw_amount" gorm:"column:withdraw_amount;comment:当日总出款成功总金额;type:bigint(20);"`

	// 投注&充值类
	BettingRechargeAmount int `orm:"betting_recharge_amount" json:"betting_recharge_amount" form:"betting_recharge_amount" gorm:"column:betting_recharge_amount;comment:当日充值并投注总金额;type:bigint(20);"`
	BettingRechargeCount  int `orm:"betting_recharge_count" json:"betting_recharge_count" form:"betting_recharge_count" gorm:"column:betting_recharge_count;comment:当日充值并投注总人数;type:int;"`
	BettingRechargeRate   int `orm:"betting_recharge_rate" json:"betting_recharge_rate" form:"betting_recharge_rate" gorm:"column:betting_recharge_rate;comment:当日投注充值比;type:bigint(20);"`

	WinLoseRechargeRate int `orm:"win_lose_recharge_rate" json:"win_lose_recharge_rate" form:"win_lose_recharge_rate" gorm:"column:win_lose_recharge_rate;comment:当日输赢充值比;type:bigint(20);"`
	RevenueRechargeRate int `orm:"revenue_recharge_rate" json:"revenue_recharge_rate" form:"revenue_recharge_rate" gorm:"column:revenue_recharge_rate;comment:当日营收充值比;type:bigint(20);"`
	WinLoseBettingRate  int `orm:"win_lose_betting_rate" json:"win_lose_betting_rate" form:"win_lose_betting_rate" gorm:"column:win_lose_betting_rate;comment:当日输赢投注比;type:bigint(20);"`
	BonusRechargeRate   int `orm:"bonus_recharge_rate" json:"bonus_recharge_rate" form:"bonus_recharge_rate" gorm:"column:bonus_recharge_rate;comment:当日彩金充值比;type:bigint(20);"`
	WinLoseTaxRate      int `orm:"win_lose_tax_rate" json:"win_lose_tax_rate" form:"win_lose_tax_rate" gorm:"column:win_lose_tax_rate;comment:当日输赢税收比;type:bigint(20);"`

	BettingCount     int     `orm:"betting_count" json:"betting_count" form:"betting_count" gorm:"column:betting_count;comment:当日投注总人数;type:int;"`
	BettingAmount    int     `orm:"betting_amount" json:"betting_amount" form:"betting_amount" gorm:"column:betting_amount;comment:当日总有效打码;type:bigint(20);"`
	BettingAvgAmount int     `orm:"betting_avg_amount" json:"betting_avg_amount" form:"betting_avg_amount" gorm:"column:betting_avg_amount;comment:当日人均投注金额;type:bigint(20);"`
	WinLoseAmount    float64 `orm:"win_lose_amount" json:"win_lose_amount" form:"win_lose_amount" gorm:"column:win_lose_amount;comment:当日总输赢金额;type:bigint(20);"`
	TaxAmount        int     `orm:"tax_amount" json:"tax_amount" form:"tax_amount" gorm:"column:tax_amount;comment:当日税收总金额;type:bigint(20);"`
	RevenueAmount    int     `orm:"revenue_amount" json:"revenue_amount" form:"revenue_amount" gorm:"column:revenue_amount;comment:当日平台营收金额;type:bigint(20);"`

	// 新增
	AllBets       int `orm:"all_bets" json:"all_bets" form:"all_bets" gorm:"column:all_bets;comment:投注金额;type:bigint;size:19;"`     //  总下注
	AllBonus      int `orm:"all_bonus" json:"all_bonus" form:"all_bonus" gorm:"column:all_bonus;comment:中奖金额;type:bigint;size:19;"` //  总奖金
	JpBonus       int `orm:"jp_bonus" json:"jp_bonus" form:"jp_bonus" gorm:"column:jp_bonus;comment:JP奖金;type:bigint;size:19;"`
	JpContri      int `orm:"jp_contri" json:"jp_contri" form:"jp_contri" gorm:"column:jp_contri;comment:JP贡献;type:bigint;size:19;"`
	BettingNumber int `orm:"betting_number" json:"betting_number" form:"betting_number" gorm:"column:betting_number;comment:投注次数;type:bigint;size:19;"`

	// 返佣、彩金、返水、返奖类
	BonusAmount      int     `orm:"bonus_amount" json:"bonus_amount" form:"bonus_amount" gorm:"column:bonus_amount;comment:当日总发放彩金总金额;type:bigint(20);"`
	BonusCount       int     `orm:"bonus_count" json:"bonus_count" form:"bonus_count" gorm:"column:bonus_count;comment:当日彩金发放总人数;type:int;"`
	CommissionAmount int     `orm:"commission_amount" json:"commission_amount" form:"commission_amount" gorm:"column:commission_amount;comment:当日总发放佣金总金额;type:bigint(20);"`
	CommissionCount  int     `orm:"commission_count" json:"commission_count" form:"commission_count" gorm:"column:commission_count;comment:当日佣金发放总人数;type:int;"`
	ChipsAmount      int     `orm:"chips_amount" json:"chips_amount" form:"chips_amount" gorm:"column:chips_amount;comment:当日总发放返水总金额;type:bigint(20);"`
	RebateAmount     float64 `orm:"rebate_amount" json:"rebate_amount" form:"rebate_amount" gorm:"column:rebate_amount;comment:当日总返奖金额(玩家输赢+税收);type:bigint(20);"`
}

func (p *PlatformReportDay) TableName() string {
	return "report_platform_day"
}

//func (p *PlatformReportDay) Database() string {
//	return "usercenter"
//}

type T struct {
	A int64
}

func Test(ctx context.Context) {
	var m map[string]string
	m["ss"] = "ss"
	return

	for {
		select {
		case <-ctx.Done():
			fmt.Println("cancel")
			if ctx.Err() != nil {
				fmt.Println(ctx.Err())
				return
			}
			fmt.Println(ctx.Deadline())
			return
		default:
			func() {
				time.Sleep(20 * time.Second)
				fmt.Println("end")
			}()

		}
	}

}

var wg sync.WaitGroup

type DailyBettingPending struct {
	gorm.Model
	LogId  int64  `json:"log_id" gorm:"column:log_id;comment:日志id;type:bigint;size:19;unique_index:uniq_idx_log_id"`
	Data   string `json:"data" gorm:"type:text"`
	Status int    `json:"status" gorm:"comment:0 待处理 1 投注统计已完成 2 投注落地中 3 投注落地已完成 4 每日统计处理中 5 每日统计已完成 6 月统计处理中;index:idx_status"`
}

func (i *DailyBettingPending) TableName() string {
	return "daily_betting_pending"
}

//func getPlatform() int {
//	rand.Seed(time.Now().UnixNano())
//	intn := rand.Intn(10)
//	if intn == 0 {
//		return intn+1
//	}
//	return intn
//}
//
//
//func getApplication() int {
//	rand.Seed(time.Now().UnixNano())
//	intn := rand.Intn(10)
//	if intn == 0 {
//		return intn+1
//	}
//	return intn
//}
type UserAgentCommission struct {
	Idx       uint      `orm:"idx" json:"idx" gorm:"primaryKey"`
	CreatedAt time.Time `orm:"created_at" json:"created_at"`
	UpdatedAt time.Time `orm:"updated_at" json:"updated_at"`

	Id                string      `orm:"id" json:"id" form:"id" gorm:"column:id;comment:结算日期;type:varchar(32);size:32;"`
	Date              int         `orm:"date" json:"date" form:"date" gorm:"column:date;comment:结算日期;type:bigint;size:19;"`
	PlatformId        uint        `orm:"platform_id" json:"platform_id" form:"platform_id" gorm:"column:platform_id;comment:平台id ;type:bigint;size:19;"`
	ApplicationId     uint        `orm:"application_id" json:"application_id" form:"application_id" gorm:"column:application_id;comment:所属应用;type:bigint;size:20;"`
	AgentId           uint        `orm:"agent_id" json:"agent_id" form:"agent_id" gorm:"column:agent_id;comment:代理ID;type:bigint;size:20;"`
	ChannelId         uint        `orm:"channel_id" json:"channel_id" form:"channel_id" gorm:"column:channel_id;comment:渠道ID;type:bigint;size:20;"`
	UserId            uint        `orm:"user_id" json:"user_id" form:"user_id" gorm:"column:user_id;comment:代理ID;type:bigint;size:20;"`
	Pid               int         `orm:"pid" json:"pid" form:"pid" gorm:"column:pid;comment:直属上级id;type:bigint;size:19;"`
	TeamMemberCount   int         `orm:"team_member_count" json:"team_member_count" form:"team_member_count" gorm:"column:team_member_count;comment:团队人数;type:bigint;size:19;default:0;"`
	DirectMemberCount int         `orm:"direct_member_count" json:"direct_member_count" form:"direct_member_count" gorm:"column:direct_member_count;comment:直属人数;type:bigint;size:19;default:0;"`
	Performance       int         `orm:"performance" json:"performance" form:"performance" gorm:"column:performance;comment:流水;type:bigint;size:19;default:0;"`
	TeamPerformance   int         `orm:"team_performance" json:"team_performance" form:"team_performance" gorm:"column:team_performance;comment:团队业绩;type:bigint;size:19;"`
	CommissionRate    int         `orm:"commission_rate" json:"commission_rate" form:"commission_rate" gorm:"column:commission_rate;comment:佣金比列;type:bigint;size:19;"`
	TeamCommission    int         `orm:"team_commission" json:"team_commission" form:"team_commission" gorm:"column:team_commission;comment:团队佣金;type:bigint;size:19;"`
	Commission        int         `orm:"commission" json:"commission" form:"commission" gorm:"column:commission;comment:代理佣金;type:bigint;size:19;"`
	Deduction         int         `orm:"deduction" json:"deduction" form:"deduction" gorm:"column:deduction;comment:扣除余额;type:smallint;"`
	TrueCommission    int         `orm:"true_commission" json:"true_commission" form:"true_commission" gorm:"column:true_commission;comment:实发佣金;type:bigint;size:19;"`
	DispatchType      int         `orm:"dispatch_type" json:"dispatch_type" form:"dispatch_type" gorm:"column:dispatch_type;comment:1:自动发放，2:手动发放;type:smallint;"`
	DispatchStatus    int         `orm:"dispatch_status" json:"dispatch_status" form:"dispatch_status" gorm:"column:dispatch_status;comment:结算状态（1是未结算，5是未发放，10是已发放，15是已取消）;gype:smallint;"`
	DispatchAt        *gtime.Time `orm:"dispatch_at" json:"dispatch_at" form:"dispatch_at" gorm:"column:dispatch_at;comment:发放实际;type:datetime"`
	Comment           string      `orm:"comment" json:"comment" form:"comment" gorm:"column:comment;comment:;type:varchar(255);size:255;"`
}

func (m *UserAgentCommission) TableName() string {
	return "user_agent_commission"
}

func main() {
	//Test()
	//select {
	//
	//}
	//ctx, cancelFunc := context.WithTimeout(context.Background(), 1*time.Second)
	//Test(ctx)
	//cancelFunc()
	//select {}
	format := time.Unix(10000, 0).Format("20060102")
	fmt.Println(format)
	//// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:123456@tcp(172.16.130.206:3306)/usercenter?"
	c := logger.Default
	c.LogMode(logger.Info)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: c,
	})
	if err != nil {
		panic(err)
	}

	if err := db.AutoMigrate(&PlatformReportDay{}); err != nil {
		panic(err)
	}
	commission := PlatformReportDay{
		PlatformId:    2,
		ApplicationId: 1002,
	}
	row, err := g.DB().Table(&commission).Save(commission)
	if err != nil {
		panic(err)
	} else {
		id, err := row.LastInsertId()
		if err != nil {
			panic("LastInsertId failed, err: " + err.Error())
			return
		} else {
			fmt.Println("id: ", id)
		}
		fmt.Printf("rowLastId: %d\n", id)
		affected, err := row.RowsAffected()
		if err != nil {
			panic("RowsAffected failed, err: " + err.Error())
		} else {
			panic(fmt.Sprintf("RowsAffected, %d", affected))
		}
	}
	//fmt.Println(id, err)
	//_, errs := g.DB().Model(&commission).Where("id = ?", commission.Id).Save(&commission)
	//fmt.Println(errs)

	//for i := 0; i < 1000; i++ {
	//	d := i
	//	go test(d)
	//}

	// log_id:  18574726104317954
	// platform_id: 3
	// application_id: 1001
	// channel_id: 10611001
	//user_id: 3485
	//month: 202204
	// date: 20220402

	//data := `{"id":0,"created_at":"0001-01-01T00:00:00Z","updated_at":"0001--01T00:00:00Z","log_id":%d,"platform_id":%d,"application_id":%d,"channel_id":%d,"agent_id":0,"user_id":%d,"month":%d,"date":%d,"is_active":0,"betting_time":0,"betting_out_time":0,"betting_duration":0,"betting_amount":250,"betting_count":0,"win_lose_amount":150,"tax_amount":0,"rebate_amount":150,"income_rate":0,"win_lose_tax_rate":0,"win_lose_betting_rate":0,"all_bets":250,"all_bonus":400,"jp_bonus":0,"jp_contri":0,"betting_number":0}'"`
	//log_id := 18574726104317954
	//platform_id :=
	//for i := 0; i < 10000000; i++ {
	//	fmt.Sprintf(data, )
	//}
	//pending := DailyBettingPending{}

	//s := "INSERT INTO `task`.`daily_betting_pending` (`created_at`, `updated_at`, `deleted_at`, `data`, `status`, `log_id`) VALUES ('2022-04-02 03:38:49', '2022-04-13 16:13:36', NULL, '{"id":0,"created_at\\\":\\\"0001-01-01T00:00:00Z\\\",\\\"updated_at\\\":\\\"0001-01-01T00:00:00Z\\\",\\\"log_id\\\":18574726104317954,\\\"platform_id\\\":3,\\\"application_id\\\":1001,\\\"channel_id\\\":10611001,\\\"agent_id\\\":0,\\\"user_id\\\":3485,\\\"month\\\":202204,\\\"date\\\":20220402,\\\"is_active\\\":0,\\\"betting_time\\\":0,\\\"betting_out_time\\\":0,\\\"betting_duration\\\":0,\\\"betting_amount\\\":250,\\\"betting_count\\\":0,\\\"win_lose_amount\\\":150,\\\"tax_amount\\\":0,\\\"rebate_amount\\\":150,\\\"income_rate\\\":0,\\\"win_lose_tax_rate\\\":0,\\\"win_lose_betting_rate\\\":0,\\\"all_bets\\\":250,\\\"all_bonus\\\":400,\\\"jp_bonus\\\":0,\\\"jp_contri\\\":0,\\\"betting_number\\\":0}', 1, 18574726104317954);\n"

	//pending := DailyBettingPending{}
	//t := db.Model(pending).Exec(s)
	//if t.Error != nil {
	//	fmt.Println("1111", t.Error)
	//}
	//time.Sleep(100 * time.Second)
	//c := db.Create(&day)
	//fmt.Println(c.Error)
	//go func(db2 *gorm.DB) {
	//	db.Transaction(func(tx *gorm.DB) error {
	//		var count int64
	//		tx.Model(&day).Where("month=?", 202204).Clauses(clause.Locking{Strength: "UPDATE"}).Count(&count)
	//		fmt.Println("111   ",count)
	//
	//		time.Sleep(20*time.Second)
	//		if count == 0 {
	//			tx.Model(&day).Create(&day)
	//		}
	//		return nil
	//	})
	//}(db)
	//var count int64
	//db.Transaction(func(tx *gorm.DB) error {
	//	_ = tx.Model(&day).Where("month=?", 202204).Clauses(clause.Locking{Strength: "UPDATE"}).Count(&count)
	//	fmt.Println("222222        ", count)
	//	time.Sleep(10* time.Second)
	//	if count == 0 {
	//		tx.Model(&day).Create(&day)
	//	}
	//	return nil
	//})
	//
	//
	//time.Sleep(20* time.Second)
	//fmt.Println(c.Error)
}

func test(i int) {
	day := PlatformReportDay{
		Month:      202204,
		PlatformId: i,
	}
	//go func() {
	//	g.DB().Transaction(context.Background(), func(ctx context.Context, tx *gdb.TX) error {
	//		var d PlatformReportDay
	//		tx.Model(day.TableName()).LockUpdate().Where(g.Map{"month": 202204}).Fields("month").Struct(&day)
	//		fmt.Println("g1: ", d.ID)
	//		fmt.Println("g1: ", d.Month)
	//		if d.ID == 0 {
	//			id, _ := tx.Model(day.TableName()).Save(&day)
	//			fmt.Println(id)
	//		}
	//		time.Sleep(1 * time.Second)
	//		return nil
	//	})
	//}()

	g.DB().Transaction(context.Background(), func(ctx context.Context, tx *gdb.TX) error {
		var d PlatformReportDay
		tx.Model(day.TableName()).LockUpdate().Where(g.Map{"month": 202204}).Fields("month").Struct(&day)
		fmt.Println("g2: ", d.ID)
		fmt.Println("g2: ", d.Month)
		if d.ID == 0 {
			id, _ := tx.Model(day.TableName()).Save(&day)
			fmt.Println(id)
		}
		time.Sleep(2 * time.Second)
		return nil
	})
}

func main3() {
	//rsp1 := "{Code:1001 Msg:success Data:}"
	//var b, _ = json.Marshal(rsp1)
	//err := json.Unmarshal(b, &rsp)
}

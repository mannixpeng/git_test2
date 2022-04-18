package main

import (
	"encoding/json"
	"fmt"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"math/rand"
	"time"
)



type User struct {
	gorm.Model
	Age int64 `gorm:"type:bigint(19);size:15"`
	Name string `gorm:"unique"`
	//CompanyID int
	//Company Company
	//CreditCards []CreditCard
	Amount decimal.Decimal `gorm:"type:decimal(10,2)"`
}



type Company struct {
	ID int
	Name string
	Code string
}

type CreditCard struct {
	gorm.Model
	Number string
	UserID uint
}

type User111 struct {
	PlatformId    int `json:"platform_id"`
	ApplicationId int `json:"application_id"`
	UserId        int `json:"user_id"`
}

type D struct {
	Data string `json:"data"`
	LogId int64 `json:"log_id"`
}
func main() {
	rand.Seed(time.Now().UnixNano())
	for i:=0;i < 100;i++ {
		n := rand.Int31n(10)
		fmt.Println(n)
	}
	return
	var d D
	d.Data = `{"id":0,"created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z","platform_id":3,"application_id":1001,"channel_id":10611001,"agent_id":0,"user_id":3485,"month":202204,"date":20220401,"is_active":0,"betting_time":0,"betting_out_time":0,"betting_duration":0,"betting_amount":250,"betting_count":0,"win_lose_amount":-250,"tax_amount":0,"rebate_amount":-250,"income_rate":0,"win_lose_tax_rate":0,"win_lose_betting_rate":0,"all_bets":250,"all_bonus":0,"jp_bonus":0,"jp_contri":0,"betting_number":0}`
	d.LogId = 1
	bytes, _ := json.Marshal(d)
	fmt.Println(string(bytes))
	return

	ticker := time.NewTicker(5 * time.Second)
	for {
		select {
		case <-ticker.C:
			fmt.Println("xxxxx")
		}
	}

	u := User111{
		PlatformId:    3,
		ApplicationId: 1001,
		UserId:        1,
	}
	marshal, _ := json.Marshal(u)
	fmt.Println(string(marshal))
	return
	//dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	//db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//strs := "2022-01-06 01:00:00"
	//parse, _ := time.ParseInLocation("2006-01-02 15:04:05", strs, time.Local)
	//fmt.Println(parse)
	duration, err := time.ParseDuration("-1h")
	if err != nil {
		panic(err)
	}
	preTime := time.Now().Add(duration)

	day := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println(day)
	hour := time.Now().Sub(preTime).Hours()
	fmt.Println(hour)
	sub := timeSub(time.Now().Local(), preTime)
	fmt.Println(sub)
	//now := time.Now()
	//hours := now.Sub(parse).Hours()
	//fmt.Println(hours)
}

func timeSub(t1, t2 time.Time) int {
	t1 = t1.UTC().Truncate(24 * time.Hour)
	t2 = t2.UTC().Truncate(24 * time.Hour)
	return int(t1.Sub(t2).Hours() / 24) % 7
}
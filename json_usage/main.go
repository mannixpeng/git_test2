package main

import (
	"fmt"
	"time"
)

type ReliefPaymentActivityExtraField struct {
	GiveAwayAmount      int `json:"give_away_amount"`
	AuditMultiple       int `json:"audit_multiple"`
	TriggerAmount       int `json:"trigger_amount"`
	GiveAwayNumberTimes int `json:"give_away_number_times"`
}

type ReChargeExtraField struct {
	AmountInterval string `json:"amount_interval"`
	GiveAwayRatio  int    `json:"give_away_ratio"`
	AuditMultiple  int    `json:"audit_multiple"`
}

type SignInExtraField struct {
	Day            int `json:"day"`
	GiveAwayAmount int64 `json:"give_away_amount"`
	AuditMultiple  int `json:"audit_multiple"`
}

func main() {
	//field := ReliefPaymentActivityExtraField{
	//	GiveAwayAmount:      18000,
	//	AuditMultiple:       10,
	//	TriggerAmount:       100,
	//	GiveAwayNumberTimes: 2,
	//}
	//marshal, _ := json.Marshal(field)
	//fmt.Println(string(marshal))
	//
	//var results []ReChargeExtraField
	//for i:= 0; i < 3; i++ {
	//	results = append(results, ReChargeExtraField{
	//		AmountInterval: fmt.Sprintf("%d-%d", i*100, (i+1)*100),
	//		GiveAwayRatio:  1,
	//		AuditMultiple:  10,
	//	})
	//}
	//d, _ := json.Marshal(results)
	//fmt.Println(string(d))

	//var results []SignInExtraField
	//for i:=0; i < 7;i++ {
	//	results = append(results, SignInExtraField{
	//		Day:            i,
	//		GiveAwayAmount: int64(i*1000),
	//		AuditMultiple:  10,
	//	})
	//}
	//d, _ := json.Marshal(results)
	//fmt.Println(string(d))
	//var a interface{}
	//err := json.Unmarshal(d, &a)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(a)
	duration, err := gtime.ParseDuration("6d")
	fmt.Println(duration, err)

}

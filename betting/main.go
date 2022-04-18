package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Println("hello world tttt")
	bettingStr := `{"agent_id": 1, "application_id": 1001, "bucket": 0, "channel_id": 3,"epoch": 1649823085625,"gamelog": {"account":"xcv1","activity_bonus":0,"agent_id":1,"all_bets":40000,"all_bonus":0,"all_dpay":-40000,"channel_id":3,"client_ip":"172.16.130.112","currency_id":2,"detail":{},"device_id":1,"effect_dama":40000,"end_chips":20080,"epoch":1649823085625,"game_id":3000602,"gt_id":30,"jp_bonus":0,"jp_contri":0,"log_id":9571368719044609,"md5":"","play_id":0,"revenue":0,"room_id":0,"room_type":1,"start_chips":60080,"third_account":"36563941067841537","third_type":1,"top_agent_id":1,"type_id":1,"user_id":36563941067841537,"user_name":"xcv1"},"gtid": 30,"key": "","parent_id": 0,"platform_id": 3}`
	var betting interface{}

	if err := json.Unmarshal([]byte(bettingStr), &betting); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v", betting)
}


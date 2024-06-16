package modules

import (
	"DBTools/conf"
	"DBTools/utils"
	"fmt"
	"github.com/jedib0t/go-pretty/table"
	"log"
	"os"
	"time"
)

var users = make([]string, 0)

func HikInfoGet() {
	fmt.Println(utils.ColorPrint(1, "******************User Info******************"))
	queryUser()
	fmt.Println(utils.ColorPrint(1, "******************Org Info******************"))
	quertOrg()
	fmt.Println(utils.ColorPrint(1, "******************Role Info******************"))
	quertRole()
	fmt.Println(utils.ColorPrint(1, "******************Region Info******************"))
	quertRegion()
	if conf.DbS.ChangePWD {
		for index, user := range users {
			fmt.Printf("%d:%s\n", index, user)
		}
		index := -1
		fmt.Printf("user index you want change:")
		fmt.Scanf("%d\n", &index)
		ExchangePassword(users[index])
	}
}
func ExchangePassword(user string) {
	rows, err := conf.DbS.Db.Query(fmt.Sprintf("SELECT user_pwd,salt,pwd_expire_time FROM irds_irdsdb.public.tb_user where user_name = '%s';", user))
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	origin_user_pwd := ""
	origin_salt := ""
	origin_pwd_expire_time := ""
	exchage_user_pwd := "57b303e3875c0e6834e388912a36215e044872f45b0d5bc01ef9eb47267a8292"
	exchange_expire_time := time.Now().AddDate(0, 0, 7)
	exchage_salt := "8e8c4210822e51efc34904279f8d716ce9a6c3f76f76d58690898eb08533ea76"
	for rows.Next() {
		err = rows.Scan(&origin_user_pwd, &origin_salt, &origin_pwd_expire_time)
	}
	fmt.Println(origin_user_pwd)
	fmt.Println(origin_salt)
	fmt.Println(origin_pwd_expire_time)
	query := `
        UPDATE irds_irdsdb.public.tb_user
        SET 
            user_pwd = $1,
            salt = $2,
            pwd_expire_time = $3
        WHERE 
            user_name = $4
    `
	_, err = conf.DbS.Db.Exec(query, exchage_user_pwd, exchage_salt, exchange_expire_time, user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Exchange Success\nUser %s's Password Change To hik12345+\n", user)
	fmt.Printf("End?")
	fmt.Scanf("\n")
	_, err = conf.DbS.Db.Exec(query, origin_user_pwd, origin_salt, origin_pwd_expire_time, user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("User Info Restore.")
}
func queryUser() {
	rows, err := conf.DbS.Db.Query("select user_name,usergroup_name from irds_irdsdb.public.tb_user;")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"USER_NAME", "USER_GROUP_NAME"})
	for rows.Next() {
		var username string
		var user_group_name string
		err = rows.Scan(&username, &user_group_name)
		users = append(users, username)
		if err != nil {
			log.Fatal(err)
		}
		t.AppendRow(table.Row{username, user_group_name})
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	t.Render()
	//fmt.Println("USER_NAME\t\tUSER_GROUP_NAME")
	// 遍历查询结果
	//for rows.Next() {
	//	var username string
	//	var user_group_name string
	//	err = rows.Scan(&username, &user_group_name)
	//	users = append(users, username)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	fmt.Printf("%s\t\t%s\n", username, user_group_name)
	//}
	//err = rows.Err()
	//if err != nil {
	//	log.Fatal(err)
	//}
}
func quertOrg() {
	rows, err := conf.DbS.Db.Query("select org_id,org_name,org_index_code from irds_irdsdb.public.tb_org;")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"ORG_ID", "ORG_NAME", "ORG_INDEX_CODE"})
	// 遍历查询结果
	for rows.Next() {
		var org_id int
		var org_name string
		var org_index_code string
		err = rows.Scan(&org_id, &org_name, &org_index_code)
		if err != nil {
			log.Fatal(err)
		}
		t.AppendRow(table.Row{org_id, org_name, org_index_code})
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	t.Render()
	//fmt.Println("ORG_ID\t\tORG_NAME\t\tORG_INDEX_CODE")
	//// 遍历查询结果
	//for rows.Next() {
	//	var org_id int
	//	var org_name string
	//	var org_index_code string
	//	err = rows.Scan(&org_id, &org_name, &org_index_code)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	fmt.Printf("%d\t\t%s\t\t%s\n", org_id, org_name, org_index_code)
	//}
	//err = rows.Err()
	//if err != nil {
	//	log.Fatal(err)
	//}
}
func quertRegion() {
	rows, err := conf.DbS.Db.Query("select region_name,parent_id from irds_irdsdb.public.tb_region;")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"REGION_NAME", "PARENT_ID"})
	// 遍历查询结果
	for rows.Next() {
		var region_name string
		var parent_id string
		err = rows.Scan(&region_name, &parent_id)
		if err != nil {
			log.Fatal(err)
		}
		t.AppendRow(table.Row{region_name, parent_id})
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	t.Render()
	//fmt.Println("REGION_NAME\t\tPARENT_ID")
	//// 遍历查询结果
	//for rows.Next() {
	//	var region_name string
	//	var parent_id string
	//	err = rows.Scan(&region_name, &parent_id)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	fmt.Printf("%s\t\t%s\n", region_name, parent_id)
	//}
	//err = rows.Err()
	//if err != nil {
	//	log.Fatal(err)
	//}
}
func quertRole() {
	rows, err := conf.DbS.Db.Query("select role_name,creator from irds_irdsdb.public.tb_role;")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"ROLE_NAME", "CREATOR"})
	// 遍历查询结果
	for rows.Next() {
		var role_name string
		var creator string
		err = rows.Scan(&role_name, &creator)
		if err != nil {
			log.Fatal(err)
		}
		t.AppendRow(table.Row{role_name, creator})
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	t.Render()
	//fmt.Println("ROLE_NAME\t\tCREATOR")
	//// 遍历查询结果
	//for rows.Next() {
	//	var role_name string
	//	var creator string
	//	err = rows.Scan(&role_name, &creator)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	fmt.Printf("%s\t\t%s\n", role_name, creator)
	//}
	//err = rows.Err()
	//if err != nil {
	//	log.Fatal(err)
	//}
}

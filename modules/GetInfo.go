package modules

import (
	"DBTools/conf"
	"DBTools/utils"
	"fmt"
	"github.com/jedib0t/go-pretty/table"
	"log"
	"os"
)

var users = make([]string, 0)

func InfoGet() {
	fmt.Println(utils.ColorPrint(1, "******************User Info******************"))
	queryUser()
	if conf.DbS.DataBase == "irds_irdsdb" {
		fmt.Println(utils.ColorPrint(1, "******************Org Info******************"))
		quertOrg()
	}
	fmt.Println(utils.ColorPrint(1, "******************Role Info******************"))
	quertRole()
	fmt.Println(utils.ColorPrint(1, "******************Region Info******************"))
	quertRegion()
}
func queryUser() {
	query := fmt.Sprintf("select user_name,usergroup_name from %s.public.tb_user;", conf.DbS.DataBase)
	rows, err := conf.DbS.Db.Query(query)
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
}
func quertOrg() {
	query := fmt.Sprintf("select org_id,org_name,org_index_code from %s.public.tb_org;", conf.DbS.DataBase)
	rows, err := conf.DbS.Db.Query(query)
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
}
func quertRegion() {
	query := fmt.Sprintf("select region_name from %s.public.tb_region;", conf.DbS.DataBase)
	rows, err := conf.DbS.Db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"REGION_NAME"})
	// 遍历查询结果
	for rows.Next() {
		var region_name string
		err = rows.Scan(&region_name)
		if err != nil {
			log.Fatal(err)
		}
		t.AppendRow(table.Row{region_name})
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	t.Render()
}
func quertRole() {
	query := fmt.Sprintf("select role_name,creator  from %s.public.tb_role;", conf.DbS.DataBase)
	rows, err := conf.DbS.Db.Query(query)
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
}

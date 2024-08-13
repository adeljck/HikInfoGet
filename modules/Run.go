package modules

import (
	"DBTools/conf"
	"DBTools/utils"
	"flag"
	"fmt"
	"os"
)

func Run() {
	flag.StringVar(&conf.DbS.ConfigFile, "f", "", "config file path.")
	flag.StringVar(&conf.DbS.Hostname, "h", "", "database host name.")
	flag.IntVar(&conf.DbS.Port, "P", 0, "database port.")
	flag.StringVar(&conf.DbS.Username, "u", "", "database user name.")
	flag.BoolVar(&conf.DbS.ChangePWD, "c", false, "change database password.")
	flag.StringVar(&conf.DbS.Password, "p", "", "database password.")
	flag.Parse()
	if conf.DbS.ConfigFile != "" {
		conf.DbS.ParseConfigFile()
	} else {
		if conf.DbS.Hostname == "" {
			fmt.Println(utils.ColorPrint(-1, "Wrong HostName."))
			return
		}
		if conf.DbS.Port <= 0 || conf.DbS.Port > 65535 {
			fmt.Println(utils.ColorPrint(-1, "Wrong Database Port."))
			return
		}
		if conf.DbS.Username == "" {
			fmt.Println(utils.ColorPrint(-1, "Wrong UserName."))
			return
		}
		if password, err := utils.DecryptData(conf.DbS.Password); err == nil {
			conf.DbS.Password = password
		}
	}
	err := conf.DbS.CreateConnection()
	if err != nil {
		fmt.Println(utils.ColorPrint(-1, "Connection Error."))
		return
	}
	fmt.Println(utils.ColorPrint(0, "Connect To DataBase Success."))

	InfoGet()
	if conf.DbS.ChangePWD && conf.DbS.DataBase == "irds_irdsdb" {
		for index, user := range users {
			fmt.Printf("%d:%s\n", index, user)
		}
		index := -1
		fmt.Printf("user index you want change:")
		fmt.Scanf("%d\n", &index)
		ChangePassword(users[index])
	} else {
		fmt.Println(utils.ColorPrint(-1, "Change PWD Not Support This Version."))
		os.Exit(0)
	}
}

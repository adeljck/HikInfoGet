package modules

import (
	"DBTools/conf"
	"fmt"
	"log"
	"time"
)

func ChangePassword(user string) {
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

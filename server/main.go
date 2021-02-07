
/* The sole purpose of this server is to keep track of the total number of wallets generated */

package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	. "./webframework"
	"log"
	"strconv"
	"os"
	"fmt"
)

var db *sql.DB

func main() {
	// mymysql.mysql.database.azure.com
	db, err := sql.Open("mysql", os.Args[1]+":"+ os.Args[2]+"@tcp("+os.Args[3]+")/bitcoinWallets?tls=true")
	if err != nil {
		log.Fatal(err)
	}

	app := Server()

	app.Get(`/`, func(req Req) UrlResp {
		// allow all origins
		req.AddHeader("Access-Control-Allow-Origin", "*")

		rows, err := db.Query("SELECT numOfWalletsGenerated FROM walletNum WHERE numOfWalletsGenerated > -1 LIMIT 1")
		if err != nil {
			return SendStr("0")
		}
		var numGen int
		for rows.Next() {
			err = rows.Scan(&numGen)
			if err != nil {
				return SendStr("0")
			}
			break
		}
		return SendStr(strconv.Itoa(numGen))
	})

	app.Get(`/new`, func(req Req) UrlResp {
		// allow all origins
		req.AddHeader("Access-Control-Allow-Origin", "*")

		insertion, err := db.Prepare("UPDATE walletNum SET numOfWalletsGenerated = numOfWalletsGenerated + 1 WHERE numOfWalletsGenerated > -1")
		if err != nil {
			log.Fatal(err)
		}
		insertion.Exec()
		return SendStr("done")
	})

	// listen on specificed port
	var port int
	port = 9990
	if len(os.Args) > 4 {
		port, err = strconv.Atoi(os.Args[4]) 
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println(port)
	app.Listen(port)
}
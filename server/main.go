
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

	db, err := sql.Open("mysql", "root:" + os.Args[1] + "@tcp(127.0.0.1:1800)/bitcoinWallets")
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
		
		fmt.Println(`new`)

		insertion, err := db.Prepare("UPDATE walletNum SET numOfWalletsGenerated = numOfWalletsGenerated + 1 WHERE numOfWalletsGenerated > -1")
		if err != nil {
			log.Fatal(err)
		}
		insertion.Exec()
		return SendStr("done")
	})

	app.Listen(9990)
}
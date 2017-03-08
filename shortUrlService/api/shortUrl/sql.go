package shortUrl

import (
	"database/sql"
)

var insertLongUrl *sql.Stmt
var selectLongUrlByShortUrlId *sql.Stmt

func init() {
	db, err := sql.Open("mysql", "root@unix(/tmp/mysql.sock)/qing")
	if err != nil {
		panic(err.Error())
	}

	insertLongUrl, err = db.Prepare("INSERT INTO defaultShortUrl (longurl) VALUES(?)")
	if err != nil {
		panic(err.Error())
	}

	selectLongUrlByShortUrlId, err = db.Prepare("SELECT longurl from defaultShortUrl WHERE shortUrlId = ?")
	if err != nil {
		panic(err.Error())
	}
}

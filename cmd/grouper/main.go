package main

import (
	"github.com/athom/grouper/web"
)

func main() {
	conf := `
{
	"storage_type": "mysql",
	"port": 7200,
	"mysql_settings": {
		"host": "localhost",
		"port": 52401,
		"username": "root",
		"password": "nopassword",
		"database": "grouper",
	}
}
`
	web.Run(conf)
}

## Grouper

### Description

This is a demo social network application for "Friends Management" purpose.
Providing APIs with features like "Friend", "Unfriend", "Block", "Receive Updates" etc.

### Features

- Relationship Connect, also called add firends.
- Single connect Relationship , also called subscribe, watch, or follow.
- Updates block, primarily use for block new feeds from certain target.
- List connected relationships, so called friends list.
- Get common relationships, used to get common friends from two targets.
- Get recipients for broadcasting one's news or messages. 

### Design 

The case base follows the [Clean Architecture](https://8thlight.com/blog/uncle-bob/2012/08/13/the-clean-architecture.html) and implement with Go.
 
![](https://8thlight.com/blog/assets/posts/2012-08-13-the-clean-architecture/CleanArchitecture-8d1fe066e8f7fa9c7d8e84c1a6b0e2b74b2c670ff8052828f4a7e73fcbbc698c.jpg)

Here is the relationship between go packages and the architecture layers.

- grouper: Entities&UseCases
- web: Controller & Web infrastrure
- storage: DB

### Requirements

- Setup mysql

```
mysql -h <your_host> -p<port> -u<username> -p<password> -e 'CREATE DATABASE IF NOT EXISTS grouper;'
```

- Go packages

install [govendoer](https://github.com/kardianos/govendor) first, then exec

```$xslt
govendor sync
```

### Run the App

#### 1. run from source code 

Go to project root, setup config in `$project_root/cmd/grouper/config.json`

```$xslt
{
	"storage_type": "mysql",
	"port": <app_port>,
	"mysql_settings": {
		"host": "your_fancy_host",
		"port": <db_port>,
		"username": "your_fancy_username",
		"password": "your_password",
		"database": "your_database"
	}
}
```

Run server.

```$xslt
go run cmd/grouper/main.go
```
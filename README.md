## Grouper


[![Build Status](https://api.travis-ci.org/athom/grouper.png?branch=master)](https://travis-ci.org/athom/grouper)


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

### Run the test

#### Prepare the DB
```$xslt
mysql -e 'CREATE DATABASE IF NOT EXISTS grouper_test;'
```

#### Trigger test cases

```$xslt
./test.sh
```

### Run the App

#### 1. Run from source code 

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

Run server with go:

```$xslt
go run cmd/grouper/main.go
```

#### 2. Run via docker

Go to deploy directory, make sure docker-compose is available, execute the script:

```$xslt
./run_server.sh
``` 
  

#### 3. Try it from demo site (temporary available)

```$xslt
curl -d '{"email": "andy@example.com"}' http://119.28.1.61:7200/v1/friends/find

curl -d '{"email": "john@example.com"}' http://119.28.1.61:7200/v1/friends/find

curl -d '{"friends": ["andy@example.com", "john@example.com"]}' http://119.28.1.61:7200/v1/friends/common
```

### TODO

- ~~Docker deployment~~
- More storage plugins support.
- Elegant converting between domain and storage.
- Refactor test, data seeds and make it more readable.
- Take care of the error messages.
- ~~CI~~
- CD
- Monitoring & Alerts.
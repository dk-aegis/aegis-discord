# Discord Bot 
Aegis Discord Bot

## Run
```BASH
git clone https://github.com/dk-aegis/aegis-discord.git
cd aegis-discord
go build 
./discord
```

## Config FILE
you need three config files under config folder
#### ./config/db_config.json
```json
{
    "type" : "mysql",
    "user" : "username",
    "password" : "password",
    "protocol" : "tcp",
    "port" : 3306,
    "host" : "127.0.0.1",
    "database" : "database's name"
}
```

#### ./config/discord.json
```json
{
    "guild_id" : "guild id",
    "welcome_channel_id" : "welcome channle id",
    "moderator_role_id" : "role id",
	"study_role_id"     : "role id",
	"graduate_role_id"  : "role id",
	"student_role_id"   : "role id"
}
```

#### ./config/token.json
```json
{
    "token" : "token value"
}
```

## 요구사항 
### Leveling System 

### Welcome System

### Attendance System
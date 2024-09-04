# Discord Bot 
Aegis Discord Bot

# nick 받아오려면 직접 설정되어 있어야함. !!!!!!!!! 
이걸 몇개월 동안 못찾았다니 말도안된다. 서버 내에서 직접 Nick 이 설정되어 있어야지 값이 나옴. 아니면 empty string 이 나오기 때문에.
Nick 기반으로 할거니까, Nickname 설정 변경 못하도록 권한을 줍시다.


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
	"student_role_id"   : "role id",
    "general_role_id"   : "role id",
    "executive_privilege": "secure code"
}
```

#### ./config/token.json
```json
{
    "token" : "token value",
    "guild_id" : "1223177363722862612"
}
```

```
db 생긴거
mysql 
discord
|
ㄴTables
 |
 ㄴattendacne     id attend_count last_seen conseq_count
 ㄴwallet         id money exp
```
## 요구사항 
### 

### 

### 
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
create .env file to manage configuration
#### .env
```
#bot config
BOT_TOKEN = 봇의 토큰
BOT_ID = 봇의 ID

#db config
DB_TYPE = mysql
DB_USER = username
DB_PSWD = password
DB_PROTOCOL = tcp
DB_PORT = 3306
DB_HOST = 127.0.0.1
DB_NAME = database's name

#guild ID
GUILD_ID = 서버의 ID

#channel ID
WELCOME_CHAN_ID = 환영인사 보낼 채널 ID

#role ID
MODROLE_ID = 봇의 특정 명령을 실행시킬 수 있는 권한을 가진 역할 ID (운영진)
STUDYROLE_ID = 스터디 관련된 역할 ID
GRADUROLE_ID = 졸업생 ID
STUDENTROLE_ID = 재학생 ID

```
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
    "moderator_role_id" : "관리자 권한 역할 id",
	"study_role_id"     : "스터디장 역할 id",
	"graduate_role_id"  : "졸업생 역할 id",
	"student_role_id"   : "재학생 역할 id"
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
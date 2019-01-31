# Golang echo boilerpalte

## 서버 자동으로 재실행 하기
### 패키지 받기
> go get github.com/codegangsta/gin
### 실행 
> run.bat

## 서버 실행하기
> go run server.go

## Package Manage
### dep 설치
> go get -u github.com/golang/dep/cmd/dep

### 패키지 추가 방법
> dep ensure -add [패키지 주소]

### 패키지 설치 하기
> dep ensure 

## Rest API 실행해 보기
`RestAPIs` 폴더의 파일을 실행 함.
VSCODE에서 [Rest-client](https://marketplace.visualstudio.com/items?itemName=humao.rest-client)를 사용해서 작성 됨

## Seed 데이터 생성하기
cli를 사용해서 seed 데이터를 생성 할 수 있습니다.
> go run server.go -seed=[all|users|tournaments] -count=[number]

도움말 보기
> go run server.go --help

사용자의 기본 패스트워드는 `test13@$` 입니다.


## 채널 채팅 샘플
> http://127.0.0.1:5000/ws




## VS Code Debug

`.vscode/launch.json` 파일을 아래와 같이 입력 해 줍니다.

```json
{
    // IntelliSense를 사용하여 가능한 특성에 대해 알아보세요.
    // 기존 특성에 대한 설명을 보려면 가리킵니다.
    // 자세한 내용을 보려면 https://go.microsoft.com/fwlink/?linkid=830387을(를) 방문하세요.
    "version": "0.2.0",
    "configurations": [
        {
            "name": "Launch",
            "type": "go",
            "request": "launch",
            "mode": "debug",
            "remotePath": "",
            "port": 2345,
            "host": "127.0.0.1",
            "program": "${workspaceRoot}/server.go",
            "env": {},
            "args": [],
            "showLog": true,
            "buildFlags": ["-v"]
        }
    ]
}
```



### Reference
#### Go
- [echo guide](https://echo.labstack.com/guide)
- [gorm](http://gorm.io/docs/)
- [govalidator](https://github.com/asaskevich/govalidator)

#### Project
- [An eSports tournament framework built with Laravel](https://github.com/g33kidd/bracket)
- [React components for rendering a tournament bracket](https://github.com/moodysalem/react-tournament-bracket)

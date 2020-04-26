# 概要
CleanArchitectureの勉強のために、Golang/GinでToDoサーバを構築してみた

# フォルダ構成
```
src
└── app
    ├── domain
    │   └── todo.go
    ├── infrastructure
    │   ├── dbhandler.go
    │   └── router.go
    ├── interfaces
    │   ├── controllers
    │   │   ├── context.go
    │   │   └── todo_controllers.go
    │   └── database
    │       ├── db_handler.go
    │       └── todo_repository.go
    ├── server.go
    └── usecase
        ├── todo_interactor.go
        └── todo_repository.go
```

# 参考
[Clean ArchitectureでAPI Serverを構築してみる](https://qiita.com/hirotakan/items/698c1f5773a3cca6193e)
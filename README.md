<center><h1>MoeNya's GoCtl HTTP Service Note</h1></center>

*By MoeNya with ❤️ Created 2024.08.09      Last modified 2024.08.09*

本文档记录了如何从零使用`goctl`生成一个http服务并且写一个小demo，这是`gozero`的一环。

## 安装goctl

`goctl`是`go-zero`内置的脚手架，它可以生成代码、文档、部署k8s yaml、dockerfile等。

你可以使用下面的代码来安装最新的`goctl`：

```cmd
go get -u install github.com/zeromicro/go-zero/tools/goct@latest

go install github.com/zeromicro/go-zero/tools/goct@latest

go install github.com/zeromicro/go-zero/tools/goct@vx.x.x
```

安装完成后可以在Terminal查看goctl的version。

## 使用生成一个HTTP服务

确保`goctl`安装成功后，可以使用`goctl`命令来生成代码了。

```cmd
goctl api new your_project_name
```

这里只是一个例子，同时命令中还可以有很多参数负责不同的功能，具体的可以去看`gozero docs`。

## 如何使用呢？

### 注册API

项目结构可以在官方文档中查看到，项目中有一个`.api`的文件，我们需要在这个文件中按照官方demo的写法取构建`type`，这个`type`对应的就是go的结构体。

接着在该文件中可以看到有一个`service`部分，在里面按照官方文档的写法注册api接口。

```go
type (
	// 通过uid对猫猫进行操作
	CatReq {
		Uid uint `path:"uid"`
	}
	CatReqPost {
		Uid uint `json:"uid"`
	}

	// 完整的猫猫数据结构体
	Cat {
		Uid  uint   `json:"uid"`
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	// 错误信息返回
	ErrResp {
		Err int `json:"err"`
	}
)

service crud_api-api {
	@handler GetCat
	get /query/:uid(CatReq) returns (Cat)
	
	@handler NewCat
	post /newcat (Cat) returns (ErrResp)
	
	@handler EditCat
	post /editcat (Cat) returns (ErrResp)
	
	@handler DeleteCat
	post /deletecat (CatReqPost) returns (ErrResp)
}
```

### 根据api文件生成HTTP代码

```cmd
goctl api go -api crud_api.api -dir .
```

其它参数和作用可以去`gozero docs`查看。

当运行该命令后，`goctl`将会生成`handler`、`logic`、`types`代码。

### 编写逻辑业务

最后我们可以在`logic`下注释提示的部分编写自己的代码逻辑。当然这个时候我们可以建立自己的其它工程结构，并且用在`logic`中。

```go
package logic

import (
	"context"
	"fmt"

	"crud_api/internal/svc"
	"crud_api/internal/types"
	"crud_api/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteCatLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteCatLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCatLogic {
	return &DeleteCatLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteCatLogic) DeleteCat(req *types.CatReqPost) (resp *types.ErrResp, err error) {
	resp = new(types.ErrResp)
	resp.Err = model.DeleteCat(req.Uid)
	if resp.Err != 200 {
		fmt.Println("删除猫猫信息失败")
		return
	}
	fmt.Println("删除猫猫信息成功")
	return
}
```

在后续文档中还会讲述如何使用gRPC，如何使用HTTP服务去调用RPC服务。

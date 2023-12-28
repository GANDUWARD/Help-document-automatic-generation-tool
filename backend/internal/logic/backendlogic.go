package logic

import (
	"context"
	"fmt"
	"os"

	"backend/internal/svc"
	"backend/internal/types"
	"path/filepath"

	"github.com/zeromicro/go-zero/core/logx"
)

type BackendLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBackendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BackendLogic {
	return &BackendLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BackendLogic) Backend(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	currentDir, _ := os.Getwd()
	// 构建目标文件夹路径
	targetDir := filepath.Join(currentDir, "taskarea")
	//获取当前目录下的所有文件和目录信息
	filepath.Walk(targetDir, func(path string, info os.FileInfo, err error) error {
		fmt.Println(path)      //打印path信息
		fmt.Println(info.Name) // 打印文件或目录名
		return nil
	})
	return
}

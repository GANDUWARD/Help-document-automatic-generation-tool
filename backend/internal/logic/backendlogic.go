package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

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
type Filenode struct {
	Name        string    `json:"name"`
	Path        string    `json:"path"`
	Firstchild  *Filenode `json:"firstchild"`
	Nextsibling *Filenode `json:"nextsibling"`
	IsDirectory bool      `json:"isDirectory"`
}

func NewBackendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BackendLogic {
	return &BackendLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}
func buildFileTree(rootPath string) (*Filenode, error) {
	var root Filenode

	err := filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		//忽略原目录
		if path == rootPath {
			return nil
		}
		// 构建节点
		node := &Filenode{
			Name:        info.Name(),
			Path:        path,
			IsDirectory: info.IsDir(),
		}

		// 将节点添加到树中
		root.addFileNode(node)

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &root, nil
}

func (node *Filenode) addFileNode(newNode *Filenode) {
	if node.Firstchild == nil {
		node.Firstchild = newNode
	} else {
		currentNode := node.Firstchild
		for currentNode.Nextsibling != nil {
			currentNode = currentNode.Nextsibling
		}
		currentNode.Nextsibling = newNode
	}
}
func GetFileName(s string) string {
	names := strings.Split(s, "\\")
	return names[len(names)-1]
}
func IsDirectory(s string) bool {
	for _, v := range s {
		if v == rune("."[0]) {
			return false
		}
	}
	return true
}
func (l *BackendLogic) Backend(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	currentDir, _ := os.Getwd()
	respon := new(types.Response)
	resp = respon
	// 构建目标文件夹路径
	targetDir := filepath.Join(currentDir, "taskarea")
	fileTree, err := buildFileTree(targetDir)
	if err != nil {
		fmt.Println("Error building file tree:", err)
		return
	}
	// 将树形结构转为JSON
	jsonData, err := json.MarshalIndent(fileTree.Firstchild, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}
	resp.Data = string(jsonData)
	fmt.Println(string(jsonData))
	/*

		//构建所有文件结构切片
		filelist := new([]Filenode)
		//获取当前目录下的所有文件和目录信息

		filepath.Walk(targetDir, func(path string, info os.FileInfo, err error) error {
			fmt.Println(path)      //打印path信息
			fmt.Println(info.Name) // 打印文件或目录名
			f := new(Filenode)
			f.Path = path
			f.Name = GetFileName(path)
			f.IsDirectory = IsDirectory(f.Name)
			*filelist = append(*filelist, *f)
			return nil
		})
	*/
	return
}
func (l *BackendLogic) BackendPOST(req *types.PostRequest) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	respon := new(types.Response)
	resp = respon
	//读取指定文件
	filepath := req.Path
	// 读取文件内容
	content, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}
	resp.Data = string(content)
	return
}
func (l *BackendLogic) BackendSAVE(req *types.SaveRequest) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	respon := new(types.Response)
	resp = respon
	//读取指定文件
	// 指定文件目录和文件名
	directory := req.Path
	fileName := req.Name

	// 指定文件内容
	fileContent := req.Content
	// 拼接文件路径
	filePath := filepath.Join(directory, fileName)

	// 写入文件内容
	err = os.WriteFile(filePath, []byte(fileContent), os.ModePerm)
	if err != nil {
		fmt.Println("写入错误", err)
		return
	}

	fmt.Println("创建成功：", filePath)
	return
}

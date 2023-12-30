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
func extractAnnotations(content string) []string {
	lines := strings.Split(content, "\n")
	var annotations []string

	for _, line := range lines {
		if strings.Contains(line, "@") {
			index := strings.Index(line, "@")
			if index != -1 {
				annotations = append(annotations, strings.TrimSpace(line[index+1:]))
			}
		} else if strings.Contains(line, "{") {
			index := strings.Index(line, "{")
			if index != -1 {
				annotations = append(annotations, strings.TrimSpace(line[:index]))
			}
		}
	}
	return annotations
}
func BuildHTML(filePath string, fileName string, annotations []string) {
	fileALLPATH := filePath + ".html"
	var fileContent string
	fileContent += "<!DOCTYPE html>\r\n<html>\r\n<head>\r\n<title>" + fileName + " Documentation</title>\r\n</head>\r\n<body>\r\n"

	fileContent += "<h1>" + fileName + " Documentation</h1>\r\n"
	var filediscription string
	var filediscription_key int
	//找文件描述
	for k, v := range annotations {
		if strings.Contains(v, "brief") {
			filediscription = v
			filediscription_key = k
			break
		}

	}
	fileContent += "<p>" + filediscription + "</p>\r\n"

	fileContent += "<h2>Functions</h2>\r\n"
	//开始函数描述
	type Function struct {
		Name   string
		Brief  string
		Params []string
		Re     string
	}
	var ALLFunction []Function
	thisF := &Function{}
	for i := filediscription_key + 1; i < len(annotations); i++ {
		if strings.Contains(annotations[i], "brief") {
			thisF.Brief = annotations[i]
		} else if strings.Contains(annotations[i], "param") {
			thisF.Params = append(thisF.Params, annotations[i])
		} else if strings.Contains(annotations[i], "return") {
			thisF.Re = annotations[i]
		} else {
			thisF.Name = annotations[i]
			ALLFunction = append(ALLFunction, *thisF)
			thisF = new(Function)
		}
	}
	for _, v := range ALLFunction {
		fileContent += "<h3>" + v.Name + "</h3>\r\n"
		fileContent += "<p>" + v.Brief + "</p>\r\n"
		fileContent += "<ul>\r\n"

		for _, value := range v.Params {
			//第一个空格后表参数值
			index := strings.Index(value, " ")
			//第二个空格后表参数描述
			indexend := strings.Index(value[index+1:], " ")
			fileContent += "<li><strong>" + value[index+1:index+indexend+1] + ":</strong> " + value[index+indexend+1:] + "</li>\r\n"
		}

		fileContent += "<li><strong>Returns:</strong> " + v.Re + "</li>\r\n"
		fileContent += "</ul>\r\n"
	}

	fileContent += "</body>\r\n</html>"

	//写入文件内容
	err := os.WriteFile(fileALLPATH, []byte(fileContent), os.ModePerm)
	if err != nil {
		fmt.Println("写入错误", err)
		return
	}

	fmt.Println("创建成功：", fileName)
}
func (l *BackendLogic) BackendEXPORT(req *types.ExportRequest) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	respon := new(types.Response)
	resp = respon
	//读取指定文件
	// 指定文件目录和文件名
	directory := req.Path
	fileName := req.Name
	// 拼接文件路径
	filePath := filepath.Join(directory, fileName)

	// 读取文件内容
	content, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("读取错误", err)
		return
	}
	/*
		//用正则表达式匹配@+关键词的内容
		flysnowRegexp := regexp.MustCompile(`\/\/[^\n]*@([^\n]+)|\/\*.*?@([^\n]+).*?\*\/`)
		params := flysnowRegexp.FindAllStringSubmatch(string(content), -1)
		for _, v := range params {
			fmt.Println(v[1])
		}
	*/
	//通过字符串读取的方式逐行读取注释
	annotations := extractAnnotations(string(content))
	BuildHTML(filePath, fileName, annotations)
	return
}

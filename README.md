# GoLearning.GO
Go语言傻瓜指南（主要参考Go语言中文网）

环境搭建
1.安装go语言包
	打开下载页面https://studygolang.com/dl，选择"推荐下载->Microsoft Windows"
	下载后直接安装（它会自动设置windows环境变量）
	安装好之后，在cmd命令行中输入"go version"，可以看到版本号
	cd到main.go所在路径，输入"go run main.go"，即可编译并执行

2.安装IDE
	几款GoIDE介绍
		https://studygolang.com/articles/16981?fr=sidebar
		https://www.zhihu.com/question/25012617/answer/688120283(知乎推荐用Goland)	
	我安装过sublime+go+gdb插件，结果装好之后发现gdb不支持sublime3,最后放弃了
	现在安装网上推荐的Go最强IDE Goland, 因为是收费的，需要破解，找了一阵之后，
	最后找到了2020版的安装包，打开goland-2020.1.exe安装就行了
	打开新建工程，右下角会弹黄色警告，Windows Defender might be impacting your build...
	这个时候打开电脑的设置->搜索windows defender->添加排除项->排除文件夹->把你的工程根目录加进去就可以了
	编辑框头部会出现GOPATH is empty，直接点开编辑，或者File->Settings->GO->GOPATH->Project GOPATH->把你的工程根目录加进去就可以了
	开启保存时自动对齐 File->Settings->Tools->File Watchers->"+"->go fmt->ok (go fmt是go的官方代码格式化工具)
	安装之后稍微用一下就会发现好爽^_^

3.开始Go旅行
	按照傻瓜教程一步步学习，碰到教程没有说明的函数查阅标准库文档
		傻瓜教程http://tour.studygolang.com/basics/2
		菜鸟教程https://www.runoob.com/go/go-tutorial.html
		标准库文档https://studygolang.com/pkgdoc
	傻瓜教程的最后给出了一些其他参考、指南、博客，可以进一步学习其中的
		"Go Web编程https://github.com/astaxie/build-web-application-with-golang"
		"Go入门指南https://github.com/Unknwon/the-way-to-go_ZH_CN"
		
	Go 程序的执行（程序启动）顺序如下：
		1.按顺序导入所有被 main 包引用的其它包，然后在每个包中执行如下流程：
		2.如果该包又导入了其它的包，则从第一步开始递归执行，但是每个包只会被导入一次。
		3.然后以相反的顺序在每个包中初始化常量和变量，如果该包含有 init 函数的话，则调用该函数。
		4.在完成这一切之后，main 也执行同样的过程，最后调用 main 函数开始执行程序。

	new和make的区别
		https://www.jb51.net/article/126703.htm
	值类型
		int, float, bool, string, 数组, struct
	引用类型
		指针, slice, map, chan 接口
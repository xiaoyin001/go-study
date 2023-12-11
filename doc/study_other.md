# 其他笔记

## 1、工作区【一个文件夹下多个项目】
    https://segmentfault.com/a/1190000041681242
    https://blog.csdn.net/weixin_49369316/article/details/129612456

    个人总结：
        在跟目录(go-study)下初始化mod，会生成一个go.mod文件，然后分别在目录(go-study/server)和(go-study/client)下，
        执行 go work init 初始化工作区，然后会生成一个go.work文件，这个时候虽然没有报错，
        但是在(go-study/server)和(go-study/client)下的main方法不能执行，会提示没有go.mod文件，
        上面工作区(woek)初始化完毕后，需要将自己用到的go.mod文件引用到这两个go.work文件中，
        就分别在两个路径中执行 go work use ../ ，到这里工作区就初始化完毕了，
        然后就可以在(go-study/server)和(go-study/client)中执行 go run main.go 了

## 2、检查go代码中内存占用
    https://studygolang.com/articles/29812?fr=sidebar

    pprof是Golang提供的一个性能分析工具，它可以对程序的CPU占用和内存占用进行分析
    首先，在代码中导入pprof包 import _ "net/http/pprof"
    然后在程序的main函数中，添加以下代码
    go func() {
        http.ListenAndServe("localhost:8080", nil)
    }()
    在程序启动后，我们就可以通过访问"http://localhost:8080/debug/pprof/"来使用pprof工具
    可以在浏览器中输入"http://localhost:8080/debug/pprof/heap"来查看程序的堆内存占用情况
    可以使用go tool pprof的命令行工具来分析记录的数据。比如，我们可以使用以下命令来查看程序中使用最多内存的函数
    go tool pprof http://localhost:8080/debug/pprof/heap













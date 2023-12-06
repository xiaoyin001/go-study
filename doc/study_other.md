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



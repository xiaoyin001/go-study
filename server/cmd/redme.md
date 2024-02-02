
## Cobra使用方法

### 1、在CMD面板中将路径跳转到指定的的 main.go 文件的路径下面

### 2、在CMD面板中输入将程序遍历出来 [go build]，有依赖需要弄的话先 [go mod tidy] 然后在编译

### 3、最后就是在CMD面板中输入 [exe文件名.exe 先关命令] 就可以运行了
       例如：
            .\exe文件名.exe config
            .\exe文件名.exe server


## Command参数
    这些参数是Go语言中cobra库中的Command结构体的字段，用于定义命令行工具的行为和选项。它们的作用如下：

    Use: 命令名称。
    Aliases: 命令的别名。
    SuggestFor: 命令建议使用的单词列表。
    Short: 命令简短描述。
    GroupID: 命令所属的命令组。
    Long: 命令详细描述。
    Example: 命令的使用示例。
    ValidArgs: 命令接受的参数列表。
    ValidArgsFunction: 命令用于提供动态参数补全的函数。
    Args: 命令的位置参数列表。
    ArgAliases: 位置参数的别名。
    BashCompletionFunction: 生成Bash补全的函数。
    Deprecated: 命令是否已经过时的标志。
    Annotations: 命令的附加注释信息。
    Version: 命令版本号。
    PersistentPreRun: 每次执行该命令之前都会执行的函数。
    PersistentPreRunE: 每次执行该命令之前都会执行的返回错误的函数。
    PreRun: 每次执行该命令之前都会执行的函数。
    PreRunE: 每次执行该命令之前都会执行的返回错误的函数。
    Run: 执行命令的函数。
    RunE: 执行命令的返回错误的函数。
    PostRun: 每次执行该命令之后都会执行的函数。
    PostRunE: 每次执行该命令之后都会执行的返回错误的函数。
    PersistentPostRun: 每次执行该命令之后都会执行的函数。
    PersistentPostRunE: 每次执行该命令之后都会执行的返回错误的函数。
    FParseErrWhitelist : 忽略特定的解析错误
    CompletionOptions :控制 shell 自动完成的选项
    TraverseChildren: 解析父命令的标志后再执行子命令
    Hidden : 隐藏命令，不在可用命令列表中显示
    SilenceErrors : 静默下游错误
    SilenceUsage : 静默错误时不显示用法
    DisableFlagParsing : 禁用标志解析
    DisableAutoGenTag : 禁用自动生成的标记
    DisableFlagsInUseLine : 在打印帮助或生成文档时禁用“[flags]”在用法行中的添加
    DisableSuggestions : 禁用基于Levenshtein距离的建议
    SuggestionsMinimumDistance : 显示建议的最小Levenshtein距离



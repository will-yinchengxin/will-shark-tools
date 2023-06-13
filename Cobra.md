# Cobra
Cobra 是一个库，提供了一个简单的界面来创建类似于 git & go 工具的强大的现代 CLI 界面。

### [命令(Command)](https://pkg.go.dev/github.com/spf13/cobra#Command)
命令是应用程序的核心。应用程序提供的每一个交互都包含在 Command 中。一个命令可以有子命令和可选的运行一个动作。
在 cobra 中，命令和子命令都是用Command结构表示的。

`WC_Linux c all --user`, 其中 `c` 是（子）命令

### 参数（Arg）
命令的参数，即要操作的对象

`WC_Linux c all --user`, 其中 c 是（子）命令，`all` 是参数

### 标志/选项(Flags)
一个标志是一种修饰命令行为的方式。Cobra 支持完全符合 POSIX（可移植操作系统接口） 的标志和 Go flag 包。

`WC_Linux c all --user`, 其中 c 是（子）命令，all 是参数 `--user` 是标志

标志功能由 pflag 库提供，它是标志标准库的一个分支，它在添加 POSIX 合规性的同时保持相同的接口。

### 常用的cobra Command字段及说明
#### Use
命令的名称，用于在终端中调用该命令。
````go
var cmd = &cobra.Command{
    Use:   "hello",
    Short: "Prints hello world",
    Long:  `A longer description that spans multiple lines and likely contains examples and usage of using your command.`,
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("hello world")
    },
}
````
在终端中输入hello即可调用该命令。

**可选参数**
Use: "greet [name]" 表示该命令的名称为 greet，并且可以接受一个参数 name。方括号 [] 表示该参数是可选的，如果不提供参数，则默认为 ""。
在终端中输入 greet Alice，则 greet 为命令名称，Alice 为参数值。在命令的 Run 函数中，可以通过 args 参数获取到该参数的值，例如：
```go
Run: func(cmd *cobra.Command, args []string) {
    if len(args) > 0 {
        fmt.Printf("Hello, %s!\n", args[0])
    } else {
        fmt.Println("Hello, world!")
    }
}
````
如果不提供参数，则输出 Hello, world!。如果提供了参数，则输出 Hello, 参数值!。

**必选参数**
设置必选参数，可以将方括号 [] 改为尖括号 <>，例如：`Use: "greet <name>"`
这样就表示 name 参数是必选的，如果不提供参数，则会提示用户缺少参数。在命令的 Run 函数中，可以通过 args 参数获取到该参数的值，例如：
```go
Run: func(cmd *cobra.Command, args []string) {
    fmt.Printf("Hello, %s!\n", args[0])
}
````
这样就可以保证 args[0] 一定存在，不需要再进行判断。

**可变参数**
使用省略号 ... 表示可变参数，例如：`Use: "sum <numbers...>"`

这样就表示 numbers 参数可以接受任意数量的变量名，可以通过 args 参数获取到一个字符串切片，其中每个元素都是一个数字字符串。
但至少要有一个

**多个选项**
使用竖线 | 表示多个选项，例如：`Use: "copy [--recursive | -r] <src> <dest>"`

这样就表示 copy 命令有一个可选的 --recursive 或 -r 选项，以及两个必选参数 src 和 dest。
[--recursive | -r] 表示一个可选参数，用于指定是否递归处理目录。如果使用了 -r 或 --recursive 参数，则会递归处理目录中的所有子目录和文件。

使用大括号 {} 表示多个选项中的一个，例如：`Use: "search {file|dir} <name>"`

这样就表示 search 命令有一个必选参数 name，以及一个可选的 {file|dir} 选项，表示要搜索的是文件还是目录

{file|dir} 表示一个必选参数，用于指定要处理的文件或目录。{file|dir} 表示这个参数可以是文件或目录，但是必须指定一个。

如果指定的是目录，则根据是否使用了 -r 或 --recursive 参数 (``Use: "search {file|dir}  [--recursive | -r] <name>"``) 来决定是否递归处理目录中的所有子目录和文件。

### Short/Long/Example
- Short：命令的简短描述，用于在帮助文档中显示。
- Long：命令的详细描述，用于在帮助文档中显示。
- Example：命令的使用示例。
```go
var cmd = &cobra.Command{
    Use:   "hello",
    Short: "Prints hello world",
    Long:  `A longer description that spans multiple lines and likely contains examples and usage of using your command.`,
    Example: "hello"
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("hello world")
    },
}
````
### Args
命令的参数列表，用于定义命令需要接受的参数。
```go
var cmd = &cobra.Command{
    Use:   "greet [name]",
    Short: "Greet someone",
    Long: `This command greets someone. If the name is not provided, it will greet the world.`,
    Args: cobra.MaximumNArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        if len(args) > 0 {
            fmt.Printf("Hello, %s!\n", args[0])
        } else {
            fmt.Println("Hello, world!")
        }
    },
}
````
在终端中输入 greet 或 greet Alice 即可调用该命令。

Args 函数可以接受多个参数，用于定义参数的数量、类型、名称等信息。以下是 Args 函数的常用参数：

**MinimumNArgs(n int)：定义命令需要接受的最小参数数量**
````
var cmd = &cobra.Command{
    Use:   "greet [name]",
    Short: "Greet someone",
    Long: `This command greets someone. If the name is not provided, it will greet the world.`,
    Args: cobra.MinimumNArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Printf("Hello, %s!\n", args[0])
    },
}
````
在终端中输入 greet Alice 即可调用该命令，并输出 Hello, Alice!。

**MaximumNArgs(n int)：定义命令需要接受的最大参数数量。**
````
var cmd = &cobra.Command{
    Use:   "greet [name]",
    Short: "Greet someone",
    Long: `This command greets someone. If the name is not provided, it will greet the world.`,
    Args: cobra.MaximumNArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        if len(args) > 0 {
            fmt.Printf("Hello, %s!\n", args[0])
        } else {
            fmt.Println("Hello, world!")
        }
    },
}
````
在终端中输入 greet 或 greet Alice Bob 即可调用该命令，并输出 Hello, Alice! 或 Error: accepts at most 1 arg(s), received 2。

**ExactArgs(n int)：定义命令需要接受的精确参数数量。**
````
var cmd = &cobra.Command{
    Use:   "greet [name]",
    Short: "Greet someone",
    Long: `This command greets someone. If the name is not provided, it will greet the world.`,
    Args: cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Printf("Hello, %s!\n", args[0])
    },
}
````
在终端中输入 greet Alice 即可调用该命令，并输出 Hello, Alice!，但输入 greet 或 greet Alice Bob 则会报错。

**ArbitraryArgs()：定义命令可以接受任意数量的参数。**
````
var cmd = &cobra.Command{
    Use:   "echo [args...]",
    Short: "Echo back the input",
    Long: `This command echoes back the input.`,
    Args: cobra.ArbitraryArgs(),
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println(strings.Join(args, " "))
    },
}
````
在终端中输入 echo hello world 即可调用该命令，并输出 hello world。

**ArgsFunc参数解析函数**
````
var cmd = &cobra.Command{
    Use:   "greet [name]",
    Short: "Greet someone",
    Long: `This command greets someone. If the name is not provided, it will greet the world.`,
    ArgsFunc: func(cmd *cobra.Command, args []string) error {
        if len(args) > 1 {
            return errors.New("accepts at most 1 arg")
        }
        return nil
    },
    Run: func(cmd *cobra.Command, args []string) {
        if len(args) > 0 {
            fmt.Printf("Hello, %s!\n", args[0])
        } else {
            fmt.Println("Hello, world!")
        }
    },
}
````
在终端中输入 greet Alice 即可调用该命令，并输出 Hello, Alice!，但输入 greet Alice Bob 则会报错。

### Run
命令的执行函数，用于定义命令的具体行为。
```go
var cmd = &cobra.Command{
    Use:   "hello",
    Short: "Prints hello world",
    Long:  `A longer description that spans multiple lines and likely contains examples and usage of using your command.`,
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("hello world")
    },
}
````
在终端中输入hello即可调用该命令，并输出hello world。

### Flags
命令的选项列表，用于定义命令需要接受的选项
```go
var cmd = &cobra.Command{
    Use:   "greet",
    Short: "Greet someone",
    Long: `This command greets someone. If the name is not provided, it will greet the world.`,
    Args: cobra.MaximumNArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        language, _ := cmd.Flags().GetString("language")
    
        if language == "spanish" {
            fmt.Println("¡Hola, mundo!")
        } else {
            if len(args) > 0 {
                fmt.Printf("Hello, %s!\n", args[0])
            } else {
                fmt.Println("Hello, world!")
            }
        }
    },
}

cmd.Flags().StringP("language", "l", "english", "language for the greeting")
````
在终端中输入greet --language=spanish即可调用该命令，并输出 `¡Hola, mundo!`。

### PersistentFlags
命令的全局选项列表，用于定义命令及其子命令需要接受的全局选项。
```go
var rootCmd = &cobra.Command{
    Use:   "myapp",
    Short: "My application",
    Long: `My application is a tool that does some things.`,
    Run: func(cmd *cobra.Command, args []string) {
        // Do Stuff Here
    },
}

rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.myapp.yaml)")
````
在终端中输入myapp --config=/path/to/config.yaml即可调用该命令，并指定配置文件路径。

在子命令中，可以使用以下代码来获取该参数的值：
````
Run: func(cmd *cobra.Command, args []string) {
    config, _ := cmd.Parent().PersistentFlags().GetString("config")
    
    if len(config) > 0 {
        fmt.Println("Config Path: ", config)
    } else {
        fmt.Println("Do Not Get Config Path")
    }
},
````

## 1 安装
````
go get -u github.com/spf13/cobra@latest
````
引用
````
import "github.com/spf13/cobra"
````

## 2.使用指南
一般使用该工具的目录结构为
````
▾ appName/
▾ cmd/
    root.go
    command-1.go
    command-2.go
    command-3.go
  main.go
````


## 附录
- code: https://github.com/spf13/cobra
- 使用指南: https://github.com/spf13/cobra/blob/main/user_guide.md
- doc: https://cobra.dev/
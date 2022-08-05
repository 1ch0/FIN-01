## 主要命令
这些都是 cue 命令最常用的命令

cue def
def 命令将统一后的配置打印为同一个文件，在运行代码前查看求值器结果非常有用。

cue eval
eval 命令计算结果、验证并打印配置。

提供很多参数可以展示已求值的内容和其他很多内容。

cue export
根据参数计算并输出配置。

通常用于将 CUE 代码生成 Yaml 和 JSON。

cue vet
根据定义的结构验证数据，使用 -c 参数来确认具体数值，-d 参数来指定要验证的结构。

通过通配符可以执行多个文件。

## import
import 表示将其他外部资源转换为 CUE 代码

cue import 将各种格式数据转为 CUE
cue get go 使用 Go 类型生成 CUE
查看 first-steps/import-configuration 以获取更多信息。

## Modules
cue mod 目前只有 init 命令， 最终它应该会很像 go mod。

查看 modules and packages 和 依赖管理 以获取更多信息。

## Scripting
cue cmd 是如何在工具链中调用命令，你可以运行 cue cmd <anme 或 cue <name>，<name> 不能与官方命令冲突。

看 first-steps/scripting 以获取更多信息。

## 其他命令
cue version 版本信息
cue fmt 标准化和一致性代码
cue completion bash, zsh, fish 和 powershell 的自动补全
cue fix 用于语法修复和重写
cue trim 根据结构定义减少数据冗余

## Help
当然所有的命令都可以通过 cue help 查看帮助，你可以用 help topics 找到一些补充信息。

Additional help topics:
cue commands   user-defined commands
cue filetypes  supported file types and qualifiers
cue flags      common flags for composing packages
cue injection  inject files or values into specific fields for a build
cue inputs     package list, patterns, and files
变量：
变量（Variable）的功能是存储数据，不同的变量保存的数据类型可能会不一样，基本类型有：整型，浮点型，布尔型，字符串。高级类型有：数组，切片，map,指针，结构体，函数,接口，channnel。Go语言中的每一个变量都有自己的类型，并且变量必须经过声明才能开始使用。

变量声明：go语言变量必须声明后才能使用并且声明后必须使用。同一作用域不支持重复使用。var 变量名 变量类型。分为单一声明，批量声明（未说明一样指定类型）

变量初始化：
Go语言在声明变量的时候，会自动对变量对应的内存区域进行初始化操作。每个变量会被初始化成其类型的默认值，例如： 整型和浮点型变量的默认值为0。 字符串变量的默认值为空字符串。 布尔型变量默认为false。 切片、函数、指针变量的默认为nil,map类型的变量默认初始值为nil.但要想使用slice和map这样的引用型变量，必须初始化后才能使用，初始化后才能分配相应的内存，比如用make函数初始化。当然我们也可在声明变量的时候为其指定初始值。变量初始化的标准格式如下：var 变量名 类型 = 表达式。因此产生了类型推导（编译器会根据等号右边的值来推导变量的类型完成初始化。）函数内部：短变量声明。匿名变量：不占用命名空间，不会分配内存，所以匿名变量之间不存在重复声明

注意点：函数外的每个语句都必须以关键字开始（var、const、func等）
:=不能使用在函数外。
_多用于占位，表示忽略值。

常量：常量的声明和变量声明非常类似，只是把var换成了const，常量在定义的时候必须赋值。批量声明：const同时声明多个常量时，如果省略了值则表示和上面一行的值相同。
iota：go语言的常量计数器，在const关键字出现时将被重置为0。const中每新增一行常量声明将使iota计数一次(iota可理解为const语句块中的行索引)。 使用iota能简化定义，在定义枚举时很有用。



运算符：
注意点： ++（自增）和--（自减）在Go语言中是单独的语句，并不是运算符。
算术，关系，逻辑，位，赋值运算符

流程控制：最常用的流程控制有if和for，而switch和goto主要是为了简化代码、降低重复代码而生的结构，属于扩展类的流程控制。
if：两种写法
for:基础和特殊（省略初始但要写；  省略初始和结束，一个；都不要），无限循环。for循环可以通过break、goto、return、panic语句强制退出循环。
for range:数组、切片、字符串:索引和值、map :键和值及通道（channel）:只返回通道内的值
switch:每个switch只能有一个default分支,一个分支可以有多个值，多个case值中间使用英文逗号分隔。分支还可以使用表达式，这时候switch语句后面不需要再跟判断变量,fallthrough语法可以执行满足条件的case的下一个case.每个case后是有编译器默认的break
goto:可以接标签
break:可以接标签
continue:可以接标签






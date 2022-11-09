unicode  它为每种语言中的每个字符设定了统一并且唯一的二进制编码，以满足跨语言、跨平台进行文本转换、处理的要            Unicode就是世界各国合作开发的一种语言。

utf-8   UTF-8以字节为单位对Unicode进行编码。从Unicode到UTF-8的编码方式如下：
Unicode编码(16进制)　║　UTF-8 字节流(二进制)

Go 中的字符串是兼容 Unicode 编码的，并且使用 UTF-8 进行编码。所以字符串是一个字节切片。但放眼国际，有很多语言的一个单词编码超过一个字节，我们就不能用byte类型，要用rune类型，int32的别称，也是go的内建类型。rune代表一个代码点，这个代码点无论占用多少个字节，都可以用一个rune来表示

修改字符串：要修改字符串，需要先将其转换成[]rune或[]byte，完成后再转换为string。无论哪种转换，都会重新分配内存，并复制字节数组。

字符串长度：utf8 package 包中的 func RuneCountInString(s string) (n int) 方法用来获取字符串的长度。这个方法传入一个字符串参数然后返回字符串中的 rune 的数量。

循环处理:
//byte
    func printBytes(s string) {
    for i:= 0; i < len(s); i++ {
        fmt.Printf("%x ", s[i])
    }
}
//rune
func printChars(s string) {
    runes := []rune(s)
    for i:= 0; i < len(runes); i++ {
        fmt.Printf("%c ",runes[i])
    }
}
for index, rune := range s {
        fmt.Printf("%c starts at byte %d\n", rune, index)
    }
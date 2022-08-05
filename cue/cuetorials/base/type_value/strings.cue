// 转义方式
// \a   U+0007 alert or bell
//\b   U+0008 退格
//\f   U+000C 换页
//\n   U+000A 换行
//\r   U+000D 回车
//\t   U+0009 水平制表符
//\v   U+000b 垂直制表符
//\/   U+002f 斜杠(solidus)
//\\   U+005c 反斜杠
//\'   U+0027 单引号 (只在单引号引用的字符串中生效)
//\"   U+0022 双引号 (只在双引号引用的字符串中生效)
//
//\nnn   八进制      (只在单引号引用的字符串中生效)
//\xnn   十六进制    (只在单引号引用的字符串中生效)
//
//\uXXXX  for unicode
//\UXXXXXXXX for longer unicode
str:    "hello world"
smile:  "\U0001F60A"
quoted: "you can \"quote by escaping \\ \""
multiline: """
	hello world
	a "quoted string in a string"
	down under
	   - some author
	"""
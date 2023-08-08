# e-call-go-dll

易语言调用 GoLang 编译的 DLL 示例

![screenshot.png](screenshot.png)

主要问题和解决方案:

- DLL 要编译 32 位的, 请查看 Makefile
- GBK 和 UTF8 编码转换, 解决汉字乱码
- Go 导出声明 __stdcall
- 易语言声明 DLL 命令要加 @ 开头, 解决堆栈错误, 默认是 __cdcel

协程练习
go中多线程使用go运行,加上go关键字就和main主函数一起运行了,为了防止main函数结束
而协程没有结束,可以给main方法设置延时关闭
使用time包   延迟10秒 10 * time.Second


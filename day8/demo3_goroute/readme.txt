before是没有使用goroute,发现运行时间是三秒
使用goroute之后,执行三个goroute程序,运行时间是1秒

不同goroutine之间如何通讯?
1.全局变量和锁(wait)同步
2.channel
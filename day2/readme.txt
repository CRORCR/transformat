iota常量
rand.Intn(100) 随机数
获取主机名,返回字符串和error信息
name, err := os.Hostname()
//获得环境变量的PATH信息
val := os.Getenv("PATH")

//字符串有两种表示形式,"" 这种比较常见,还有一种`` 这种方式换行也会输出,不需要其他连接符

获取时间和时间戳
now := time.Now()
second := now.Unix()
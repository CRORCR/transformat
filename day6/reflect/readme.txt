反射
两个函数
	TypeOf() 获取变量类型
	ValueOf() 获取变量的值
	Value.Kind() 返回一个常量(变量类别)
	Value.Interface() 转换成interface{}类型

	错误演示:
		var i int = 100
		//类型转为valueof结构体
		valueInfo := reflect.ValueOf(i)
		//通过反射改变变量的值
		valueInfo.SetInt(200)
	1.因为valueof是值拷贝,不能改变i的值,必须是指针类型才能改变值
	2.传入指针,修改值就需要加上*,取出指针指向的变量,但是反射没法写*,只能通过Elem()代替
		var i int = 100
		//类型转为valueof结构体
		valueInfo := reflect.ValueOf(&i)
		//通过反射改变变量的值
		valueInfo.Elem().SetInt(200)



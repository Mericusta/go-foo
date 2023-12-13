package todofoo

func todoFoo() {
	// 假设
	// - 存在结构 Basic
	// - 存在结构 Derivative，组合 Basic
	// 是否有一种设计模式，可以使得
	// - Basic 实现控制逻辑，负责业务的开始和结束
	// - Derivative 实现业务逻辑，负责业务的输入和输出

	type basic struct{
		event chan any
		cancel chan any
	}
	
	f := func() {
		for := 
	}
}

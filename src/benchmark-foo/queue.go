package benchmarkfoo

type PoolDequeue interface {
	PushHead(val any) bool
	PopHead() (any, bool)
	PopTail() (any, bool)
}

func NewPoolDequeue(n int) PoolDequeue {
	d := &poolDequeue{
		vals: make([]eface, n),
	}
	d.headTail = d.pack(0, 0)
	return d
}

func (d *poolDequeue) PushHead(val any) bool {
	return d.pushHead(val)
}

func (d *poolDequeue) PopHead() (any, bool) {
	return d.popHead()
}

func (d *poolDequeue) PopTail() (any, bool) {
	return d.popTail()
}

func NewPoolChain() PoolDequeue {
	return new(poolChain)
}

func (c *poolChain) PushHead(val any) bool {
	c.pushHead(val)
	return true
}

func (c *poolChain) PopHead() (any, bool) {
	return c.popHead()
}

func (c *poolChain) PopTail() (any, bool) {
	return c.popTail()
}

// func NewRevertPoolDequeue(n int) PoolDequeue {
// 	d := &revertPoolDequeue{
// 		vals: make([]eface, n),
// 	}
// 	d.headTail = d.pack(1<<dequeueBits-500, 1<<dequeueBits-500)
// 	return d
// }

// func (d *revertPoolDequeue) PushHead(val any) bool {
// 	return d.pushHead(val)
// }

// func (d *revertPoolDequeue) PopHead() (any, bool) {
// 	return d.popHead()
// }

// func (d *revertPoolDequeue) PopTail() (any, bool) {
// 	return d.popTail()
// }

// func NewRPoolChain() PoolDequeue {
// 	return new(revertPoolChain)
// }

// func (c *revertPoolChain) PushHead(val any) bool {
// 	c.pushHead(val)
// 	return true
// }

// func (c *revertPoolChain) PopHead() (any, bool) {
// 	return c.popHead()
// }

// func (c *revertPoolChain) PopTail() (any, bool) {
// 	return c.popTail()
// }

func NewLockFreeQueue() PoolDequeue {
	return NewQueue[int]()
}

func (d *Queue[int]) PushHead(val any) bool {
	d.Enqueue(val.(int))
	return true
}

func (d *Queue[int]) PopHead() (any, bool) {
	return nil, false
}

func (d *Queue[int]) PopTail() (any, bool) {
	return d.Dequeue()
}

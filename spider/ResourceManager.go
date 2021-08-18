package main
//控制线程并发数

type ResourceManager struct {
	tc chan uint8
}

func NewResourceManagerChan(num uint8) *ResourceManager {
	tc := make(chan uint8,num)
	return &ResourceManager{tc:tc}
}

func (r *ResourceManager) GetOne()  {
	r.tc <- 1
}

func (r *ResourceManager) FreeOne()  {
	<-r.tc
}

func (r *ResourceManager) Cap() int {
	return cap(r.tc)
}

func (r *ResourceManager) Has() int {
	return len(r.tc)
}

func (r *ResourceManager) Left() int {
	return cap(r.tc) - len(r.tc)
}

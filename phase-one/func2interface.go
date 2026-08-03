package main

// 函数替代接口只适用于一个接口只有一个方法需要实现的情况
// 如果一个接口里有多个方法需要实现，就不能用函数替代接口了

type Selector interface {
	Select([]string) int
}

type ConnectionPool struct {
	Servers      []string
	LoadBalancer Selector // 成员变量是接口类型
}

type ConnectionPool2 struct {
	Servers      []string
	LoadBalancer func([]string) int // 成员变量是接口类型
}

func f1([]string) int {
	return 0
}

func f2([]string) int {
	return 0
}

type RoundRobin struct{}

func (RoundRobin) Select(s []string) int { return f1(s) }

type Interleave struct{}

func (Interleave) Select(s []string) int { return f2(s) }

func main33() {
	cp := ConnectionPool{
		Servers: []string{"127.0.0.1:1234", "127.0.0.1:5678"},
		// LoadBalancer: RoundRobin{},
		LoadBalancer: Interleave{},
	}
	// _ = cp
	cp.LoadBalancer.Select(cp.Servers)

	cp2 := ConnectionPool2{
		Servers: []string{"127.0.0.1:1234", "127.0.0.1:5678"},
		// LoadBalancer: f1,
		LoadBalancer: f2,
	}
	cp2.LoadBalancer(cp.Servers)
}

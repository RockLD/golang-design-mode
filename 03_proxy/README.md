## 代理模式

*为其他对象提供一种代理以控制对这个对象的访问*

例子：

Subject类 定义了 RealSubject和Proxy的共用接口，这样在任何使用RealSubject的地方都可以使用Proxy。
```go
package proxy
type GiveGift interface{
    GiveDolls()
    GiveFlowers()
    GiveChocolate()
}
```

RealSubject 类,定义Proxy所代表的真实实体
```go
package proxy

import "fmt"

type SchoolGirl struct{
    name string
}

type Pursuit struct {
	mm SchoolGirl
}

func NewPursuit(mm *SchoolGirl) *Pursuit {
	return &Pursuit{mm: *mm}
}

func (p *Pursuit) GiveDolls() {
	fmt.Printf("\n%s 送你洋娃娃", p.mm.name)
}

func (p *Pursuit) GiveFlowers() {
	fmt.Printf("\n%s 送你鲜花", p.mm.name)
}

func (p *Pursuit) GiveChocolate() {
	fmt.Printf("\n%s 送你巧克力", p.mm.name)
}
```

Proxy类，保存一个引用，似的代理可以访问实体，并提供一个与Subject的接口相同的接口，这样代理就可以用来替代实体

```go
package proxy

type Proxy struct {
	gg Pursuit
}

func NewProxy(gg *Pursuit) *Proxy {
	return &Proxy{gg: *gg}
}

func (p *Proxy) GiveDolls() {
	p.gg.GiveDolls()
}

func (p *Proxy) GiveFlowers() {
	p.gg.GiveFlowers()
}

func (p *Proxy) GiveChocolate() {
	p.gg.GiveChocolate()
}

```

## 代理模式应用
- 远程代理
    - 为一个对象在不同的地址空间提供局部代表，这样可以隐藏一个对象存在于不同地址空间的事实
    - 比如：WebService
- 虚拟代理
    - 根据需要创建开销很大的对象，通过他来存放实例化需要很长时间的真实对象
    - 比如：浏览器中的下载
- 安全代理
    - 用来控制真实对象访问时的权限。一般用于对象应该有不同的访问权限的时候。
- 智能指引
    - 指当调用真实的对象时，代理处理另外一些事。
    - 如真实计算对象的引用次数，这样当改对象没有引用时，可以自动释放它；
    -或当第一次引用一个持久对象时，将它装入内存；
    - 或在访问一个实际对象前，检查是否已经锁定它，以确保其他对象不能改变它。
package counter

type Counter struct {
	count int
}

func New() *Counter {
	c := &Counter{
		count: 0,
	}
	return c
}

func (c *Counter) increment() int {
	c.count = c.count + 1
	return c.count
}

func (c *Counter) decrement() int {
	c.count = c.count - 1
	return c.count
}

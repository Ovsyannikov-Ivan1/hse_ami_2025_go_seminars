package tasks

type counter struct {
	value int
}

func newCounter(c int) *counter {
	return &counter{value: c}
}

func (c *counter) Increment() {
	c.value++
}

func (c *counter) Decrement() {
	c.value--
}

func (c *counter) GetValue() int {
	return c.value
}

func (c *counter) Reset() {
	c.value = 0
}

func (c *counter) Add(n int) {
	c.value += n
}

func (c *counter) Subtract(n int) {
	c.value -= n
}

package implementDS

import (
	"fmt"
	"maze/models"
	"sync"
)

type CustomQueue struct {
	Stack []models.NodeStack
	lock  sync.RWMutex
}

func (c *CustomQueue) Push(name models.NodeStack) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.Stack = append(c.Stack, name)
}

func (c *CustomQueue) Pop() error {
	len := len(c.Stack)
	if len > 0 {
		c.lock.Lock()
		defer c.lock.Unlock()
		c.Stack = c.Stack[:len-1]
		return nil
	}
	return fmt.Errorf("Pop Error: Queue is empty")
}

func (c *CustomQueue) Top() models.NodeStack {
	len := len(c.Stack)
	if len > 0 {
		c.lock.Lock()
		defer c.lock.Unlock()
		return c.Stack[len-1]
	}
	return models.NodeStack{0, 0, 0}
}

func (c *CustomQueue) Size() int {
	return len(c.Stack)
}

func (c *CustomQueue) Empty() bool {
	return len(c.Stack) == 0
}

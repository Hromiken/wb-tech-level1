package main

/*18. Реализовать структуру-счетчик, которая будет инкрементироваться в
конкурентной среде. По завершению программа должна выводить итоговое
значение счетчика.
*/
import (
	"fmt"
	"sync"
)

type counter struct {
	mu     sync.Mutex
	number int
}

func main() {
	var wg sync.WaitGroup
	cnt := counter{number: 0}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go cnt.incr(&wg)
	}
	wg.Wait()
	fmt.Println(cnt.number)
}

func (c *counter) incr(wg *sync.WaitGroup) {
	defer wg.Done()
	c.mu.Lock()
	c.number++
	c.mu.Unlock()
}

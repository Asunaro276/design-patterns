package cache

import "fmt"

type Lfu struct{}

func (l *Lfu) evict(c *Cache) {
	fmt.Println("Evicting by lfu strategy")
}

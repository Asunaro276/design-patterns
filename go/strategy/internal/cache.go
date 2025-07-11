package cache

type EvictionAlgo interface {
	evict(c *Cache)
}

type Cache struct {
	storage        map[string]string
	evictionAlgo   EvictionAlgo
	capacity       int
	maxCapacity    int
}

func InitCache(e EvictionAlgo) *Cache {
	storage := make(map[string]string)
	return &Cache{
		storage:        storage,
		evictionAlgo:   e,
		capacity:       0,
		maxCapacity:    2,
	}
}

func (c *Cache) SetEvictionAlgo(e EvictionAlgo) {
	c.evictionAlgo = e
}

func (c *Cache) Add(key, value string) {
	if c.capacity == c.maxCapacity {
		c.Evict()
	}
	c.capacity++
	c.storage[key] = value
}

func (c *Cache) Get(key string) string {
	return c.storage[key]
}

func (c *Cache) Evict() {
	c.evictionAlgo.evict(c)
	c.capacity--
}

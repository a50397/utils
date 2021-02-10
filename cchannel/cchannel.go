package cchannel

// Cchannel fixed sized non-blocking buffer implemented using channels, Only last n values will be kept
type Cchannel struct {
	c        chan interface{}
	Capacity int
	Free     int
}

// NewChan create new
func (c *Cchannel) NewChan(size int) {
	c.c = make(chan interface{}, size)
	c.Capacity = size
	c.Free = size
}

// Add add item to the cchannel
func (c *Cchannel) Add(item interface{}) {
	select {
	case c.c <- item:
		c.Free--
	default:
		_ = <-c.c
		c.c <- item
	}
}

// Get non-blocking get
func (c *Cchannel) Get() (interface{}, bool) {
	var ret interface{}
	if c.Capacity == c.Free {
		return ret, false
	}
	select {
	case ret = <-c.c:
		c.Free++
		return ret, true
	default:
		return ret, false
	}
}

// GetB blocking get
func (c *Cchannel) GetB() interface{} {
	return <-c.c
}

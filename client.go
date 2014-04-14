package triblab
//package main

import(
	"trib"
	"net/rpc"
)

type client struct {
	// private fields
	addr string
}

// implents Key-value pair interfaces
func (self *client) Get(key string, value *string) error{
	conn, e := rpc.DialHTTP("tcp", self.addr)
	if e != nil {
		return e
	}

	e = conn.Call("Storage.Get", key, value)
	if e != nil {
		conn.Close()
		return e
	}

	return conn.Close()
}

func (self *client) Set(kv *trib.KeyValue, succ *bool) error{
	conn, e := rpc.DialHTTP("tcp", self.addr)
	if e != nil {
		return e
	}

	e = conn.Call("Storage.Set", kv, succ)
	if e != nil {
		conn.Close()
		return e
	}

	return conn.Close()
}

func (self *client) Keys(p *trib.Pattern, list *trib.List) error{
	panic("todo")
}

// implements Key-list interfaces.
func (self *client) ListGet(key string, list *trib.List) error{
	panic ("todo")
}

func (self *client) ListAppend(kv *trib.KeyValue, succ *bool) error{
	panic("todo")
}

func (self *client) ListRemove(kv *trib.KeyValue, n *int) error{
	panic("todo")
}

func (self *client) ListKeys(p *trib.Pattern, list *trib.List) error{
	panic("todo")
}

//implements clock interface
func (self *client) Clock(atLeast uint64, ret *uint64) error{
	panic("todo")
}

/*
func main(){

	var _ trib.Storage = new (client)
}
*/


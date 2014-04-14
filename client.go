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
	conn, e := rpc.DialHTTP("tcp", self.addr)
	if e != nil {
		return e
	}

	e = conn.Call("Storage.Keys", p, list)
	if e != nil {
		conn.Close()
		return e
	}

	return conn.Close()
}

// implements Key-list interfaces.
func (self *client) ListGet(key string, list *trib.List) error{
	conn, e := rpc.DialHTTP("tcp", self.addr)
	if e != nil {
		return e
	}

	e = conn.Call("Storage.ListGet", key, list)
	if e != nil {
		conn.Close()
		return e
	}

	return conn.Close()
}

func (self *client) ListAppend(kv *trib.KeyValue, succ *bool) error{
	conn, e := rpc.DialHTTP("tcp", self.addr)
	if e != nil {
		return e
	}

	e = conn.Call("Storage.ListAppend", kv, succ)
	if e != nil {
		conn.Close()
		return e
	}

	return conn.Close()
}

func (self *client) ListRemove(kv *trib.KeyValue, n *int) error{
	conn, e := rpc.DialHTTP("tcp", self.addr)
	if e != nil {
		return e
	}

	e = conn.Call("Storage.ListRemove", kv, n)
	if e != nil {
		conn.Close()
		return e
	}

	return conn.Close()
}

func (self *client) ListKeys(p *trib.Pattern, list *trib.List) error{
	conn, e := rpc.DialHTTP("tcp", self.addr)
	if e != nil {
		return e
	}

	e = conn.Call("Storage.ListKeys", p, list)
	if e != nil {
		conn.Close()
		return e
	}

	return conn.Close()
}

//implements clock interface
func (self *client) Clock(atLeast uint64, ret *uint64) error{
	conn, e := rpc.DialHTTP("tcp", self.addr)
	if e != nil {
		return e
	}

	e = conn.Call("Storage.Clock", atLeast, ret)
	if e != nil {
		conn.Close()
		return e
	}

	return conn.Close()
}

/*
func main(){

	var _ trib.Storage = new (client)
}
*/


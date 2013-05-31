package attr

type String string

func (self *String) Set(v string) {
	*self = String(v)
}

package attr

type String string

func (self *String) Set(v string) {
	*self = String(v)
}

func (self *String) Get() string {
	return string(*self)
}

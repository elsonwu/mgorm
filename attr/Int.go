package attr

type Int int

func (self *Int) Get() int {
	return int(*self)
}

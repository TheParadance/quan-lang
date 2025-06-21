package object

type Value interface{}
type Object struct {
	Properties map[string]Value
}

func (o *Object) SetProperty(name string, val Value) {
	o.Properties[name] = val
}

func (o *Object) GetProperty(name string) (Value, bool) {
	val, ok := o.Properties[name]
	return val, ok
}

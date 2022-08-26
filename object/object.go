package object

// objects in Go are structs / key-value pairs
type Object struct {
	Field1 string // if the field name starts with a capital letter it is "public"
	field2 string // if the field name starts with a lower case letter it is "private"
}

// in Go we can specify instance methods for the object
func (o *Object) SetField2(field2 string) {
	o.field2 = field2
}

func (o *Object) GetField2() string {
	return o.field2
}

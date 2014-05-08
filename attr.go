package grapho

// Attr is a set of attributes associated to a node/edge.
// Keys are strings. Values can be anything.
// Set associates a new key-value pair.
// Get returns the associated value to the given key, and
// a bool set to true if the key was found, false otherwise
type Attr interface {
	Set(key string, value interface{})
	Get(key string) (interface{}, bool)
}

// NewAttr creates an empty set of attributes.
func NewAttr() Attr {
	return newAttrMap()
}

// attrMap implements Attr, wrapping a map holding the data
type attrMap struct {
	attr map[string]interface{}
}

func newAttrMap() *attrMap {
	attr := attrMap{make(map[string]interface{})}
	return &attr
}

func (attr *attrMap) Set(key string, value interface{}) {
	attr.attr[key] = value
}

func (attr *attrMap) Get(key string) (interface{}, bool) {
	v, ok := attr.attr[key]
	return v, ok
}

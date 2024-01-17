
## First way of implementing constructor  
```go
type Polygon interface {
	Area() int
	Perimeter() int
}

type Triangle struct {
	Edge1 int
	Edge2 int
	Edge3 int
}

func (t *Triangle) Area() int {
	return t.Edge1 * t.Edge2
}
func (t *Triangle) Perimeter() int {
	return t.Edge1 * t.Edge2
}
```

## Second way of implementing constructor  
### struct is private
```go
type Polygon interface {
	Area() int
	Perimeter() int
}

type triangle struct {
	Edge1 int
	Edge2 int
	Edge3 int
}

func NewTriangle(e1,e2,e3 int) Polygon {
	return &triangle{Edge1: e1, Edge2: e2, Edge3: e3}
}

func (t *triangle) Area() int {
	return t.Edge1 * t.Edge2
}
func (t *triangle) Perimeter() int {
	return t.Edge1 * t.Edge2
}
```


## named return
```go
// returns <nil>
// when the panic() is called, it stops execution, so line return err does not run
// next defer func is called, which assigns "ssss" to the "var err error"
// once defer is complete, it returns via the named returned defined in the function name "(error)"
// in this case it is just error with no name, which was initialised to nil, hence <nil> is returned
func Sth() error {
	var err error
	defer func() {
		if p := recover(); p != nil {
			err = errors.New("ssss")
		}
	}()
	panic("STH")
	return err
}

// returns "ssss"
// notice the named return "err error", this is initialised to nil
// when panic() is called, stops execution
// defer function is called, in this case it assigns "ssss" to the named return (err error)
// once defer ends, it returns the named variable (err error)
func Sthh() (err error) {
	defer func() {
		if p := recover(); p != nil {
			err = errors.New("ssss")
		}
	}()
	panic("STH")
	return
}
```
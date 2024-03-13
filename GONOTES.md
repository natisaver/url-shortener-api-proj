
# Interfaces
- What is its purpose? Code only needs to interact with the methods of the interface
- It does not need to know the implementation of the methods
- You can thus change the implementation of the methods without affecting the code

## First way of implementing interface methods  
- NOTE: struct is **public** 
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

## Second way of implementing interface methods
## using a constructor  
- NOTE: struct is **private**
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
if you declare a variable name for the return part of your function
then this variable can be used in the function body


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

# Git


### Working on bugs/features
#### Create a branch first
thereafter once changes are complete, commit to the branch. You can then make a pull request back to master for review to merge those changes


### Merging feature branch with latest changes from master
```bash
# go to the master branch
git pull
# go to the feature branch
git rebase master featurebranchname

```

### Format for git branch names
We are adapting to git-flow way of doing things.
```
Branch
  1. master                                      // primary source of truth, QA/Production readiness depend on the CICD pipelines
  2. develop/<feature>                           // secondary source of truth, typically more applicable when handling large scale project where feature per team level
  3. feature/<type-of-feature>-<name-of-feature> // all the features include, this branch will either merge to master branch or develop branch on the scale of the project
  4. hotfix/vX.X.X.X                             // this is usually a fix require to be done on the production, this branch usually merge directly to master
  5. release/vX.X.X.X                            // this is usually snapshot of all the production readiness, this allow revert of release should new release serious failure in production, typically this are usually last resort scenarios.

type-of-features
  1. document
  2. component
  3. page
  4. api          // this include hooks, state management, endpoints calls
  5. test
  6. prototype-vX // usually this compose of various feature branch for showcase of product to client
  7. fix          // for the feature but not use for staging / production yet

Notes: Depend on how the CICD works, usually latest commit of features or develop branch will be push to DEV environment. Feel free to discuss with the leads, should there be any possible improvements.
```
## References:

1. Git Flow - http://danielkummer.github.io/git-flow-cheatsheet/



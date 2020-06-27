package rand

// Dog is a dog!
type Dog struct {
	color  string
	height int
	name   string
}

// NewDog is a Dog constructor
func NewDog(color string, height int, name string) Dog {
	return Dog{color: color, height: height, name: name}
}

// Puppy is a Puppy!!
type Puppy struct {
	parent Dog
	name   string
}

// NewPuppy is a Puppy constructor
func NewPuppy(color string, height int, name string) Puppy {
	return Puppy{
		parent: Dog{color: color, height: height, name: name},
		name:   name,
	}
}

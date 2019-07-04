// A simple File interface
interface {
	Read(b Buffer) bool
	Write(b Buffer) bool
	Close()
}

// An empty struct.
struct {}

// A struct with 6 fields.
struct {
	x, y int
	u float32
	_ float32  // padding
	A *[]int
	F func()
}
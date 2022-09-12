package main

// interfaceの設計指針としてメソッドリストは極力少なくする。
// Bad
type DuckTester interface {
	Cry() string
	Footsteps() string
}

// Good
type Crier interface {
	Cry() string
}
type Footstepper interface {
	Footsteps() string
}

// インターフェース自体に別のインターフェースを埋め込むことで合成ができる
type CryFootstepper interface {
	Crier
	Footstepper
}

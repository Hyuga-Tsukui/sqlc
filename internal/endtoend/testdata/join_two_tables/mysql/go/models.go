// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0

package querytest

import ()

type Bar struct {
	ID uint64
}

type Baz struct {
	ID uint64
}

type Foo struct {
	BarID uint64
	BazID uint64
}

//go:build cgo

package hello

func init() {
	panic("linked with cgo")
}

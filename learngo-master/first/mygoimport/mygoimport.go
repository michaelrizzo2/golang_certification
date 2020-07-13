package mygoimport

import "runtime"

func Myversion() string {
	return runtime.Version()
}

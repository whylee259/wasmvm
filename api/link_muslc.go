// +build linux,muslc

package api

// #cgo LDFLAGS: -L${SRCDIR} -lwasmvm_muslc
import "C"

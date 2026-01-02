package main

// #include "autoexports.h"
import "C"
import (
	"reflect"
	"unsafe"

	"github.com/untrustedmodders/go-plugify"
)

var _ = reflect.TypeOf(0)
var _ = unsafe.Sizeof(0)
var _ = plugify.Plugin.Loaded

// Exported methods

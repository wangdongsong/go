package io

import (
	"jvmgo/rtda"
	"jvmgo/rtda/heap"
)

func init() {
}

func _fd(method func(frame *rtda.Frame), name, desc string) {
	heap.RegisterNativeMethod("java/io/FileDescriptor", name, desc, method)
}

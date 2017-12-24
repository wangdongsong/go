package io

import (
	"jvmgo/rtda"
	"jvmgo/rtda/heap"
)

func init() {
	_osc(initNative, "initNative", "()V")
}

func _osc(method func(frame *rtda.Frame), name, desc string) {
	heap.RegisterNativeMethod("java/io/ObjectStreamClass", name, desc, method)
}

// private static native void initNative();
// ()V
func initNative(frame *rtda.Frame) {
	// todo
}

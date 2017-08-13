package native

import (
	"jvmgo/rtda"
	"jvmgo/rtda/heap"
	_ "jvmgo/native/java/awt"
	_ "jvmgo/native/java/io"
	_ "jvmgo/native/java/lang"
	_ "jvmgo/native/java/lang/invoke"
	_ "jvmgo/native/java/lang/reflect"
	_ "jvmgo/native/java/net"
	_ "jvmgo/native/java/security"
	_ "jvmgo/native/java/util"
	_ "jvmgo/native/java/util/concurrent/atomic"
	_ "jvmgo/native/java/util/jar"
	_ "jvmgo/native/java/util/zip"
	_ "jvmgo/native/sun/awt"
	_ "jvmgo/native/sun/java2d/opengl"
	_ "jvmgo/native/sun/management"
	_ "jvmgo/native/sun/misc"
	_ "jvmgo/native/sun/nio/ch"
	_ "jvmgo/native/sun/reflect"
)

func init() {
	heap.SetEmptyNativeMethod(emptyNativeMethod)
}

func emptyNativeMethod(frame *rtda.Frame) {
	// do nothing
}

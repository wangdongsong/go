package awt

import (
	"jvmgo/rtda"
	"jvmgo/rtda/heap"
)

func init() {
}

func _container(method func(frame *rtda.Frame), name, desc string) {
	heap.RegisterNativeMethod("java/awt/Container", name, desc, method)
}

package awt

import (
	"jvmgo/rtda"
	"jvmgo/rtda/heap"
)

func init() {
}

func _comp(method func(frame *rtda.Frame), name, desc string) {
	heap.RegisterNativeMethod("java/awt/Component", name, desc, method)
}

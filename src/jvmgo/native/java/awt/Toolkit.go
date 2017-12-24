package awt

import (
	"jvmgo/rtda"
	"jvmgo/rtda/heap"
)

func init() {
}

func _tk(method func(frame *rtda.Frame), name, desc string) {
	heap.RegisterNativeMethod("java/awt/Toolkit", name, desc, method)
}

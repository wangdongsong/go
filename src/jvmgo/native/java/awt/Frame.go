package awt

import (
	"jvmgo/rtda"
	"jvmgo/rtda/heap"
)

func _frame(method func(frame *rtda.Frame), name, desc string) {
	heap.RegisterNativeMethod("java/awt/Frame", name, desc, method)
}

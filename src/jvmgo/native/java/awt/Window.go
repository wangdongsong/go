package awt

import (
	"jvmgo/rtda"
	"jvmgo/rtda/heap"
)

func _window(method func(frame *rtda.Frame), name, desc string) {
	heap.RegisterNativeMethod("java/awt/Window", name, desc, method)
}

package awt

import (
	"jvmgo/rtda"
	"jvmgo/rtda/heap"
)

func init() {
}

func _cursor(method func(frame *rtda.Frame), name, desc string) {
	heap.RegisterNativeMethod("java/awt/Cursor", name, desc, method)
}

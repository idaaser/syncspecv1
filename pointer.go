package syncspec

// Pointer 转换T类型的值v为指针类型*T
func Pointer[T any](v T) *T {
	return &v
}

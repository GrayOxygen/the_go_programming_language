package chapters

import "testing"

//测试strings.Join()和自己手动+拼接的性能
func BenchmarkEcho3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Echo3()
	}
}

func BenchmarkEcho(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Echo()
	}
}

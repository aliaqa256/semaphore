package queue

import (
	"fmt"
	"testing"
)

// benchmarking
func BenchmarkQueue(b *testing.B) {
	seq := []int{1, 2, 3}
	q := Queue{}
	for _,seq := range seq {
		b.Run(fmt.Sprintf("Queue-%d", seq), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				q.Enqueue(seq)
			}
			
		})

	}

}

// TestQueue tests the queue
func BenchmarkDequeue(b *testing.B) {
	q := Queue{}
	for i := 0; i < 71396013; i++ {
		q.Enqueue(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		q.Dequeue()
	}
}
// go  test  ./queue  -run=xxxx -v  -bench=.  benchtime3s -benchmem 




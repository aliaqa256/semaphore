package queue

import (
	"fmt"
	"testing"
)

// benchmarking
func BenchmarkQueue(b *testing.B) {
	seq := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
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
	for i := 0; i < b.N; i++ {
		q.Enqueue(i)
	}
	for i := 0; i < b.N; i++ {
		q.Dequeue()
	}
}





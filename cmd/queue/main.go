package main

import (
	"fmt"
	"math/rand/v2"
	"time"

	"github.com/dev-dmitrii-g/Algorithms-Data-Structures-Tomorrow-School-Club/internal/queue"
)

const SIZE int = 10

type EmailJob struct {
	From string
	To   string
}

func main() {
	q := queue.New[EmailJob]()

	for i := SIZE; i >= 0; i-- {
		from := fmt.Sprintf("user%v@email.com", rand.IntN(SIZE))
		to := fmt.Sprintf("user%v@email.com", rand.IntN(SIZE))
		email := EmailJob{From: from, To: to}
		q.Enqueue(email)
	}

	for email := range q.Drain() {
		fmt.Println(q.String())
		fmt.Printf("from : %v\t to: %v\nQueue length: %v\n", email.From, email.To, q.Len())
		time.Sleep(time.Duration(rand.IntN(100)) * time.Millisecond * 50)
	}
	fmt.Println("Job finished!")
}

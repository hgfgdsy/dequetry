package deque

import (
	"log"

	"github.com/joncrlsn/dque"
)

// Item is what we'll be storing in the queue.  It can be any struct
// as long as the fields you want stored are public.
type Item struct {
	Name string
	Id   int
}

// ItemBuilder creates a new item and returns a pointer to it.
// This is used when we load a segment of the queue from disk.
func ItemBuilder() interface{} {
	return &Item{}
}

func main() {
	ExampleDQue_main()
}

// ExampleQueue_main() show how the queue works
func ExampleDQue_main() {
	qName := "item-queue"
	qDir := "/tmp"
	segmentSize := 50

	// Create a new queue with segment size of 50
	q, err := dque.New(qName, qDir, segmentSize, ItemBuilder)
	...

	// Add an item to the queue
	err := q.Enqueue(&Item{"Joe", 1})
	...


	// You can reconsitute the queue from disk at any time
	// as long as you never use the old instance
	q, err = dque.Open(qName, qDir, segmentSize, ItemBuilder)
	...

	// Peek at the next item in the queue
	var iface interface{}
	if iface, err = q.Peek(); err != nil {
		if err != dque.ErrEmpty {
			log.Fatal("Error peeking at item ", err)
		}
	}

	// Dequeue the next item in the queue
	if iface, err = q.Dequeue(); err != nil {
		if err != dque.ErrEmpty {
			log.Fatal("Error dequeuing item ", err)
		}
	}

	// Assert type of the response to an Item pointer so we can work with it
	item, ok := iface.(*Item)
	if !ok {
		log.Fatal("Dequeued object is not an Item pointer")
	}

	doSomething(item)
}

func doSomething(item *Item) {
	log.Println("Dequeued", item)
}

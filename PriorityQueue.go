package mygoProject

import (
	"container/heap"
	"errors"
	"github.com/mitchellh/copystructure"
	"sync"
)

type Item struct {
	Key string
	// Value is an unspecified type that implementations can use to store
	// information
	Value interface{}
	// Priority determines ordering in the queue, with the lowest value being the
	// highest priority
	Priority int64
	// index is an internal value used by the heap package, and should not be
	// modified by any consumer of the priority queue
	index int
}
type ItemList []Item

func (l *ItemList) Pop() *Item {
	old := *l
	value := old[len(*l)-1]
	value = (*l)[len(*l)-1]
	*l = old[:len(*l)-1]
	return &value
}

var ErrEmpty = errors.New("queue is empty")
var ErrDuplicateItem = errors.New("depulicate item")

func NewPriorityQueue() *PriorityQueue {
	pq := PriorityQueue{
		data:    make(queue, 0),
		dataMap: make(map[string]*Item),
	}
	heap.Init(&pq.data)
	return &pq
}

type PriorityQueue struct {
	data    queue
	dataMap map[string]*Item
	lock    sync.RWMutex
}
type queue []*Item

func (q *queue) Len() int {
	//TODO implement me
	panic("implement me")
}

func (q *queue) Less(i, j int) bool {
	//TODO implement me
	panic("implement me")
}

func (q *queue) Swap(i, j int) {
	//TODO implement me
	panic("implement me")
}

func (q *queue) Push(x any) {
	len := len(*q)
	item := x.(*Item)
	item.index = len
	*q = append(*q, item)
}

func (q *queue) Pop() any {
	len := len(*q)
	item := (*q)[len-1]
	*q = (*q)[:len-1]
	return item
}

func (pq *PriorityQueue) Len() int {
	pq.lock.RLock()
	defer pq.lock.RUnlock()
	return len(pq.data)

}
func (pq *PriorityQueue) Pop() (*Item, error) {
	pq.lock.Lock()
	defer pq.lock.Unlock()
	if pq.Len() == 0 {
		return nil, ErrEmpty
	}
	item := heap.Pop(&pq.data).(*Item)
	delete(pq.dataMap, item.Key)
	return item, nil
}
func (pq *PriorityQueue) Push(i *Item) error {
	if i == nil || i.Key == "" {
		return errors.New("error adding item")
	}
	pq.lock.Lock()
	defer pq.lock.Unlock()
	if _, ok := pq.dataMap[i.Key]; ok {
		return ErrDuplicateItem
	}
	clone, err := copystructure.Copy(i)
	if err != nil {
		return err
	}
	pq.dataMap[i.Key] = clone.(*Item)
	heap.Push(&pq.data, clone)
	return nil
}
func (q *queue) push(x interface{}) {
	n := len(*q)
	item := x.(*Item)
	item.index = n
	*q = append(*q, item)
}

// 任务对象
type task struct {
	priority int64 //优先级
	value    interface{}
	key      string
}
type PriorityQueueTask struct {
	mLock    sync.Mutex
	pushChan chan *task
	pq       *PriorityQueue
}

func NewPriorityQueueTask() *PriorityQueueTask {
	//使用&，使用了结构体的指针来初始化 pq。这意味着 pq 是一个指向 PriorityQueueTask 类型的指针
	//如果用pq := PriorityQueueTask{...} 直接使用值来初始化 pq。
	//这种方式创建了 PriorityQueueTask 类型的一个实例，并且 pq 是这个实例的值拷贝
	pq := &PriorityQueueTask{
		pushChan: make(chan *task, 1000),
		pq:       NewPriorityQueue(),
	}
	go pq.ListenPushChan()
	return pq
}
func (pq *PriorityQueueTask) ListenPushChan() {
	for {
		select {
		case task := <-pq.pushChan:
			pq.mLock.Lock()
			pq.pq.Push(&Item{Key: task.key, Value: task.value, Priority: task.priority})
			pq.mLock.Unlock()
		}
	}
}

func (pq *PriorityQueueTask) Push(priority int64, value interface{}, key string) {
	pq.pushChan <- &task{
		key:      key,
		value:    value,
		priority: priority,
	}
}

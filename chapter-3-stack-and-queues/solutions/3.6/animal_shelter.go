package animal_shelter

import (
	"container/list"
	"strings"
)

/*
	3.6: Animal Shelter: An animal shelter, which holds only dogs and cats, operates on a strictly "first in ,
	first out" basis. People must adopt either the "oldest" (based on arrival time) of all animals at the shelter,
	or they can select whether they would prefer a dog or a cat (and will receive the oldest animal of that type).
	They cannot select which specific animal they would like. Create the data structure to maintain this system
	and implement operations such as enqueue, dequeueAny, dequeueDog, and dequeueCAt. You may used a linked-list
	data structure.
*/

/*
We could explore a variety of solutions to this problem. For instance, we could maintain a single queue.
This would make dequeueAny easy, but dequeueDog and dequeueCAt would require iteration through the queue
to find the first dog or cat. This would increase the complexity of the solution and decrease the efficiency.

An alternative approach is simple, clean, and efficient is to simply use separate queues for dogs and cats, and to
place them within a wrapper class called AnimalQueue. We then store some sort of timestamp to mark when each animal
was enqueued. When we call dequeueAny, we peek at the heads of both the dog and cat and return the oldest.
*/

type Animal struct {
	Order int
	Name  string
	Type  string
}

func (a *Animal) IsOlderThan(b *Animal) bool {
	return a.Order < b.Order
}

type AnimalQueue struct {
	Dogs  *list.List
	Cats  *list.List
	Order int
}

func CreateAnimalQueue() *AnimalQueue {
	return &AnimalQueue{
		Dogs:  list.New().Init(),
		Cats:  list.New().Init(),
		Order: 0,
	}
}

func (q *AnimalQueue) Enqueue(a Animal) {
	// order is used as a sort of timestamp, so that we can compare the insertion
	// order of a dog to a cat
	a.Order = q.Order
	q.Order++
	switch strings.ToLower(a.Type) {
	case "dog":
		q.Dogs.PushBack(a)
	case "cat":
		q.Cats.PushBack(a)
	}
}

func (q *AnimalQueue) DequeueAny() *Animal {
	// look at tops of dog and cat queues, and pop the queue with the oldest value
	if q.Dogs.Len() == 0 {
		return q.DequeueCats()
	} else if q.Cats.Len() == 0 {
		return q.DequeueDogs()
	}
	dogF := q.Dogs.Front()
	dog, ok := dogF.Value.(*Animal)
	if !ok {
		return nil
	}
	catF := q.Cats.Front()
	cat, ok := catF.Value.(*Animal)
	if !ok {
		return nil
	}
	if dog.IsOlderThan(cat) {
		return q.DequeueDogs()
	} else {
		return q.DequeueCats()
	}
}

func (q *AnimalQueue) DequeueDogs() *Animal {
	v := q.Dogs.Front()
	q.Dogs.Remove(v)
	a, ok := v.Value.(Animal)
	if !ok {
		return nil
	}
	return &a
}

func (q *AnimalQueue) DequeueCats() *Animal {
	v := q.Cats.Front()
	q.Cats.Remove(v)
	a, ok := v.Value.(Animal)
	if !ok {
		return nil
	}
	return &a
}

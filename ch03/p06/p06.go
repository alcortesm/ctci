package p06

import (
	"errors"
	"fmt"

	"github.com/alcortesm/ctci/ch03/p06/internal"
)

var ErrNoAnimals = errors.New("no animals")

type Animal struct {
	Name  string
	IsDog bool // if it is not a dog, it is a cat.
}

func (a Animal) String() string {
	return a.Name
}

type Shelter struct {
	seen       int // number of animals seen so far
	dogs, cats *internal.Queue
}

func NewShelter() *Shelter {
	return &Shelter{
		dogs: &internal.Queue{},
		cats: &internal.Queue{},
	}
}

func (s *Shelter) Enqueue(a *Animal) {
	q := s.cats
	if a.IsDog {
		q = s.dogs
	}

	q.Enqueue(&internal.Intake{
		Item:           a,
		AdmissionOrder: s.seen,
	})

	s.seen++
}

func (s *Shelter) DequeueDog() (*Animal, error) {
	return dequeueAnimalFrom(s.dogs)
}

func (s *Shelter) DequeueCat() (*Animal, error) {
	return dequeueAnimalFrom(s.cats)
}

// returns ErrNoAnimals if there are no animals in the queue
func dequeueAnimalFrom(q *internal.Queue) (*Animal, error) {
	intake, err := q.Dequeue()
	if err != nil {
		if errors.Is(err, internal.ErrEmpty) {
			return nil, ErrNoAnimals
		}
		return nil, err
	}

	return intakeToAnimal(intake)
}

func intakeToAnimal(i *internal.Intake) (*Animal, error) {
	if i == nil {
		return nil, fmt.Errorf("found nil intake")
	}

	animal, ok := i.Item.(*Animal)
	if !ok {
		return nil, fmt.Errorf("not an animal found in intake")
	}

	return animal, nil
}

func (s *Shelter) DequeueAny() (*Animal, error) {
	noDogs := false
	dogIntake, err := s.dogs.Peek()
	if err != nil {
		if errors.Is(err, internal.ErrEmpty) {
			noDogs = true
		} else {
			return nil, fmt.Errorf("peeking at dogs queue: %v", err)
		}
	}

	noCats := false
	catIntake, err := s.cats.Peek()
	if err != nil {
		if errors.Is(err, internal.ErrEmpty) {
			noCats = true
		} else {
			return nil, fmt.Errorf("peeking at cats queue: %v", err)
		}
	}

	switch {
	case noDogs && noCats:
		return nil, ErrNoAnimals
	case noDogs:
		return dequeueAnimalFrom(s.cats)
	case noCats:
		return dequeueAnimalFrom(s.dogs)
	default:
		containsOldest := s.dogs
		if catIntake.IsOldest(dogIntake) {
			containsOldest = s.cats
		}

		return dequeueAnimalFrom(containsOldest)
	}
}

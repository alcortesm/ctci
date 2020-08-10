package p06_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/alcortesm/ctci/ch03/p06"
	"github.com/google/go-cmp/cmp"
)

func TestShelter_DequeueAny(t *testing.T) {
	t.Parallel()

	d1 := &p06.Animal{Name: "dog1", IsDog: true}
	d2 := &p06.Animal{Name: "dog2", IsDog: true}

	c1 := &p06.Animal{Name: "cat1"}
	c2 := &p06.Animal{Name: "cat2"}

	subtests := [][]*p06.Animal{
		{},
		{d1},
		{c1},
		{d1, d2},
		{c1, c2},
		{d1, d2, c1, c2},
		{c1, c2, d1, d2},
		{d1, c1, d2, c2},
		{c1, d1, c2, d2},
	}

	for _, animals := range subtests {
		animals := animals
		name := fmt.Sprintf("%s", animals)

		t.Run(name, func(t *testing.T) {
			shelter := p06.NewShelter()

			for _, a := range animals {
				shelter.Enqueue(a)
			}

			got := []*p06.Animal{}
			for {
				a, err := shelter.DequeueAny()
				if err != nil {
					if errors.Is(err, p06.ErrNoAnimals) {
						break
					}
					t.Fatal(err)
				}

				got = append(got, a)
			}

			if diff := cmp.Diff(animals, got); diff != "" {
				t.Errorf("(-want +got)\n%s", diff)
			}
		})
	}
}

func TestShelter_DequeueDogs(t *testing.T) {
	t.Parallel()

	d1 := &p06.Animal{Name: "dog1", IsDog: true}
	d2 := &p06.Animal{Name: "dog2", IsDog: true}
	d3 := &p06.Animal{Name: "dog3", IsDog: true}

	c1 := &p06.Animal{Name: "cat1"}
	c2 := &p06.Animal{Name: "cat2"}
	c3 := &p06.Animal{Name: "cat3"}

	subtests := []struct {
		intakes []*p06.Animal
		want    []*p06.Animal
	}{
		{
			intakes: []*p06.Animal{},
			want:    []*p06.Animal{},
		},
		{
			intakes: []*p06.Animal{c1},
			want:    []*p06.Animal{},
		},
		{
			intakes: []*p06.Animal{c1, c2},
			want:    []*p06.Animal{},
		},
		{
			intakes: []*p06.Animal{d1},
			want:    []*p06.Animal{d1},
		},
		{
			intakes: []*p06.Animal{d1, d2},
			want:    []*p06.Animal{d1, d2},
		},
		{
			intakes: []*p06.Animal{c1, d1, d2},
			want:    []*p06.Animal{d1, d2},
		},
		{
			intakes: []*p06.Animal{d1, c1, d2},
			want:    []*p06.Animal{d1, d2},
		},
		{
			intakes: []*p06.Animal{d1, d2, c1},
			want:    []*p06.Animal{d1, d2},
		},
		{
			intakes: []*p06.Animal{c1, d1, c2, d2, c3, d3},
			want:    []*p06.Animal{d1, d2, d3},
		},
	}

	for _, test := range subtests {
		test := test
		name := fmt.Sprintf("%s", test.intakes)

		t.Run(name, func(t *testing.T) {
			shelter := p06.NewShelter()

			for _, a := range test.intakes {
				shelter.Enqueue(a)
			}

			got := []*p06.Animal{}
			for {
				a, err := shelter.DequeueDog()
				if err != nil {
					if errors.Is(err, p06.ErrNoAnimals) {
						break
					}
					t.Fatal(err)
				}

				got = append(got, a)
			}

			if diff := cmp.Diff(test.want, got); diff != "" {
				t.Errorf("(-want +got)\n%s", diff)
			}
		})
	}
}

func TestShelter_DequeueCats(t *testing.T) {
	t.Parallel()

	d1 := &p06.Animal{Name: "dog1", IsDog: true}
	d2 := &p06.Animal{Name: "dog2", IsDog: true}
	d3 := &p06.Animal{Name: "dog3", IsDog: true}

	c1 := &p06.Animal{Name: "cat1"}
	c2 := &p06.Animal{Name: "cat2"}
	c3 := &p06.Animal{Name: "cat3"}

	subtests := []struct {
		intakes []*p06.Animal
		want    []*p06.Animal
	}{
		{
			intakes: []*p06.Animal{},
			want:    []*p06.Animal{},
		},
		{
			intakes: []*p06.Animal{d1},
			want:    []*p06.Animal{},
		},
		{
			intakes: []*p06.Animal{d1, d2},
			want:    []*p06.Animal{},
		},
		{
			intakes: []*p06.Animal{c1},
			want:    []*p06.Animal{c1},
		},
		{
			intakes: []*p06.Animal{c1, c2},
			want:    []*p06.Animal{c1, c2},
		},
		{
			intakes: []*p06.Animal{d1, c1, c2},
			want:    []*p06.Animal{c1, c2},
		},
		{
			intakes: []*p06.Animal{c1, d1, c2},
			want:    []*p06.Animal{c1, c2},
		},
		{
			intakes: []*p06.Animal{c1, c2, d1},
			want:    []*p06.Animal{c1, c2},
		},
		{
			intakes: []*p06.Animal{d1, c1, d2, c2, d3, c3},
			want:    []*p06.Animal{c1, c2, c3},
		},
	}

	for _, test := range subtests {
		test := test
		name := fmt.Sprintf("%s", test.intakes)

		t.Run(name, func(t *testing.T) {
			shelter := p06.NewShelter()

			for _, a := range test.intakes {
				shelter.Enqueue(a)
			}

			got := []*p06.Animal{}
			for {
				a, err := shelter.DequeueCat()
				if err != nil {
					if errors.Is(err, p06.ErrNoAnimals) {
						break
					}
					t.Fatal(err)
				}

				got = append(got, a)
			}

			if diff := cmp.Diff(test.want, got); diff != "" {
				t.Errorf("(-want +got)\n%s", diff)
			}
		})
	}
}

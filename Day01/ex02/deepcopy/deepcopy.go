package deepcopy

import (
	"errors"
)

func DeepCopy(slice_nb []int, lower_i int, upper_i int) ([]int, error) {
	a := []int{};
	
	switch {
	case lower_i > upper_i :
		return a, errors.New("lower index > upper index");	
	case len(slice_nb) < (upper_i + 1) :
		return a, errors.New("upper index strictly superior than len of the slice");	
	case lower_i < 0 :
		return a, errors.New("lower index inferior to zero");	
	}

	for lower_i <= upper_i {
		a = append(a, slice_nb[lower_i]);
		lower_i++;
	}
	return a, nil;
}
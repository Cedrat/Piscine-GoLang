package find 

func SortSlice(list *[]int) {
	def_list := *list;
	len := len(def_list);

	//isn't an exercice for sorting, implement this for go fast in this exercice
	for i := 0 ; i < (len - 1) ; {
		if def_list[i] > def_list[i + 1] {
			def_list[i], def_list[i + 1] = def_list[i + 1] , def_list[i]; 
			i = 0;
		} else {
			i++;
		}
	}
	*list = def_list;
}

func Find(x int, list_numbers ...int) []int {

	var list []int;

	for _, value := range list_numbers {
		if value < x {
			list = append(list, value)
		}
	}

	SortSlice(&list);
	return list;
	
}
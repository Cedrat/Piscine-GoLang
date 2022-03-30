package hackaton

const (
	DevPresent = 0
	DevAbsent  = 1
)

func Find(slice []int, value_to_find int) bool {
	for _, value := range slice {
		if value == value_to_find {
			return true
		}
	}
	return false
}

func IsDevLeft(devs []int) bool {
	for _, dev := range devs {
		if dev == DevPresent {
			return true
		}
	}
	return false
}

func DevsInRelation(devs []int, devs_left []int, nb int) []int {
	var devs_present []int

	for !Find(devs_present, devs[nb]) && devs_left[devs[nb]] == DevPresent {
		devs_present = append(devs_present, devs[nb])
		nb = devs[nb]
	}
	return devs_present
}

func removeDevs(dev_left *[]int, devs_in_relation []int) {
	for _, dev := range devs_in_relation {
		(*dev_left)[dev] = DevAbsent
	}
}

func removeDev(dev_left *[]int, dev int) {
	(*dev_left)[dev] = DevAbsent
}

func NextDev(dev_left []int) int {
	for i, dev := range dev_left {
		if dev == DevPresent {
			return i
		}
	}
	return 0
}

func SortSlice(list *[]int) {
	def_list := *list
	len := len(def_list)

	//isn't an exercice for sorting, implement this for go fast in this exercice
	for i := 0; i < (len - 1); {
		if def_list[i] > def_list[i+1] {
			def_list[i], def_list[i+1] = def_list[i+1], def_list[i]
			i = 0
		} else {
			i++
		}
	}
	*list = def_list
}

func Match(devs []int) [][]int {
	dev_left := make([]int, len(devs))
	var team_list [][]int
	var team_alone []int

	nb_dev := 0

	for IsDevLeft(dev_left) {
		devs_in_relation := DevsInRelation(devs, dev_left, nb_dev)
		SortSlice(&devs_in_relation)
		if devs_in_relation == nil {
			team_alone = append(team_alone, nb_dev)
			removeDev(&dev_left, nb_dev)
			nb_dev = NextDev(dev_left)
		} else {
			team_list = append(team_list, devs_in_relation)
			removeDevs(&dev_left, devs_in_relation)
			nb_dev = NextDev(dev_left)
		}
	}
	team_list = append(team_list, team_alone)
	return team_list
}

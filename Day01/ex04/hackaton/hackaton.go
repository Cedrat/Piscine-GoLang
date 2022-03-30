package hackaton

import "fmt"

func Find(slice []int, value_to_find int) bool {
	for _, value := range slice {
		if value == value_to_find {
			return true
		}
	}
	return false
}

func Match(devs []int) [][]int {
	dev_left := make([]int, len(devs), len(devs))
	fmt.Println(dev_left)
	var team_list [][]int
	// var dev_alone []int

	for i := 0; i < len(devs); i++ {
		var dev_encounter []int
		nb_dev := i

		for i := 0; i < len(devs); i++ {
			if dev_left[i] != -1 {
				nb_dev = i
				break
			}
		}

		for {
			nb_dev = devs[nb_dev]
			if Find(dev_encounter, nb_dev) || Find(dev_left, nb_dev) {
				team_list = append(team_list, dev_encounter)
				for _, i := range dev_encounter {
					dev_left[i] = -1
				}
				fmt.Println(dev_left)
				break
			}
			dev_encounter = append(dev_encounter, nb_dev)
		}
		team_list = append(team_list, dev_encounter)
	}

	return team_list
}

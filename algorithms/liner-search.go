package algorithms

func LinerSearch(datalist[] int, search int) (bool){
	for _, item := range datalist{
		if item == search{
			return true
		}
	}
	return false
}
package whattoeat

type FoodGroup struct {
	GroupName string
	GroupUser []int64
	Food      struct {
		Location string
		Name     string
		Rank     int8 //Ten out of ten XD
		Comment  string
	}
}

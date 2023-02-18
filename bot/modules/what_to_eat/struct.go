package whattoeat

type FoodGroup struct {
	Name string `json:"name"`
	Food []Food `json:"food"`
}

type Database struct {
	Groups []*FoodGroup `json:"groups"`
	Users  []*FoodEater `json:"users"`
}

type FoodEater struct {
	ID        int64  `json:"id"`
	GroupName string `json:"group"`
}

type Food struct {
	Location string
	Name     string
	Rank     int8 //Ten out of ten XD
	Comment  string
}

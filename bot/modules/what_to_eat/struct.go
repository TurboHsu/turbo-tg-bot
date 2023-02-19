package whattoeat

import (
	"time"
)

type FoodGroup struct {
	Name string  `json:"name"`
	Food []*Food `json:"food"`
	// ReviewInterval is the duration between a food recommendation is sent
	// and its interview is requested
	ReviewInterval time.Duration `json:"review_interval"`
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
	Location  string `json:"location"`
	Name      string `json:"name"`
	Rank      int8   `json:"rank"` //Ten out of ten XD
	Comment   string `json:"comment"`
	Thumbnail string `json:"icon"`
}

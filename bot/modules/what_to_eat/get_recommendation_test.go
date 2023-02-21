package whattoeat

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
)

func TestGetRecommendation(t *testing.T) {
	const SAMPLES = 1000000
	const FoodCount = 10
	const Threshold = 0.02

	group := makeGroup("testing")
	group.Food = make([]*Food, FoodCount)
	sum := 0
	for i := 0; i < FoodCount; i++ {
		group.Food[i] = &Food{
			Name: fmt.Sprintf("food#%d", i),
			Rank: int8(rand.Intn(100)) + 1,
		}
		sum += int(group.Food[i].Rank)
	}

	recommendations := make(map[*Food]int, FoodCount)
	for i := 0; i < SAMPLES; i++ {
		food := getRecommendation(&group)
		recommendations[food]++
	}

	errorSum := 0.0
	for _, food := range group.Food {
		count := float64(food.Rank) * float64(SAMPLES) / float64(sum)
		actual := recommendations[food]
		err := math.Abs(count-float64(actual)) / count

		errorSum += err
	}

	averErr := errorSum / FoodCount
	if averErr > Threshold {
		t.Errorf("Average error is %f", averErr)
	}
}

package whattoeat

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/TurboHsu/turbo-tg-bot/utils/log"
)

var Data Database

func Init() error {
	//Read the Data store
	if _, err := os.Stat("./database"); os.IsNotExist(err) {
		err = os.Mkdir("./database", os.ModePerm)
		Data = Database{}
		saveChanges()
		return err
	}
	dbFile, err := os.OpenFile("./database/whattoeat.json", os.O_CREATE|os.O_RDONLY, 0666)
	if err != nil {
		return err
	}
	defer dbFile.Close()
	byteVal, err := io.ReadAll(dbFile)
	if err != nil {
		return err
	}
	json.Unmarshal(byteVal, &Data)
	return nil
}

func saveChanges() {
	//Write to file
	data, err := json.Marshal(Data)
	log.HandleError(err)
	dbFile, err := os.OpenFile("./database/whattoeat.json", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	log.HandleError(err)
	defer dbFile.Close()
	_, err = io.WriteString(dbFile, string(data))
	log.HandleError(err)
}

// FindGroup looks up the database
// and return the exact match in name.
// If no match exists, nil will be returned.
func (data Database) FindGroup(name string) *FoodGroup {
	if name == "" {
		return nil
	}

	for _, group := range data.Groups {
		if group.Name == name {
			return group
		}
	}
	return nil
}

// FindUser looks up the database
// and return the exact match in name.
// Returns nil if no match found.
func (data Database) FindUser(id int64) *FoodEater {
	for _, user := range data.Users {
		if user.ID == id {
			return user
		}
	}
	return nil
}

// Members gets all members belonging to
// one group
func (group FoodGroup) Members(data Database) []*FoodEater {
	var eaters []*FoodEater
	for _, eater := range data.Users {
		if eater.GroupName == group.Name {
			eaters = append(eaters, eater)
		}
	}
	return eaters
}

func makeGroup(name string) FoodGroup {
	return FoodGroup{
		Name:           name,
		Food:           make([]*Food, 0),
		ReviewInterval: 600,
	}
}

func makeFoodEater(id int64) FoodEater {
	return FoodEater{ID: id}
}

func (food *Food) RankString() string {
	return fmt.Sprintf("%.1f / 10", float32(food.Rank)/10)
}

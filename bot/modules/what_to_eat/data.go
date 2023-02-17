package whattoeat

import (
	"encoding/json"
	"io"
	"os"
)

var Database []FoodGroup

func Init() error {
	//Read the data store
	dbFile, err := os.OpenFile("./database/whattoeat.json", os.O_CREATE|os.O_RDONLY, 0666)
	if err != nil {
		return err
	}
	defer dbFile.Close()
	byteVal, err := io.ReadAll(dbFile)
	if err != nil {
		return err
	}
	json.Unmarshal(byteVal, &Database)
	return nil
}

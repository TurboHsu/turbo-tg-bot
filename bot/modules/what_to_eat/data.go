package whattoeat

import (
	"encoding/json"
	"io"
	"os"

	"github.com/TurboHsu/turbo-tg-bot/utils/log"
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

func saveChanges() {
	//Write to file
	data, err := json.Marshal(Database)
	log.HandleError(err)
	dbFile, err := os.OpenFile("./database/whattoeat.json", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	log.HandleError(err)
	defer dbFile.Close()
	_, err = io.WriteString(dbFile, string(data))
	log.HandleError(err)
}

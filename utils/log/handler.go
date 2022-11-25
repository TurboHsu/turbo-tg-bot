package log

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/TurboHsu/turbo-tg-bot/utils/config"
)

func HandleError(err error) {
	if err != nil {
		//Generate log
		errMsg := fmt.Sprintf("%s [E] [%s]", getTime(), err.Error())

		//Handle console output
		if !config.Config.Common.Silent {
			fmt.Println(errMsg)
		}

		if config.Config.Common.WriteLog {
			//Write log to file
			file, innerErr := os.OpenFile(config.Config.Common.LogPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				fmt.Printf("%s [EE] Error when writing log:{%s}, reason is {%s}\n",
					getTime(), err.Error(), innerErr.Error())
				defer file.Close()
				writer := bufio.NewWriter(file)
				_, innerErr = writer.WriteString(errMsg + "\n")
				if innerErr != nil {
					fmt.Printf("%s [EE] Error when writing log:{%s}, reason is {%s}\n",
						getTime(), err.Error(), innerErr.Error())
				}
				writer.Flush()
			}

		}
	}
}

func HandleInfo(msg string) {
	if msg != "" {
		//Generate log
		infoMsg := fmt.Sprintf("%s [I] [%s]", getTime(), msg)

		//Handle console output
		if !config.Config.Common.Silent {
			fmt.Println(infoMsg)
		}

		if config.Config.Common.WriteLog && config.Config.Common.Debug {
			//Write log to file
			file, err := os.OpenFile(config.Config.Common.LogPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				fmt.Printf("%s [EE] Error when writing log:{%s}, reason is {%s}\n",
					getTime(), msg, err.Error())
			}
			defer file.Close()
			writer := bufio.NewWriter(file)
			_, err = writer.WriteString(infoMsg + "\n")
			if err != nil {
				fmt.Printf("%s [EE] Error when writing log:{%s}, reason is {%s}\n",
					getTime(), msg, err.Error())
			}
			writer.Flush()
		}
	}
}

func getTime() string {
	t := time.Now()
	return fmt.Sprintf("[%d-%02d-%02d %02d:%02d:%02d]", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
}

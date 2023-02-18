package whattoeat

import (
	"fmt"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/TurboHsu/turbo-tg-bot/utils/log"
	"math/rand"
)

/*
	Structure:
	Members -> Members Group -> Eat Data -> Random generate -> send
*/

func EatQueryResultHandler(ctx *ext.Context) *gotgbot.InlineQueryResultArticle {
	description, text := foodGenerate(ctx.EffectiveSender.Id())
	return &gotgbot.InlineQueryResultArticle{
		Id:                  "0",
		Title:               "Decide what to eat!",
		Description:         description,
		InputMessageContent: gotgbot.InputTextMessageContent{MessageText: text},
		ThumbUrl:            "https://api.tcloud.site/static/rice.jpg",
		ThumbWidth:          612,
		ThumbHeight:         455,
	}
}

func foodGenerate(uid int64) (description string, text string) {
	return "nil", "nil"
}

// CommandHandler Handles the command
func CommandHandler(bot *gotgbot.Bot, ctx *ext.Context) error {
	parameter := ctx.Args()
	if len(parameter) > 1 {
		senderId := ctx.EffectiveSender.Id()
		var res string
		switch parameter[1] {
		case "group":
			res = handleGroupCommand(senderId, parameter)
		default:
			res = "Unknown action."
		}
		if res != "" {
			_, err := ctx.EffectiveMessage.Reply(bot, res, &gotgbot.SendMessageOpts{})
			log.HandleError(err)
		}
	} else {
		return handleRecommendCommand(bot, ctx)
	}
	return nil
}

func handleGroupCommand(senderId int64, parameter []string) string {
	/*
		Group opeation:
		join <group name> - Join one group. If you are already in one group, then u cant do it. If group does not exist, then create one. If group exists, then join it.
		quit - Quit one group. If you are in one group, then quit it. If the group is empty, then delete it.
		show - Show the current group info.
	*/
	if len(parameter) > 2 {
		switch parameter[2] {
		case "join":
			if len(parameter) > 3 {
				groupName := parameter[3]
				joined, created := joinUser(groupName, senderId)
				if joined {
					saveChanges()
					if created {
						return fmt.Sprintf("Group %s does not exist.\nDon't worry, you have created it!", groupName)
					} else {
						ret := fmt.Sprintf("Joint group [%s] successfully. Now the group contains: [", groupName)
						group := data.FindGroup(groupName)
						members := group.Members(data)
						for _, eater := range members {
							ret += fmt.Sprintf("%d, ", eater.ID)
						}
						ret = ret[:len(ret)-2] + "]."
						return ret
					}
				} else {
					return "You're already in one group. Can't join more!"
				}
			} else {
				return "Too few argument"
			}
		case "quit":
			user := data.FindUser(senderId)
			if user == nil || user.GroupName == "" {
				return "You are not in any of the groups! Go join or create one."
			}
			group := data.FindGroup(user.GroupName)
			user.GroupName = ""
			if len(group.Members(data)) > 0 {
				saveChanges()
				return fmt.Sprintf("You have successfully quit [%s]!", group.Name)
			} else {
				index := -1
				for i, g := range data.Groups {
					if group.Name == g.Name {
						index = i
						break
					}
				}
				if index >= 0 {
					data.Groups = append(data.Groups[:index], data.Groups[index+1:]...)
					saveChanges()
				}
				return fmt.Sprintf("You have successfully quit [%s]. Because the group contains no users, it has beed deleted!", group.Name)
			}
		case "show":
			user := data.FindUser(senderId)
			if user != nil {
				group := data.FindGroup(user.GroupName)
				if group != nil {
					ret := fmt.Sprintf("You are in group [%s], there are these users in the group:[", group.Name)
					members := group.Members(data)
					for _, member := range members {
						ret += fmt.Sprintf("%d, ", member.ID)
					}
					ret = ret[:len(ret)-2] + "]."
					return ret
				}
			}
			return "You are not in any of the groups! Go ahead and join one."
		default:
			return "Unknown action."
		}
	} else {
		return "Too few argument"
	}
}

func handleRecommendCommand(bot *gotgbot.Bot, ctx *ext.Context) error {
	user := data.FindUser(ctx.EffectiveSender.Id())
	if user == nil || user.GroupName == "" {
		_, err := ctx.EffectiveMessage.Reply(bot,
			"You haven't joined a group yet. To get an recommendation, join one.",
			&gotgbot.SendMessageOpts{})
		return err
	}

	group := data.FindGroup(user.GroupName)
	if group == nil {
		ctx.EffectiveMessage.Reply(bot,
			"<i>In internal error occurred</i>",
			&gotgbot.SendMessageOpts{ParseMode: "html"})
		_, err := ctx.EffectiveMessage.Reply(bot,
			fmt.Sprintf("Group %s has been removed, but somehow you're still in it. "+
				"To continue, please run /eat group quit", user.GroupName),
			&gotgbot.SendMessageOpts{})
		return err
	}

	speech := getRecommendation(group).getRecommendationString()
	_, err := ctx.EffectiveMessage.Reply(bot, speech, &gotgbot.SendMessageOpts{})
	return err
}

func getRecommendation(group *FoodGroup) *Food {
	sum := 0
	for _, food := range group.Food {
		sum += int(food.Rank)
	}

	for i, food := range group.Food {
		if i == len(group.Food)-1 || byChance(float32(food.Rank)/float32(sum)) {
			return &food
		}
	}
	return nil
}

func (food *Food) getRecommendationString() string {
	if food == nil {
		return "No recommendations. To get started, run /eat some"
	}
	if food.Location != "" {
		return fmt.Sprintf("Recommending %s at %s, %.1f / 10", food.Name, food.Location, float32(food.Rank)/10)
	} else {
		return fmt.Sprintf("Recommending %s, %.1f / 10", food.Name, float32(food.Rank)/10)
	}
}

func byChance(chance float32) bool {
	return rand.Float32() <= chance
}

// GenerateHelp TODO
func GenerateHelp() string {
	return `
	| Help message.
	| Maybe?
	| WIP
	`
}

// joinUser appends one user to a food group.
// The first return value indicates success
// of the join, while the other indicates whether
// a new group has been created. Changes are not saved.
func joinUser(groupName string, userId int64) (bool, bool) {
	user := data.FindUser(userId)
	// Make sure user exists
	if user == nil {
		user = NewFoodEater(userId)
		data.Users = append(data.Users, user)
	}
	// Make sure the user is not in other groups.
	if user.GroupName != "" && user.GroupName != groupName {
		return false, false
	}

	// Do some joining
	group := data.FindGroup(groupName)
	user.GroupName = groupName
	if group != nil {
		return true, false
	}

	data.Groups = append(data.Groups, NewGroup(groupName))
	return true, true
}

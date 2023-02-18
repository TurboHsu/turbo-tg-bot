package whattoeat

import (
	"fmt"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/TurboHsu/turbo-tg-bot/utils/log"
	"math/rand"
	"strconv"
	"strings"
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
		case "some":
			res = handleAddCommand(senderId, parameter, bot, ctx)
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
				return "Too few arguments"
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
			return "Unknown action"
		}
	} else {
		return "Too few arguments"
	}
}

func handleRecommendCommand(bot *gotgbot.Bot, ctx *ext.Context) error {
	user := data.FindUser(ctx.EffectiveSender.Id())
	if user == nil || user.GroupName == "" {
		_, err := ctx.EffectiveMessage.Reply(bot,
			"You haven't joined a group yet.\nTo get some recommendations, join one.",
			&gotgbot.SendMessageOpts{})
		return err
	}

	group := data.FindGroup(user.GroupName)
	if group == nil {
		sendInternalError(bot, ctx)
		_, err := ctx.EffectiveMessage.Reply(bot,
			fmt.Sprintf("Group %s has been removed, but somehow you're still in it.\n"+
				"To continue, please run /eat group quit", user.GroupName),
			&gotgbot.SendMessageOpts{})
		return err
	}

	speech := getRecommendation(group).getRecommendationString()
	_, err := ctx.EffectiveMessage.Reply(bot, speech, &gotgbot.SendMessageOpts{})
	return err
}

func handleAddCommand(senderId int64, parameter []string, bot *gotgbot.Bot, ctx *ext.Context) string {
	user := data.FindUser(senderId)
	if user == nil || user.GroupName == "" {
		return "You haven't joined a group yet.\nTo recommend a food, join one."
	}
	if len(parameter) < 4 {
		return "Too few arguments. Please input something"
	}

	location := ""
	rate := int8(-1)
	name := ""
	isRate := func(param string) (bool, int) {
		if len(param) <= 0 {
			return false, -1
		}
		slash := -1
		for i, c := range param {
			if c == '/' {
				if i == 0 {
					return false, 0
				}
				if slash >= 0 {
					return false, slash
				}
				slash = i
			} else if c < '0' || c > '9' {
				return false, slash
			}
		}
		return true, slash
	}
	for i := 2; i < len(parameter); i++ {
		param := parameter[i]
		if param == "at" {
			for i+1 < len(parameter) {
				i++
				if r, _ := isRate(parameter[i]); r {
					i--
					break
				}
				location += " "
				location += parameter[i]
			}
		} else if r, slash := isRate(param); r {
			if slash >= 0 {
				deno, err := strconv.ParseFloat(param[slash+1:], 8)
				if err != nil {
					return fmt.Sprintf("%s is not a legal number", param[slash+1:])
				}
				num, err := strconv.ParseFloat(param[:slash], 8)
				if err != nil {
					return fmt.Sprintf("%s is not a legal number", param[:slash])
				}
				rate = int8(num / deno * 100)
			} else {
				num, err := strconv.ParseInt(param, 10, 8)
				if err != nil {
					return fmt.Sprintf("%s is not a legal number", param)
				}
				rate = int8(num * 10)
			}
		} else {
			name += " "
			name += param
		}
	}

	name = strings.Trim(name, " ")
	location = strings.Trim(location, " ")
	if name == "" || rate < 0 {
		return "Too few arguments. Please give me more"
	}

	group := data.FindGroup(user.GroupName)
	if group == nil {
		sendInternalError(bot, ctx)
		return fmt.Sprintf("Group %s has been removed, but somehow you're still in it.\n"+
			"To continue, run /eat group quit", user.GroupName)
	}

	group.Food = append(group.Food, &Food{
		Location: location,
		Name:     name,
		Rank:     rate,
		Comment:  "",
	})
	saveChanges()
	return "Successfully recommended this food"
}

func sendInternalError(bot *gotgbot.Bot, ctx *ext.Context) {
	ctx.EffectiveMessage.Reply(bot,
		"<i>An internal error occurred</i>",
		&gotgbot.SendMessageOpts{ParseMode: "html"})
}

func getRecommendation(group *FoodGroup) *Food {
	sum := 0
	for _, food := range group.Food {
		sum += int(food.Rank)
	}

	for _, food := range group.Food {
		if byChance(float32(food.Rank) / float32(sum)) {
			return food
		}
	}
	if len(group.Food) > 0 {
		for _, food := range group.Food {
			if food.Rank > 0 {
				return food
			}
		}
		return group.Food[0]
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
	return rand.Float32() < chance
}

// GenerateHelp TODO
func GenerateHelp() string {
	return `/eat -- Get a food recommendation
/eat group [quit|show] -- Quit or reveal current group
/eat group join <name> -- Join or create a group
/eat some <name> [at <where>] <rank> -- Recommend some food to your group
		The rank field can be one of the following:
			<x>/<y> -- Fractional format, should be no more than 1.0
			<x> -- 10 based format, should be no more than 10
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

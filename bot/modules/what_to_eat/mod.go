package whattoeat

import (
	"fmt"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

/*
	Structure:
	User -> User Group -> Eat Data -> Random generate -> send
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
	var ret string
	parameter := ctx.Args()
	senderId := ctx.EffectiveSender.Id()
	if len(parameter) > 1 {
		switch parameter[1] {
		case "group":
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
						//Make sure the user is not in other groups.
						inGroup := false
						for _, group := range Database {
							for _, user := range group.GroupUser {
								if senderId == user {
									inGroup = true
									ret = fmt.Sprintf("You are already in group [%s], cannot join more groups.", group.GroupName)
									break
								}
							}
							if inGroup {
								break
							}
						}

						if !inGroup {
							//Do some join
							groupName := parameter[3]
							flag := false
							for i := 0; i < len(Database); i++ {
								if groupName == Database[i].GroupName {
									Database[i].GroupUser = append(Database[i].GroupUser, senderId)
									ret = fmt.Sprintf("Joint group [%s] successfully. Now the group contains: [", groupName)
									for _, uid := range Database[i].GroupUser {
										ret += fmt.Sprintf("%d, ", uid)
									}
									ret = ret[:len(ret)-2] + "]."
									flag = true
								}
							}
							if !flag { //Group does not exist
								Database = append(Database, FoodGroup{
									GroupName: groupName,
									GroupUser: []int64{senderId},
								})
								ret = fmt.Sprintf("Group %s does not exist.\nDon't worry, you have created it!", groupName)
								flag = true
							}

							if flag {
								saveChanges()
							}
						}

					} else {
						ret = "Too few argument"
					}
				case "quit":
					flag := false
					for i := 0; i < len(Database); i++ {
						for _, user := range Database[i].GroupUser {
							if senderId == user && len(Database[i].GroupUser) > 1 {
								flag = true
								//Delete the user from the group
								var newUser []int64
								for _, u := range Database[i].GroupUser {
									if senderId != u {
										newUser = append(newUser, u)
									}
								}
								Database[i].GroupUser = newUser
								ret = fmt.Sprintf("You have successfully quit [%s]!", Database[i].GroupName)
								break
							} else if senderId == user && len(Database[i].GroupUser) == 1 {
								flag = true
								var newDatabase []FoodGroup
								//The user have only this user left, delete it.
								for _, g := range Database {
									if g.GroupName != Database[i].GroupName {
										newDatabase = append(newDatabase, g)
									}
								}
								ret = fmt.Sprintf("You have successfully quit [%s]. Because the group contains no users, it has beed deleted!", Database[i].GroupName)
								Database = newDatabase
								break
							}
						}
						if flag {
							break
						}
					}
					if flag {
						saveChanges()
					} else {
						ret = "You are not in any of the groups! Go join or create one."
					}
				case "show":
					flag := false
					for _, group := range Database {
						for _, user := range group.GroupUser {
							if senderId == user {
								flag = true
								ret = fmt.Sprintf("You are in group [%s], there are these users in the group:[", group.GroupName)
								for _, u := range group.GroupUser {
									ret += fmt.Sprintf("%d, ", u)
								}
								ret = ret[:len(ret)-2] + "]."
								break
							}
						}
						if flag {
							break
						}
					}
					if !flag {
						ret = "You are not in any of the groups! Go ahead and join one."
					}
				default:
					ret = "Unknown action."
				}
			} else {
				ret = "Too few argument"
			}
		default:
			ret = "Unknown action."
		}
	} else {
		ret = "Too few argument."
	}

	ctx.EffectiveMessage.Reply(bot, ret, &gotgbot.SendMessageOpts{})
	return nil
}

// TODO
func GenerateHelp() string {
	return `
	| Help message.
	| Maybe?
	| WIP
	`
}

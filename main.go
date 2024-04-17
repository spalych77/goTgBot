package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var gBot *tgbotapi.BotAPI
var gToken string
var gChatId int64

var gUsersInChat Users

var gUsefulActivities = Activities{
	// это бруклин)
	{"skip", "Прогулять пары (80 минут)", 5},
	{"buy_beer", "Сгонять за пивом для МРК (10 минут)", 7},
	{"cheating", "Списать (20 минут)", 4},
	{"expell", "Отчислиться", 100},
	{"cross_road", "Перейти дорогу в неположенном месте (10 секунд)", 6},
	{"sitting", "Отсидеть пару и ничего не делать", 2},

	//чё за бизнес
	{"attend_courses", "Посетить курсы (пару дней)", 3},
	{"loan", "Дать кенту 200 рублей в столовую (пару сек)", 5},
	{"help", "Дать кенту списать", 3},
	{"project", "Создать проект (несолько часов)", 7},
	{"cleaning", "Убраться в кабинете (10 минут)", 2},
	{"reading", "Прочесть статьи о программировании (1 час)", 1},
}

var gRewards = Activities{
	// это бруклин)
	{"cigarette", "Купить пачку сижек", 100},
	{"beer", "Купить пивка", 100},
	{"hookah", "Сходить с кентами в парилку", 150},

	// чё за бизнес
	{"buy_courses", "Купить курсы", 60},
	{"friends", "Сходить с друзьями куда-нибудь", 40},
	{"business", "Открыть бизнес", 200},
}

type User struct {
	id    int64
	name  string
	coins uint16
}
type Users []*User

type Activity struct {
	code, name string
	coins      uint16
}
type Activities []*Activity

func init() {
	// Uncomment and update token value to set environment variable for Telegram Bot Token given by BotFather.
	// Delete this line after setting the env var. Keep the token out of the public domain!
	_ = os.Setenv(TOKEN_NAME_IN_OS, "6409293985:AAEpyGRZnCJLD_Of6kjiP7sOSI15ThwfOaA")

	if gToken = os.Getenv(TOKEN_NAME_IN_OS); gToken == "" {
		panic(fmt.Errorf(`failed to load environment variable "%s"`, TOKEN_NAME_IN_OS))
	}

	var err error
	if gBot, err = tgbotapi.NewBotAPI(gToken); err != nil {
		log.Panic(err)
	}
	gBot.Debug = true
}

func isStartMessage(update *tgbotapi.Update) bool {
	return update.Message != nil && update.Message.Text == "/start"
}

func isCallbackQuery(update *tgbotapi.Update) bool {
	return update.CallbackQuery != nil && update.CallbackQuery.Data != ""
}

func delay(seconds uint8) {
	time.Sleep(time.Second * time.Duration(seconds))
}

func sendStringMessage(msg string) {
	gBot.Send(tgbotapi.NewMessage(gChatId, msg))
}

func sendMessageWithDelay(delayInSec uint8, message string) {
	sendStringMessage(message)
	delay(delayInSec)
}

func printIntro(update *tgbotapi.Update) {
	sendMessageWithDelay(1, "Предлагаю тебе сыграть в игру и почувствовать на себе, что значит учиться в удивительном месте.. В КТС...")
	sendMessageWithDelay(3, EMOJI_ANGRY)
	sendMessageWithDelay(7, "Здесь ты научишься жизни, а может даже программированию. Да, ты правильно прочёл: жизни. Как списать, прогулять, в целом жить 4 года в своё удовольствие и быть непойманным за руку!")
	sendMessageWithDelay(1, EMOJI_WOW)
	sendMessageWithDelay(5, "В этой мини-игре, взвависимости от выполнения заданий, ты сможешь выбрать 2 пути: быть 'это бруклин, мусо..' или 'чё за бизнес'.")
	sendMessageWithDelay(1, EMOJI_SMILE)
	sendMessageWithDelay(7, `На выбор у тебя есть 2 пункта: 'Совершить какое-то действие (полезное или не очень)' и 'Купить что-то'. Покупка осуществляется за очки репутации🙌, которые ты будешь получать взависимости от выполненных действий!`)
	sendMessageWithDelay(1, EMOJI_REP)
	sendMessageWithDelay(10, `Например, если ты выбрал путь 'это бруклин, мусо..', можешь прогулять пару, потратив 80 минут и получить за это 5 очков респекта! Если же ты выбрал путь 'чё за бизнес', то сможешь сходить на курсы по фронтенду, получив тем самым сертификат 'Конструкторщика сайтов' и 3 очка респекта!`)
	sendMessageWithDelay(3, `Вот и всё! Выполняй задания на свой вкус и основываясь на выбранном пути, трать очки респекта на дроп и развлекайся!`)
	sendMessageWithDelay(1, EMOJI_SUNGLASSES)
}

func getKeyboardRow(buttonText, buttonCode string) []tgbotapi.InlineKeyboardButton {
	return tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData(buttonText, buttonCode))
}

func askToPrintIntro() {
	msg := tgbotapi.NewMessage(gChatId, "Короче, интро. Хочешь скипай, хочешь нет. Скипнешь - будешь тупить, не скипнешь - узнаешь лор и как играть!")
	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		getKeyboardRow(BUTTON_TEXT_PRINT_INTRO, BUTTON_CODE_PRINT_INTRO),
		getKeyboardRow(BUTTON_TEXT_SKIP_INTRO, BUTTON_CODE_SKIP_INTRO),
	)
	gBot.Send(msg)
}

func showMenu() {
	msg := tgbotapi.NewMessage(gChatId, "Выбери один из пунктов:")
	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		getKeyboardRow(BUTTON_TEXT_BALANCE, BUTTON_CODE_BALANCE),
		getKeyboardRow(BUTTON_TEXT_USEFUL_ACTIVITIES, BUTTON_CODE_USEFUL_ACTIVITIES),
		getKeyboardRow(BUTTON_TEXT_REWARDS, BUTTON_CODE_REWARDS),
	)
	gBot.Send(msg)
}

func showBalance(user *User) {
	msg := fmt.Sprintf("%s, твоя репутация на нуле! %s \nЗаймись делами, чтобы заработать +rep", user.name, EMOJI_DOWN)
	if coins := user.coins; coins > 0 {
		msg = fmt.Sprintf("%s, у тебя %d %s", user.name, coins, EMOJI_REP)
	}
	sendStringMessage(msg)
	showMenu()
}

func callbackQueryFromIsMissing(update *tgbotapi.Update) bool {
	return update.CallbackQuery == nil || update.CallbackQuery.From == nil
}

func getUserFromUpdate(update *tgbotapi.Update) (user *User, found bool) {
	if callbackQueryFromIsMissing(update) {
		return
	}

	userId := update.CallbackQuery.From.ID
	for _, userInChat := range gUsersInChat {
		if userId == userInChat.id {
			return userInChat, true
		}
	}
	return
}

func storeUserFromUpdate(update *tgbotapi.Update) (user *User, found bool) {
	if callbackQueryFromIsMissing(update) {
		return
	}

	from := update.CallbackQuery.From
	user = &User{id: from.ID, name: strings.TrimSpace(from.FirstName + " " + from.LastName), coins: 0}
	gUsersInChat = append(gUsersInChat, user)
	return user, true
}

func showActivities(activities Activities, message string, isUseful bool) {
	activitiesButtonsRows := make([]([]tgbotapi.InlineKeyboardButton), 0, len(activities)+1)
	for _, activity := range activities {
		activityDescription := ""
		if isUseful {
			activityDescription = fmt.Sprintf("+ %d %s: %s", activity.coins, EMOJI_REP, activity.name)
		} else {
			activityDescription = fmt.Sprintf("- %d %s: %s", activity.coins, EMOJI_REP, activity.name)
		}
		activitiesButtonsRows = append(activitiesButtonsRows, getKeyboardRow(activityDescription, activity.code))
	}
	activitiesButtonsRows = append(activitiesButtonsRows, getKeyboardRow(BUTTON_TEXT_PRINT_MENU, BUTTON_CODE_PRINT_MENU))

	msg := tgbotapi.NewMessage(gChatId, message)
	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(activitiesButtonsRows...)
	gBot.Send(msg)
}

func showUsefulActivities() {
	showActivities(gUsefulActivities, "Выбери действие или иди домой", true)
}

func showRewards() {
	showActivities(gRewards, "Выбери награду или иди домой", false)
}

func findActivity(activities Activities, choiceCode string) (activity *Activity, found bool) {
	for _, activity := range activities {
		if choiceCode == activity.code {
			return activity, true
		}
	}
	return
}

func processUsefulActivity(activity *Activity, user *User) {
	errorMsg := ""
	if activity.coins == 0 {
		errorMsg = fmt.Sprintf(`дело "%s" столько не стоит`, activity.name)
	} else if user.coins+activity.coins > MAX_USER_COINS {
		errorMsg = fmt.Sprintf("у тебя не может быть больше %d %s", MAX_USER_COINS, EMOJI_REP)
	}

	resultMessage := ""
	if errorMsg != "" {
		resultMessage = fmt.Sprintf("%s, братанчик, сорян, конечно, но %s %s репутации не хватает чёт..", user.name, errorMsg, EMOJI_SAD)
	} else {
		user.coins += activity.coins
		resultMessage = fmt.Sprintf(`%s, "%s" выполнено! %d %s респекта добавлено. %s%s Сейчас утебя %d %s`,
			user.name, activity.name, activity.coins, EMOJI_REP, EMOJI_BICEPS, EMOJI_SUNGLASSES, user.coins, EMOJI_REP)
	}
	sendStringMessage(resultMessage)
}

func processReward(activity *Activity, user *User) {
	errorMsg := ""
	if activity.coins == 0 {
		errorMsg = fmt.Sprintf(`Награда "%s" дороже!`, activity.name)
	} else if user.coins < activity.coins {
		errorMsg = fmt.Sprintf(`у тебя есть %d %s. Ты не можешь позволить "%s" за %d %s`, user.coins, EMOJI_REP, activity.name, activity.coins, EMOJI_REP)
	}

	resultMessage := ""
	if errorMsg != "" {
		resultMessage = fmt.Sprintf("%s братанчик, сорян, конечно, но %s %s репутации не хватает чёт %s", user.name, errorMsg, EMOJI_SAD, EMOJI_DOWN)
	} else {
		user.coins -= activity.coins
		resultMessage = fmt.Sprintf(`%s, держи подарок за выполнение "%s"! %d %s потрачено репутации, теперь у тебя осталось: %d %s`, user.name, activity.name, activity.coins, EMOJI_REP, user.coins, EMOJI_REP)
	}
	sendStringMessage(resultMessage)
}

func updateProcessing(update *tgbotapi.Update) {
	user, found := getUserFromUpdate(update)
	if !found {
		if user, found = storeUserFromUpdate(update); !found {
			sendStringMessage("Unable to identify the user")
			return
		}
	}

	choiceCode := update.CallbackQuery.Data
	log.Printf("[%T] %s", time.Now(), choiceCode)

	switch choiceCode {
	case BUTTON_CODE_BALANCE:
		showBalance(user)
	case BUTTON_CODE_USEFUL_ACTIVITIES:
		showUsefulActivities()
	case BUTTON_CODE_REWARDS:
		showRewards()
	case BUTTON_CODE_PRINT_INTRO:
		printIntro(update)
		showMenu()
	case BUTTON_CODE_SKIP_INTRO:
		showMenu()
	case BUTTON_CODE_PRINT_MENU:
		showMenu()
	default:
		if usefulActivity, found := findActivity(gUsefulActivities, choiceCode); found {
			processUsefulActivity(usefulActivity, user)

			delay(2)
			showUsefulActivities()
			return
		}

		if reward, found := findActivity(gRewards, choiceCode); found {
			processReward(reward, user)

			delay(2)
			showRewards()
			return
		}

		log.Printf(`[%T] !!!!!!!!! ERROR: Unknown code "%s"`, time.Now(), choiceCode)
		msg := fmt.Sprintf("%s, I'm sorry, I don't recognize code '%s' %s Please report this error to my creator.", user.name, choiceCode, EMOJI_SAD)
		sendStringMessage(msg)
	}
}

func main() {
	log.Printf("Authorized on account %s", gBot.Self.UserName)

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = UPDATE_CONFIG_TIMEOUT

	for update := range gBot.GetUpdatesChan(updateConfig) {
		if isCallbackQuery(&update) {
			updateProcessing(&update)
		} else if isStartMessage(&update) {
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
			gChatId = update.Message.Chat.ID
			askToPrintIntro()
		}
	}
}

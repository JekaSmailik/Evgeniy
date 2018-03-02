package main

import (
	"github.com/gin-gonic/gin"
	"github.com/Syfaro/telegram-bot-api"
	"log"
	"os"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	v1 := router.Group("api/v1")
	v1.GET("/start", Starttest)
	v1.GET("/bot", Telegrambot)
	v1.GET("/finish", Finishtest)

	return router
}
func Starttest(c *gin.Context) {
	// Вывод информации на странице
	c.JSON(200, gin.H{"not errors": "this Test!"})

	start := "login to the server"
	JekaBot(start)
}

func Telegrambot(c *gin.Context) {
	// Вывод информации на странице
	c.JSON(200, gin.H{"start telegrambot": "bot jekabot"})

	// бот перенаправляет мои сообщения ему в группу от его имени
	// или отвечат тем же сообщением что ты ему пишешь

	// подключаемся к боту (jekabot) с помощью токена
	bot, err := tgbotapi.NewBotAPI("546364619:AAGEdZQsSaruUnqXdIlKxt3vN6VBuIf0pPc")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	// инициализируем канал, куда будут прилетать обновления от API
	var ucfg tgbotapi.UpdateConfig = tgbotapi.NewUpdate(0)

	// Обработка телеграм каждые 60 секунд проверка канала
	ucfg.Timeout = 6000

	upd, _ := bot.GetUpdatesChan(ucfg)

	// читаем обновления из канала
	for {
		select {
		case update := <-upd:
			// Пользователь, который написал боту
			UserName := update.Message.From.UserName

			// ID чата/диалога.
			// Может быть идентификатором как чата с пользователем
			// (тогда он равен UserID) так и публичного чата/канала
			ChatID := update.Message.Chat.ID

			// Текст сообщения
			Text := update.Message.Text

			log.Printf("[%s] %d %s", UserName, ChatID, Text)

			// Ответим пользователю его же сообщением
			reply := Text
			// Созадаем сообщение
			// println(ChatID)
			//	msg := tgbotapi.NewMessage(ChatID, reply)
			msq := tgbotapi.NewMessage(-190017184, reply)
			// и отправляем его
			//	bot.Send(msg)
			bot.Send(msq)
		}
	}
}

func Finishtest(c *gin.Context) {
	c.JSON(200, gin.H{"exit": "work finish"})

	finish := "Goodbye"
	JekaBot(finish)

	// Завершение программы main
	os.Exit(0)
}

func JekaBot(x string) {
	// подключаемся к боту с помощью токена
	bot, err := tgbotapi.NewBotAPI("546364619:AAGEdZQsSaruUnqXdIlKxt3vN6VBuIf0pPc")
	if err != nil {
		log.Panic(err)
	}
	// Отправка сообщения telegram лично мне
	//	bot.Send(tgbotapi.NewMessage(331814005, "Hi"))
	// Отправка сообщения telegram группе
	bot.Send(tgbotapi.NewMessage(-190017184, x))
}

func main() {
	router := SetupRouter()
	router.Run(":8888")
}

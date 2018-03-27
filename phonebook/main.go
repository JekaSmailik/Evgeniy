package main

import (
	"github.com/jinzhu/gorm"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"net/http"
	"fmt"
	"io/ioutil"
	"strings"
)

// Открытие или создание базы данных
func Database() *gorm.DB {
	// открыть соединение db, если отсутствует то создает новое по указаному пути.
	db, err := gorm.Open("sqlite3", "/Users/rs/workspace/myproject/phonebook/db/phonebook.db")
	if err != nil {
		panic("не удалось подключить базу данных")
	}
	return db
}

// Выполнение маршрутизатора согласно его протокола
func SetupRouter() *gin.Engine {
	// Default возвращается экземпляр Engine с уже прикрепленным промежуточным программным обеспечением Logger и Recovery.
	// Engine - это экземпляр фреймворка, он содержит настройки мультиплексора, промежуточного программного обеспечения и конфигурации.
	router := gin.Default()
	// Group создает новую группу маршрутизаторов. Вы должны добавить все маршруты с обычными средними или одинаковыми префиксами.
	router.GET("/phonebook", PhoneBook)
	router.POST("/phonebook", PhoneBook)
	router.GET("/spisok", SpisokContact)
	router.GET("/spisok/sort")
	router.GET("/add", AddPhones)
	router.POST("/newsave", PhoneBook)
	//router.POST("/delete", PhoneDelete)
	return router
}

// !!!!! Маршрутизатор ко всем файлам в папке (не исполбзуется... пока не используется).
func Conect(c *gin.Context) {
	name := c.Param("./")
	html := c.Param("/Users/rs/workspace/myproject/phonebook/main.go")
	c.String(http.StatusOK, html+name)

	/* Parameters in path
	func main() {
		router := gin.Default()

		// This handler will match /user/john but will not match neither /user/ or /user
		router.GET("/user/:name", func(c *gin.Context) {
			name := c.Param("name")
			c.String(http.StatusOK, "Hello %s", name)
		})

		// However, this one will match /user/john/ and also /user/john/send
		// If no other routers match /user/john, it will redirect to /user/john/
		router.GET("/user/:name/*action", func(c *gin.Context) {
			name := c.Param("name")
			action := c.Param("action")
			message := name + " is " + action
			c.String(http.StatusOK, message)
		})

		router.Run(":8080")
	}*/
}

// Открывает ссылку на вебстраницу и отправляет введенные данные с помощью функции на сохранение в базу данных,
// так же проверяет на повторение и выводит получившийся результат.
func PhoneBook(c *gin.Context) {
	stroka := ""

	// ReadFile Читает файл с именем filename и возврощает содержимое
	b, err := ioutil.ReadFile("./html/phonebook.html")
	if err != nil {
		fmt.Print(err)
	}
	// Конвектирует содержимое в строчку
	stroka = string(b)
	// Заменяет слова в строке
	stroka = strings.Replace(stroka, "{{alreadyExists}}", "", -1)
	stroka = strings.Replace(stroka, "{{addNewContact}}", "", -1)
	stroka = strings.Replace(stroka, "{{Сonnected}}", "", -1)
	stroka = strings.Replace(stroka, "{{addNewTelephone}}", "", -1)

	// При выполнении условия направляет данные с помощью функции на сохранение в базу данных.
	if c.Request.Method == "POST" {
		stroka = SaveContacts(c.PostForm("surname"), c.PostForm("name"), c.PostForm("telphone"))
	} else {
		// Открытие файла по указанному пути.
		c.File("./html/phonebook.html")
	}

	// Отправляет заголовок ответа HTTP с кодом состояния
	c.Writer.WriteHeader(http.StatusOK)
	// Записывает данные в соединение как часть ответа HTTP
	c.Writer.Write([]byte(stroka))
}

func AddPhones(c *gin.Context) {
	var contact Contacts
	var numbers []NumberTelephones

	// Вывод информации о данном контакте в веб странице
	//if c.Request.Method == "GET" {
	contactId, _ := c.GetQuery("contactid")

	db.Where("id = ?", contactId).First(&contact)
	fmt.Println("ID =", contact.ID, ", фамилия =", contact.Surname, ", Имя =", contact.Name)

	db.Where("id_contact = ?", contact.ID).Find(&numbers)

	num := ""
	for _, number := range numbers {
		num += number.NumberTelephone + "<br>"
	}

	contactsee := "<tr><td><h1>" + contact.Name + "</h1></td><td><h1>" + contact.Surname + "</h1></td></tr><tr><td width=\"100%\" colspan=\"100%\" align=\"center\">" + num + "</td></tr>"

	// ReadFile Читает файл с именем filename и возврощает содержимое
	b, err := ioutil.ReadFile("./html/add.html")
	if err != nil {
		fmt.Print(err)
	}

	// Конвектирует содержимое в строчку
	str := string(b)
	// Заменяет слова в строке
	str = strings.Replace(str, "{{add}}", contactsee, -1) // -1 заменяет все слова в строке
	str = strings.Replace(str, "{{surname}}", fmt.Sprint(contact.Surname), -1)
	str = strings.Replace(str, "{{name}}", fmt.Sprint(contact.Name), -1)
	// Отправляет заголовок ответа HTTP с кодом состояния
	c.Writer.WriteHeader(http.StatusOK)
	// Записывает данные в соединение как часть ответа HTTP
	c.Writer.Write([]byte(str))
	//c.Writer.Write([]byte(stroka45))
}

func SpisokContact(c *gin.Context) {
	// Получает контакты из базы
	var contacts []Contacts

	// Сортировка по алфавиту --------->
	sortSurname, _ := c.GetQuery("sortSurname")
	sortName, _ := c.GetQuery("sortName")

	// Find находит записи, соответствующие заданным условиям
	db.Find(&contacts)

	if sortSurname == "1" {
		db.Order("surname ASC").Find(&contacts)
	} else if sortSurname == "2" {
		db.Order("surname DESC").Find(&contacts)
	}

	if sortName == "1" {
		db.Order("name ASC").Find(&contacts) // сортировка по возрастанию (ASC) или по убыванию (DESC)
	} else if sortName == "2" {
		db.Order("name DESC").Find(&contacts)
	}

	// Получает контакты из базы
	var numbers []NumberTelephones

	list := ""

	for _, contact := range contacts {
		addPhone := "<a href=\"/add?contactid=" + fmt.Sprint(contact.ID) + "\">добавить</a>"
		db.Where("id_contact = ?", contact.ID).Find(&numbers)
		num := ""
		for _, number := range numbers {
			num += number.NumberTelephone + "<br>"
		}
		list = list + "<tr><td>" + contact.Surname + "</td><td>" + contact.Name + "</td><td>" + num + "</td><td>" + addPhone + "</td>" + "</tr>"
	}

	// ReadFile Читает файл с именем filename и возврощает содержимое
	b, err := ioutil.ReadFile("./html/spisok.html")
	if err != nil {
		fmt.Print(err)
	}
	// Конвектирует содержимое в строчку
	content := string(b)

	if sortName == "1" {
		content = strings.Replace(content, "{{sortName}}", "2", -1)
		fmt.Println("sortName =", sortName)
		sortirovkaName := "Список отсортирован по возростанию Имен"
		content = strings.Replace(content, "{{sortListSurname}}", sortirovkaName, -1)
	} else {
		content = strings.Replace(content, "{{sortName}}", "1", -1)
		if sortName == "2" {
			fmt.Println("sortName =", sortName)
			sortirovkaName := "Список отсортирован по убыванию Имен"
			content = strings.Replace(content, "{{sortListSurname}}", sortirovkaName, -1)
		}
	}

	if sortSurname == "1" {
		content = strings.Replace(content, "{{sortSurname}}", "2", -1)
		fmt.Println("sortSurname =", sortSurname)
		sortirovkaSurname := "Список отсортирован по возростанию Фамилий"
		content = strings.Replace(content, "{{sortListSurname}}", sortirovkaSurname, -1)
	} else {
		content = strings.Replace(content, "{{sortSurname}}", "1", -1)
		if sortSurname == "2" {
			fmt.Println("sortSurname =", sortSurname)
			sortirovkaSurname := "Список отсортирован по убыванию Фамилий"
			content = strings.Replace(content, "{{sortListSurname}}", sortirovkaSurname, -1)
		}
	}

	// Заменяет слова в строке
	content = strings.Replace(content, "{{sortListSurname}}", "", -1)
	content = strings.Replace(content, "{{html}}", list, -1) // -1 заменяет все слова в строке

	// Отправляет заголовок ответа HTTP с кодом состояния
	c.Writer.WriteHeader(http.StatusOK)
	// Записывает данные в соединение как часть ответа HTTP
	c.Writer.Write([]byte(content))

}

// Добавляет в базу данных новые контакты, без повторения
func SaveContacts(surnames string, names string, telphone string) (string) {
	number := NumberTelephones{}
	contact := Contacts{}
	repeatContact := Contacts{}

	// ReadFile Читает файл с именем filename и возврощает содержимое
	b, err := ioutil.ReadFile("./html/phonebook.html")
	if err != nil {
		fmt.Print(err)
	}
	// Конвектирует содержимое в строчку
	content := string(b)
	// Заменяет слова в строке

	// Where филтрует записи в базе данных с указанными условиями
	if db.Where("surname = ? AND name = ?", surnames, names).First(&contact); contact.ID != 0 {
		fmt.Println("Контакт", contact.Surname, contact.Name, " уже создан")

		Repeat := "<tr><td>" + "Контакт " + contact.Surname + contact.Name + " уже создан" + "</td></tr>"
		content = strings.Replace(content, "{{alreadyExists}}", Repeat, -1) // -1 заменяет все слова в строке
		content = strings.Replace(content, "{{addNewContact}}", "", -1)     // -1 заменяет все слова в строке

	} else {
		contact.Surname = surnames
		contact.Name = names
		db.Save(&contact)
		fmt.Println("Сохранен контакт", surnames, names)

		Repeat := "<tr><td>" + "Сохранен контакт " + surnames + names + "</td></tr>"
		content = strings.Replace(content, "{{alreadyExists}}", "", -1)     // -1 заменяет все слова в строке
		content = strings.Replace(content, "{{addNewContact}}", Repeat, -1) // -1 заменяет все слова в строке
	}

	if db.Where("number_telephone = ?", telphone).First(&number); number.ID != 0 {
		db.Where("id = ?", number.IDContact).First(&repeatContact)
		fmt.Println("Телефон", telphone, "привязан к", repeatContact.Surname, repeatContact.Name)

		Repeat := "<tr><td>" + "Телефон " + telphone + " привязан к " + repeatContact.Surname + repeatContact.Name + "</td></tr>"
		content = strings.Replace(content, "{{Сonnected}}", Repeat, -1)   // -1 заменяет все слова в строке
		content = strings.Replace(content, "{{addNewTelephone}}", "", -1) // -1 заменяет все слова в строке

	} else {
		IDphone := NumberTelephones{
			NumberTelephone: telphone,
			IDContact:       contact.ID,
		}
		db.Save(&IDphone)
		fmt.Println("К контакту", surnames, names, "добавлен телефон", telphone)

		Repeat := "<tr><td>" + "Добавлен телефон " + telphone + " к контакту " + contact.Surname + contact.Name + "</td></tr>"
		content = strings.Replace(content, "{{Сonnected}}", "", -1)           // -1 заменяет все слова в строке
		content = strings.Replace(content, "{{addNewTelephone}}", Repeat, -1) // -1 заменяет все слова в строке
	}

	return content
}

// База данных контактов телефонной книги
type Contacts struct {
	gorm.Model     // Определение базовой модели модели, включая поля `ID`,` CreatedAt`, `ОбновленAt`,` DeletedAt`, которые могут быть встроены в ваши модели
	Surname string // Фамилия контакта
	Name    string // Имя контакта
}

// База данных телефонов контактов
type NumberTelephones struct {
	gorm.Model             // Определение базовой модели модели, включая поля `ID`,` CreatedAt`, `ОбновленAt`,` DeletedAt`, которые могут быть встроены в ваши модели
	NumberTelephone string // Номер телефона
	IDContact       uint   // ID контакта
}

var db *gorm.DB

func
main() {
	db = Database()
	defer db.Close()
	// AutoMigrate запускает автоматическую миграцию для данных моделей, добавляет только отсутствующие поля, не удаляет / не изменяет текущие данные
	db.AutoMigrate(&Contacts{})
	db.AutoMigrate(&NumberTelephones{})

	router := SetupRouter()
	router.Run(":8888") // Run подключает маршрутизатор к http.Server и начинает прослушивать и обслуживать HTTP-запросы.
}

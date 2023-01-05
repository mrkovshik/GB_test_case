package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	_ "github.com/lib/pq"
)

var showQry string = " SELECT vacancies.id, vacancy_name, key_skills, salary, vacancy_desc, job_types.job_type FROM vacancies JOIN job_types ON vacancies.job_type = job_types.id;"
var searchQry string = " SELECT vacancies.id, vacancy_name, key_skills, salary, vacancy_desc, job_types.job_type FROM vacancies JOIN job_types ON vacancies.job_type = job_types.id WHERE vacancy_name ILIKE '%"

const (
	host     = "172.19.97.209" //впишите сюда параметры своей базы данных
	port     = 5432
	user     = "postgres"
	password = "17pasHres19!"
	dbname   = "vacancies"
)

func connectDB() *sql.DB {
	// Connect to the database
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		log.Fatal(err)
	}

	return db
}

func showVacs(qry string) {

	if qry != showQry {
		fmt.Println("Введите название вакансии частично или полностью, либо наберите \"назад\" для выхода в предыдущее меню")
		scanner := bufio.NewScanner(os.Stdin)
		for {
			fmt.Print("> ")
			if scanner.Scan() {
				if scanner.Text() == "назад" {
					return
				}
				qry = searchQry + scanner.Text() + "%';"
				break
			}
		}
	}
	db := connectDB()
	defer db.Close()
	rows, err := db.Query(qry)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	counter:=0
	headerPrinted:=false
	for rows.Next() {
		counter++
		var id, salary int
		var vacancy_name, key_skills, vacancy_desc, job_type string
		err = rows.Scan(&id, &vacancy_name, &key_skills, &salary, &vacancy_desc, &job_type)
		if err != nil {
			panic(err.Error())
		}
		if counter>0&&!headerPrinted{
			fmt.Println("ID    |   Vacancy Name   |  Key Skills                                          |  Vacancy Description                                                     |  Salary    |  Job Type")
		fmt.Println("-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------")
		headerPrinted=true
	}
		fmt.Printf("%-4v  |  %-14v  |  %-50v  |  %-70v  |  %-8v  |  %-10v\n", id, vacancy_name, key_skills, vacancy_desc, salary, job_type)
	}
	if counter == 0 {
		fmt.Println("\nПохоже, по такому запросу в базе ничего не нашлось. Попробуйте изменить запрос")
		fmt.Println("----------------------------------------")
		return
	}
	err = rows.Err()
	if err != nil {
		panic(err.Error())
	}
}

func insert() {
	var queryString []string
	var sal, jobCode int
	var err error

	fmt.Println("введите соответствующие значения строк, разделяя их знаком \"/\": ")
	fmt.Println("название вакансии, ключевые навыки, описание вакансии, зарплата, и код типа работы: 1 для работы в офисе, 2 для удаленной работы и 3 для гибридной")
	fmt.Println("Например: \"Охранник/Решать конфликтные ситуации, обращаться с оружием, разгадывать сканворды/Человек, который следит за порядком в офисном здании/50000/1\"")
	fmt.Println("или наберите \"назад\" для выхода в предыдущее меню")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		if scanner.Scan() {
			if scanner.Text() == "назад" {
				return
			}

			queryString = strings.Split(scanner.Text(), "/")
			if len(queryString) != 5 {
				fmt.Println("Неверное количество аргументов, повторите ввод")
				continue
			}
			sal, err = strconv.Atoi(queryString[3])
			if err != nil {
				fmt.Println("Ошибка ввода данных в поле \"Зарплата\", повторите ввод")
				continue
			}
			jobCode, err = strconv.Atoi(queryString[4])
			if err != nil {
				fmt.Println("Ошибка ввода данных в поле \"код типа работы\", повторите ввод")
				continue
			}

		}
		db := connectDB()
		defer db.Close()
		// Prepare the INSERT statement
		stmt, err := db.Prepare("INSERT INTO vacancies (vacancy_name,key_skills, vacancy_desc ,  salary, job_type) VALUES($1, $2,$3,$4,$5)")
		if err != nil {
			log.Fatal(err)
		}
		defer stmt.Close()

		// Execute the INSERT statement
		_, err = stmt.Exec(queryString[0], queryString[1], queryString[2], sal, jobCode)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Запись добавлена")
		return
	}
}

func main() {
	var mainMenu string = "\n*****************************\n - Если хотите посмотреть всю таблицу вакансий, наберите \"посмотреть\", \n - Если хотите найти вакансию по названию наберите \"найти\"\n - Если хотите добавить строку - наберите \"добавить\", \n - Если хотите выйти из программы, наберите \"выход\"\n*****************************\n"
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println(mainMenu)
	// Loop indefinitely
OuterLoop:
	for {
		// Print a prompt
		fmt.Print("> ")

		// Scan for input
		if scanner.Scan() {
			switch {
			case scanner.Text() == "посмотреть":
				showVacs(showQry)
				fmt.Println(mainMenu)
			case scanner.Text() == "найти":
				showVacs("")
				fmt.Println(mainMenu)
			case scanner.Text() == "добавить":
				insert()
				fmt.Println(mainMenu)
			case scanner.Text() == "выход":
				fmt.Println("Всего хорошего!")
				break OuterLoop
				default: fmt.Println("Неверно введена команда, попробуйте еще раз")
			}
		}
	}
}

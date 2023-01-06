package main

import (
	"bufio"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	_ "github.com/lib/pq"
)

var mainMenu string = "\n*****************************\n - Если хотите посмотреть всю таблицу вакансий, наберите \"посмотреть\", \n - Если хотите найти вакансию по названию наберите \"найти\"\n - Если хотите добавить строку - наберите \"добавить\", \n - Если хотите выйти из программы, наберите \"выход\"\n*****************************\n"
var searchQry string = " SELECT vacancies.id, vacancy_name, key_skills, salary, vacancy_desc, job_types.job_type FROM vacancies JOIN job_types ON vacancies.job_type = job_types.id WHERE vacancy_name ILIKE '%"
var db *sql.DB

type dbCred struct {
	host     string
	port     int
	user     string
	password string
	dbName   string
}

type vacQery struct {
	vacName   string
	keySkills string
	vacDesc   string
	salary    int
	jobCode   int
}

func connectDB(cred dbCred) (*sql.DB, error) {

	// Connect to the database
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		cred.host, cred.port, cred.user, cred.password, cred.dbName)
	fmt.Println(psqlInfo)

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	return db, err
}

func searchDialog() (string, bool) {
	fmt.Println("Введите название вакансии частично или полностью, либо наберите \"назад\" для выхода в предыдущее меню")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		if scanner.Scan() {
			if scanner.Text() == "назад" {
				return "", false
			}
			searchKey := scanner.Text()
			return searchKey, true
		}
	}
}

func showVacs(qry string, db *sql.DB) error {
	rows, err := db.Query(qry)
	if err != nil {
		return err
	}
	defer rows.Close()
	counter := 0
	for rows.Next() {
		counter++
		var id, salary int
		var vacancy_name, key_skills, vacancy_desc, job_type string
		err = rows.Scan(&id, &vacancy_name, &key_skills, &salary, &vacancy_desc, &job_type)
		if err != nil {
			return err
		}
		if counter == 1 {
			fmt.Println("ID    |   Vacancy Name   |  Key Skills                                          |  Vacancy Description                                                     |  Salary    |  Job Type")
			fmt.Println("-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------")

		}
		fmt.Printf("%-4v  |  %-14v  |  %-50v  |  %-70v  |  %-8v  |  %-10v\n", id, vacancy_name, key_skills, vacancy_desc, salary, job_type)
	}
	if counter == 0 {
		fmt.Println("\nПохоже, по такому запросу в базе ничего не нашлось. Попробуйте изменить запрос")
		fmt.Println("----------------------------------------")
		return err
	}
	err = rows.Err()
	if err != nil {
		return err
	}
	return err
}
func insertDialog() (vacQery, bool) {
	var result vacQery
	var err error
	fmt.Println("введите соответствующие значения строк, разделяя их знаком \"/\": ")
	fmt.Println("название вакансии, ключевые навыки, описание вакансии, зарплата, и код типа работы: 1 для работы в офисе, 2 для удаленной работы и 3 для гибридной формы работы")
	fmt.Println("Например: \"Охранник/Решать конфликтные ситуации, обращаться с оружием, разгадывать сканворды/Человек, который следит за порядком в офисном здании/50000/1\"")
	fmt.Println("или наберите \"назад\" для выхода в предыдущее меню")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		if scanner.Scan() {
			if scanner.Text() == "назад" {
				return vacQery{}, false
			}

			queryString := strings.Split(scanner.Text(), "/")
			if len(queryString) != 5 {
				fmt.Println("Неверное количество аргументов, повторите ввод")
				continue
			}
			result.vacName = queryString[0]
			result.keySkills = queryString[1]
			result.vacDesc = queryString[2]
			result.salary, err = strconv.Atoi(queryString[3])
			if err != nil {
				fmt.Println("Ошибка ввода данных в поле \"Зарплата\", повторите ввод")
				continue
			}
			result.jobCode, err = strconv.Atoi(queryString[4])
			if err != nil {
				fmt.Println("Ошибка ввода данных в поле \"код типа работы\", повторите ввод")
				continue
			}
			if result.jobCode > 3 || result.jobCode < 1 {
				fmt.Println("Код работы может быть только следующих значений: 1 для работы в офисе, 2 для удаленной работы и 3 для гибридной формы работы. Ввод других значений не допускается")
				continue
			}
		}

		return result, true

	}
}
func insert(q vacQery, db *sql.DB) error {
	fmt.Println(q)
	stmt, err := db.Prepare("INSERT INTO vacancies (vacancy_name,key_skills, vacancy_desc ,  salary, job_type) VALUES($1, $2,$3,$4,$5)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(q.vacName, q.keySkills, q.vacDesc, q.salary, q.jobCode)
	if err != nil {
		return err
	}
	stmt.Close()
	fmt.Println("Запись добавлена")
	return nil
}

func mainDialog() error {
	fmt.Println(mainMenu)
	scanner := bufio.NewScanner(os.Stdin)
OuterLoop:
	for {
		// Print a prompt
		fmt.Print("> ")

		// Scan for input
		if scanner.Scan() {
			switch {
			case scanner.Text() == "посмотреть":
				err := showVacs(searchQry+"%'", db)
				if err != nil {
					fmt.Println("Ошибка обращения к базе данных", err)
					return err
				}
				fmt.Println(mainMenu)
			case scanner.Text() == "найти":
				keyWord, proceed := searchDialog()
				if proceed {
					err := showVacs(searchQry+keyWord+"%'", db)
					if err != nil {
						fmt.Println("Ошибка обращения к базе данных", err)
						return err
					}
				}
				fmt.Println(mainMenu)
			case scanner.Text() == "добавить":
				query, proceed := insertDialog()
				if proceed {
					err := insert(query, db)
					if err != nil {
						fmt.Println("Ошибка внесения данных в таблицу", err)
						return err
					}
				}
				fmt.Println(mainMenu)
			case scanner.Text() == "выход":
				fmt.Println("Всего хорошего!")
				break OuterLoop
			default:
				fmt.Println("Неверно введена команда, попробуйте еще раз")
			}
		}
	}
	return nil
}
func main() {
	var err error
	var cred dbCred
	flag.IntVar(&cred.port, "port", 5432, "Port for DB connection")
	flag.StringVar(&cred.host, "h", "localhost", "DB host IP")
	flag.StringVar(&cred.password, "p", "my_awesome_password", "DB connection password")
	flag.StringVar(&cred.user, "u", "postgres", "DB connection user name")
	flag.StringVar(&cred.dbName, "n", "vacancies", "DB name")

	db, err = connectDB(cred)
	if err != nil {
		fmt.Println("Ошибка подключения базы данных", err)
		return
	}
	defer db.Close()
	mainDialog()

}

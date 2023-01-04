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

const (
	host     = "172.19.97.209" //впишите сюда параметры своей базы данных
	port     = 5432
	user     = "postgres"
	password = "my_awesome_password"
	dbname   = "my_awesome_DB"
)

func connectDB()*sql.DB{
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



func showVacs (){
    db:=connectDB()
    defer db.Close()
    rows, err := db.Query("SELECT * FROM vacancies")
    if err != nil {
        panic(err.Error())
    }
    defer rows.Close()
    fmt.Println("ID    |   Vacancy Name   |  Key Skills                                          |  Vacancy Description                                                     |  Salary    |  Job Type")
    fmt.Println("-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------")
    for rows.Next() {
        var id, salary, job_type int
        var vacancy_name, key_skills,vacancy_desc string
            err = rows.Scan(&id, &vacancy_name, &key_skills, &salary, &job_type, &vacancy_desc)
        if err != nil {
            panic(err.Error())
        }
        fmt.Printf("%-4v  |  %-14v  |  %-50v  |  %-70v  |  %-8v  |  %-10v\n", id, vacancy_name, key_skills,vacancy_desc, salary, job_type)
    }
    err = rows.Err()
    if err != nil {
        panic(err.Error())
    }
}



func insert (vacName string, skills string, vacDesc string, salary int, jobType int){
db:=connectDB()
defer db.Close()
	// Prepare the INSERT statement
	stmt, err := db.Prepare("INSERT INTO vacancies (vacancy_name,key_skills, vacancy_desc ,  salary, job_type) VALUES($1, $2,$3,$4,$5)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	// Execute the INSERT statement
	_, err = stmt.Exec(vacName,skills,vacDesc,salary,jobType)
	if err != nil {
		log.Fatal(err)


}
}
func main (){
	var mainMenu string = "Если хотите посмотреть таблицу вакансий, наберите \"посмотреть\", если хотите добавить строку - наберите \"добавить\", если хотите выйти из программы, наберите \"выход\""
	scanner := bufio.NewScanner(os.Stdin)
fmt.Println(mainMenu)
	// Loop indefinitely
	OuterLoop:
	for {

		
		// Print a prompt
		fmt.Print("> ")

		// Scan for input
		if scanner.Scan() {
			switch{
			case scanner.Text()=="посмотреть":
				showVacs()
				fmt.Println("\n", mainMenu)
			case scanner.Text()=="добавить":
				fmt.Println("введите соответствующие значения строк, разделяя их знаком / (название вакансии, ключевые навыки, описание вакансии, зарплата, и код типа работы) или \"назад\" для выхода в предыдущее меню")
				innerLoop:
				for{
					fmt.Print("> ")
					if scanner.Scan() {
						if scanner.Text()=="назад"{
							fmt.Println(mainMenu)
								break innerLoop
						}else{
							queryString:= strings.Split(scanner.Text(), "/")
							if len(queryString)!=5{
								fmt.Println("Неверное количество аргументов, повторите ввод")
continue
							}
							sal,err:=strconv.Atoi(queryString[3])
							if err!=nil {
								fmt.Println("Ошибка ввода данных в поле \"Зарплата\", повторите ввод")
								continue
							}
							jobCode,err :=strconv.Atoi(queryString[4])
							if err!=nil {
								fmt.Println("Ошибка ввода данных в поле \"код типа работы\", повторите ввод")
								continue
							}
							insert(queryString[0], queryString[1],queryString[2], sal ,jobCode)
							fmt.Println("Запись добавлена")
							fmt.Println(mainMenu)
							break innerLoop																	}
						}
				}
				
			case scanner.Text()=="выход":
				fmt.Println("Всего хорошего!")
				break OuterLoop
			}
			}
		}
}
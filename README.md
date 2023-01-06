# GB_test_case
## тестовое задание для проекта Geek Brains

Код написан на Go для базы данных PostgreSQL. 
Для начала работы назначьте константам host, port, user, password, dbname значения, соответствующие вашей базе данных.

Для удобства ниже указан запрос для создания таблиц в PostgreSQL:

```
CREATE TABLE job_types (id SERIAL PRIMARY KEY,
                                          job_type text NOT NULL);


CREATE TABLE vacancies (id SERIAL PRIMARY KEY,
                                          vacancy_name text NOT NULL,
                                                            key_skills text NOT NULL,
                                                                            vacancy_desc text NOT NULL,
                                                                                              salary integer NOT NULL,
                                                                                                             job_type integer REFERENCES job_types(id));


INSERT INTO job_types (job_type)
VALUES ('В офисе'),
       ('Удаленно'),
       ('Гибридный');```

       разделить функции на смысловые блоки, вынести подключение к базе в мэйн, добавить пинг дб перед выполнение функции, обращающейся к дб. переделать подключение к базе через флаги. создать миграцию и положить в репу. Миграция для создания бд, create if not exist database, switch toda
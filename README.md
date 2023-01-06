# GB_test_case#
## тестовое задание для проекта Geek Brains ##

Код написан на Go для базы данных PostgreSQL. 
При запуске бинарного файла для настройки соединения к базе данных укажите следующие флаги:
- -h (адрес хоста, по умолчанию "localhost")
- -u (DB user name, по умолчанию "postgres")
- -n (DB name, по умолчанию "vacancies")
- -p (DB password, по умолчанию "my_awesome_password")
- -port (DB port, по умолчанию "5432")

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

   
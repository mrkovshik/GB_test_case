# GB_test_case
test exercise for the geek brains project

Код написан на Go для базы данных PostgreSQL. 
Для начала работы назначьте константам host, port, user, password, dbname значения, соответствующие вашей базе данных.

Для удобства ниже указан запрос для создания таблиц в PostgreSQL:



CREATE TABLE job_types (
    id SERIAL PRIMARY KEY,
    job_type text NOT NULL
);

CREATE TABLE vacancies (
    id SERIAL PRIMARY KEY,
  vacancy_name text NOT NULL,
  key_skills text NOT NULL,
  vacancy_desc text NOT NULL,
  salary integer NOT NULL, 
  job_type integer references job_types(id)
);

Insert into job_types (job_type)
values ('В офисе'), ('Удаленно'), ('Гибридный');

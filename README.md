# leadgen

Задание (Golang + PostgreSQL + Gin)

Написать сервис с тремя эндпоинтами:

1) принимает на вход Строение (Building) и записывает его в бд (поля Building: название, город, год сдачи, кол-во этажей)
2) возвращает список строений (Buildings), с возможностью фильтрации по городу, году и кол-ву этажей (параметры не обязательные)
3) OpenApi документация (генерировать из аннотаций например с помощью https://github.com/swaggo/swag)

Настройки соединения с Postgres читать из config файла:

host
port
user
password
db

решение должно быть в git-репозитории (опубликовать на github, gitlab...)


curl --silent 'http://localhost:8080/buildings?city=Moscow&year_built=2020&floor_count=10'

curl --silent 'http://localhost:8080/buildings' \
--header 'Content-Type: application/json' \
--data '{"name":"Building B112", "city":"Perm", "year_built":2021, "floor_count":10}'


go test ./internal/test/... 
go test ./internal/test/... -bench=.

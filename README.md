# Library

Веб сервис разработанный на архитектуре REST для проведения учебных НТ

## Что из себя представляет?
Представляет из себя библиотеку в котором пользователи регистрируются а после могут брать книги

## Endpoints
На данный момент реализованы следующие эндпоинты

![image](https://user-images.githubusercontent.com/67442103/177352914-319dc0a0-d136-4e40-8025-9c82af58be87.png)

1. Регистрация
2. Авторизация
3. Запрос на получение всех книг пользователя
4. Запрос на получение всех книг из библиотеки
5. Добавление книг в библиотеку
6. Добавление книг в список пользователя
7. Запрос на получение книг по id
8. Удаление книг из списка пользователя

## Какие компоненты используются?
1. Написан на Go
2. БД Postgres
3. Для регистрации используется JWT-токен

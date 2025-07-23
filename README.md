### Fullstack Web Calculator
Проект, который состоит из фронта на JavaScript(React) и бэкенда на Go.     
Веб‑интерфейс позволяет выполнять базовые арифметические операции, а серверный REST‑API принимает выражения, вычисляет результат и сохраняет историю запросов.      
Чтобы запустить проект необходимо установить: NodeJS, компилятор Golang, Docker.    

#### frontend
`cd .\frontend\`    
Команда для установки зависимостей: `npm i`     
Команда для запуска сайта: `npm start`   

#### backend
`cd .\backend\`       
Пример команды для запуска Docker контейнера с PostgreSQL:   
`docker run --name postgres-container -e POSTGRES_PASSWORD=password -d -p 5432:5432 postgres`   
Команда для запуска сервера: `go run .\cmd\app\main.go`   
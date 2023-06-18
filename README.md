# cli-test

Это приложение командной строки (CLI), 
которое позволяет отправлять строку на
указанный URL и получать самую длинную 
подстроку без повторяющихся символов из ответа сервера.

## Использование
Шаги и запустить приложение:
- clone:
```
git clone https://github.com/iKayrat/cli-test.git
```
- для запуска compose файла докер:
```
make up
```
- для запуска:
```
arg1=(*строка*) arg2=(*url*) make run
```
eg:
```
arg1=abcabcbb arg2="http://localhost:8081/api/substring" make run
```
or
```
go run cli.go abcabcbb http://localhost:8081/api/substring
```
порт:
**http://localhost:8081/api/substring**

```
go test
```

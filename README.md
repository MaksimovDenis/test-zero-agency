# Запуск  

1. Сколнировать репозиторий:
```bash   
git clone https://github.com/MaksimovDenis/test-zero-agency.git
```
2. Перейти в директорию проекта (если Вы не в ней):  
```bash    
cd test-zero-agency 
```
3. Из дериктории проекта выполнить команды:  
```bash      
docker-compose build
```
```bash    
docker-compose up 
``` 
4. После чего swagger докумнтация будет доступна по ссылке:  
```bash      
http://localhost:8000/swagger/index.html
```  
![image](https://github.com/MaksimovDenis/vk_restAPI/assets/44647373/e089ba3b-4ce3-4731-a2f7-8fda13e16aa4)

- Для авторизации необходимо зарегестрироваться, залогиниться и получить JWT токен.

# Реализация 
- Был реализован REST API сервер, с подходом The Clean Architecture, в рамках тестового задания.     
`https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html`  
- Доступные методы:   
1. Регестрация. 
2. Получение списка всех новостей.  
3. Получение новости по ID.  
4. Редактирование новости по ID.  
5. Удаление новости по ID.  

## Технологии и зависимости  
- Язык программирование: Golang 1.21.6   
- PostgreSQL latest  
- Для реализации http сервера использовались библиотека GIN.  

Инструменты:  
- Docker: Для контейнеризации и управления средой разработки.  
- Docker Compose: Для управления многоконтейнерными приложениями.  

Основные зависимости:  
- [gin-gonic/gin](https://github.com/gin-gonic/gin) v1.9.1
- [golang-jwt/jwt](https://github.com/golang-jwt/jwt) v4.5.0
- [jmoiron/sqlx](https://github.com/jmoiron/sqlx) v1.3.5
- [lib/pq](https://github.com/lib/pq) v1.10.9
- [sirupsen/logrus](https://github.com/sirupsen/logrus) v1.9.3
- [swaggo/files](https://github.com/swaggo/files) v1.0.1
- [swaggo/gin-swagger](https://github.com/swaggo/gin-swagger) v1.6.0
- [swaggo/swag](https://github.com/swaggo/swag) v1.16.3
- [yaml.v3](https://pkg.go.dev/gopkg.in/yaml.v3) v3.0.1
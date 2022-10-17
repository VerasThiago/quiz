## Running locally
```bash
docker-compose up
```

## Routes

### Quiz
```
GET    /api/v0/quiz/user/list   
POST   /api/v0/quiz/user/create 
DELETE /api/v0/quiz/user/delete 
PUT    /api/v0/quiz/user/update 
POST   /api/v0/quiz/user/publish
GET    /api/v0/quiz/published/list
GET    /api/v0/quiz/published/open/:quizid
POST   /api/v0/quiz/published/submit
POST   /api/v0/question/user/create
GET    /api/v0/question/user/list/:quizid
DELETE /api/v0/question/user/delete
PUT    /api/v0/question/user/update
POST   /api/v0/option/user/create
GET    /api/v0/option/user/list/:questionid
DELETE /api/v0/option/user/delete
PUT    /api/v0/option/user/update
GET    /api/v0/submission/user/list
GET    /api/v0/submission/user/report/:submissionid
GET    /api/v0/submission/list/:quizid
```

### Login
```
 POST   /login/v0/user/register   
 POST   /login/v0/user/login      
 DELETE /login/v0/admin/delete    
 PUT    /login/v0/admin/update    
```

### Database models
![](https://i.imgur.com/UMIArJG.png)

### Archtecture
![](https://i.imgur.com/nGxBTfM.png)

### Scripts
On `/scripts` contain a script that create 2 users (Thiago and Jhon) with 2 quizzes for each with 2 questions each

```
cd shared/scripts/ && ./populatedb.sh
```
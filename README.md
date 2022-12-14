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

On `/scripts` contain a script that create 2 users (Thiago and Jhon) with 2
quizzes for each with 2 questions each

```
cd shared/scripts/ && ./populatedb.sh
```

### TODO

- [ ] HTTPS
- [x] Disable user update is_published through PUT /quiz
- [x] Disable user create single option that bug single select
- [x] Disable user submit duplicated ids
- [ ] Enable pagination
- [x] Remove ID and leave only id
- [x] Hide specific fields from user
- [x] Validate all endpoints
- [x] Check if user already tried this quiz
- [x] Check all status codes of all responses
- [x] Check if user is choosing more than 1 option for single ans question
- [x] Refactor function "userCanExec..." to Validate class
  - [x] ValidateSyntax
  - [x] ValidateSemantic
- [x] Every quiz has a title and consists of 1-10 questions.
- [x] Every question has 1-5 possible answers.
- [x] Nginx
- [x] Users should also be able to see the solutions of other users to their own
      quizzes so they can get the statistics of how people perform on their
      quiz.
- [x] Handle db erros
- [x] Return error on submit with wrong id's
- [x] Refactor db name to Repository
- [x] Index models
- [x] Dockerfile multi stage build
- [x] Create report async (queue)
- [ ] Change Publish quiz method from POST to PUT

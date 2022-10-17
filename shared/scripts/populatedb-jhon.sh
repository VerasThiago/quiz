#!/usr/bin/env bash

username="jhon"
email="jhon@gmail.com"
password="123"

curl --location --request POST 'localhost:8081/login/v0/user/register' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "testuser",
    "username": "'"$username"'",
    "email": "'"$email"'",
    "password": "'"$password"'"
}'

TOKEN="$(curl --location --request POST 'localhost:8081/login/v0/user/login' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email": "'"$email"'",
    "password": "'"$password"'"
}' | jq -r '.token')"

echo "TOKEN from $usename:"
echo $TOKEN

geographyQuizId="$(curl --location --request POST 'localhost:8080/api/v0/quiz/user/create' \
--header "Authorization: $TOKEN" \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "Geography quiz",
    "Description": "Super quiz to test you geo knowladge"
}' | jq -r '.id')"

echo "Geography Quiz:"
echo $geographyQuizId

geographyQuestion1Id="$(curl --location --request POST 'localhost:8080/api/v0/question/user/create' \
--header "Authorization: $TOKEN" \
--header 'Content-Type: application/json' \
--data-raw '{
    "question": {
        "name": "What is the capital of Brazil?", 
        "description": "Random description",
        "quizid": "'"$geographyQuizId"'",
        "type": 1
    },
    "optionlist":[
        {
            "value": "Brasilia",
            "correctness": true
        },
        {
            "value": "California",
            "correctness": false
        },
        {
            "value": "Sidney",
            "correctness": false
        },
        {
            "value": "Camberra",
            "correctness": false
        }
    ]
}' | jq -r '.id')"

echo "Geography Question1Id:"
echo $geographyQuestion1Id

geographyQuestion2Id="$(curl --location --request POST 'localhost:8080/api/v0/question/user/create' \
--header "Authorization: $TOKEN" \
--header 'Content-Type: application/json' \
--data-raw '{
    "question": {
        "name": "What is the capital of Australia?", 
        "description": "Random description",
        "quizid": "'"$geographyQuizId"'",
        "type": 1
    },
    "optionlist":[
        {
            "value": "Brasilia",
            "correctness": false
        },
        {
            "value": "California",
            "correctness": false
        },
        {
            "value": "Sidney",
            "correctness": false
        },
        {
            "value": "Camberra",
            "correctness": true
        }
    ]
}' | jq -r '.id')"

echo "Geography Question2Id:"
echo $geographyQuestion2Id

historyQuizId="$(curl --location --request POST 'localhost:8080/api/v0/quiz/user/create' \
--header "Authorization: $TOKEN" \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "History quiz",
    "Description": "Super quiz to test you history knowladge"
}' | jq -r '.id')"

echo "History Quiz:"
echo $historyQuizId

historyQuestion1Id="$(curl --location --request POST 'localhost:8080/api/v0/question/user/create' \
--header "Authorization: $TOKEN" \
--header 'Content-Type: application/json' \
--data-raw '{
    "question": {
        "name": "When brazil got independent?", 
        "description": "Random description",
        "quizid": "'"$historyQuizId"'",
        "type": 1
    },
    "optionlist":[
        {
            "value": "September 7th, 1822",
            "correctness": true
        },
        {
            "value": "September 7th, 1775",
            "correctness": false
        },
        {
            "value": "February 18th, 1822",
            "correctness": false
        },
        {
            "value": "November 25th, 1997",
            "correctness": false
        },
        {
            "value": "February 18th, 1775",
            "correctness": false
        }
    ]
}' | jq -r '.id')"

echo "History Question1Id:"
echo $historyQuestion1Id

historyQuestion2Id="$(curl --location --request POST 'localhost:8080/api/v0/question/user/create' \
--header "Authorization: $TOKEN" \
--header 'Content-Type: application/json' \
--data-raw '{
    "question": {
        "name": "When was WW2?", 
        "description": "Random description",
        "quizid": "'"$historyQuizId"'",
        "type": 1
    },
    "optionlist":[
        {
            "value": "1935 ~ 1949",
            "correctness": false
        },
        {
            "value": "1845 ~ 1860",
            "correctness": false
        },
        {
            "value": "1939 ~ 1945",
            "correctness": true
        },
        {
            "value": "1715 ~ 1735",
            "correctness": false
        }
    ]
}' | jq -r '.id')"

echo "History Question2Id:"
echo $historyQuestion2Id

echo "TOKEN from $usename:"
echo $TOKEN
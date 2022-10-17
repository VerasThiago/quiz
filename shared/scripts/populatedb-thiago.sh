#!/usr/bin/env bash

username="thiago"
email="thiago@gmail.com"
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

echo "TOKEN:"
echo $TOKEN

astronomyQuizID="$(curl --location --request POST 'localhost:8080/api/v0/quiz/user/create' \
--header "Authorization: $TOKEN" \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "Astronomy quiz",
    "Description": "Super hard quiz to test you Astronomy knowladge"
}' | jq -r '.id')"

echo "Astronomy Quiz:"
echo $astronomyQuizID

astronomyQuestion1ID="$(curl --location --request POST 'localhost:8080/api/v0/question/user/create' \
--header "Authorization: $TOKEN" \
--header 'Content-Type: application/json' \
--data-raw '{
    "question": {
        "name": "Which planet(s) weight(s) more than 50 times than earth?", 
        "description": "Random description",
        "quizid": "'"$astronomyQuizID"'",
        "type": 2
    },
    "optionlist": [
        {
            "value": "Mercury",
            "correctness": false
        },
        {
            "value": "Venus",
            "correctness": false
        },
        {
            "value": "Mars",
            "correctness": true
        },
        {
            "value": "Jupiter",
            "correctness": true
        },
        {
            "value": "Saturn",
            "correctness": false
        }
    ]
}' | jq -r '.id')"

echo "Astronomy Question1ID:"
echo $astronomyQuestion1ID

astronomyQuestion2ID="$(curl --location --request POST 'localhost:8080/api/v0/question/user/create' \
--header "Authorization: $TOKEN" \
--header 'Content-Type: application/json' \
--data-raw '{
    "question": {
        "name": "Pluto is a planet?", 
        "description": "Random description",
        "quizid": "'"$astronomyQuizID"'",
        "type": 1
    },
    "optionlist": [
        {
            "value": "Yes",
            "correctness": false
        },
        {
            "value": "No",
            "correctness": true
        }
    ]
}' | jq -r '.id')"

echo "Astronomy Question2ID:"
echo $astronomyQuestion2ID

mathQuizID="$(curl --location --request POST 'localhost:8080/api/v0/quiz/user/create' \
--header "Authorization: $TOKEN" \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "Math quiz",
    "Description": "Super hard quiz to test you math knowladge"
}' | jq -r '.id')"

echo "Math Quiz:"
echo $mathQuizID

mathQuestion1ID="$(curl --location --request POST 'localhost:8080/api/v0/question/user/create' \
--header "Authorization: $TOKEN" \
--header 'Content-Type: application/json' \
--data-raw '{
    "question": {
        "name": "Solve the given the equation: 4x - 7(2 - x) = 3x + 2", 
        "description": "Random description",
        "quizid": "'"$mathQuizID"'",
        "type": 1
    },
    "optionlist": [
        {
            "value": "x = 2",
            "correctness": true
        },
        {
            "value": "x = 4",
            "correctness": false
        },
        {
            "value": "x = -4",
            "correctness": false
        },
        {
            "value": "x = -6",
            "correctness": false
        }
    ]
}' | jq -r '.id')"

echo "Math Question1ID:"
echo $mathQuestion1ID

mathQuestion2ID="$(curl --location --request POST 'localhost:8080/api/v0/question/user/create' \
--header "Authorization: $TOKEN" \
--header 'Content-Type: application/json' \
--data-raw '{
    "question": {
        "name": "Solve the given the equation: (4 - 2z)/3 = 3/4 - 5z/6", 
        "description": "Random description",
        "quizid": "'"$mathQuizID"'",
        "type": 1
    },
    "optionlist": [
        {
            "value": "3",
            "correctness": false
        },
        {
            "value": "14/8",
            "correctness": false
        },
        {
            "value": "-7/2",
            "correctness": true
        },
        {
            "value": "-13/5",
            "correctness": false
        }
    ]
}' | jq -r '.id')"

echo "TOKEN:"
echo $TOKEN
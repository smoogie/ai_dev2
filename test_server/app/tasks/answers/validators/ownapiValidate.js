const buildAnswerResponse = require("../buildAnswerResponse");
const axios = require('axios');

const ownapiValidate = function(data) {
    const url = data.answer
    axios.post(url, {
            question: "Nad jakim mozrzem leży Polska",
        }).then((res) => sendSecondQuestionRequest(res, url)).catch((error) => {
            console.error(error)
        })
    //TODO: switch to async and return response based on requests validations
    const isCorrect = true
    return buildAnswerResponse(isCorrect)
}

function sendSecondQuestionRequest(res, url) {
    console.log(`statusCode: ${res.statusCode}`)
    console.log(res)

    axios.post(url, {
        question: "Podaj nazwę stolicy Polski jako jedno słowo. Zwróć odpowiedź w formacie JSON w polu \"nazwa\". Nie dodawaj żadnego komentarza ani formatowania. Oczekuję formatu: {\"nazwa\":\"xxxxxxxcxxx\"}",
    }).then((res2) => {
        console.log(`statusCode: ${res2.statusCode}`)
        console.log(res2)
    }).catch((error) => {
        console.error(error)
    })
}

module.exports = ownapiValidate

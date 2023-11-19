const buildAnswerResponse = require("../buildAnswerResponse");
const axios = require("axios");

const ownapiproValidate = function(data) {

    const url = data.answer
    axios.post(url, {
        question: "Jaka jest największa rzeka Polski?",
    }).then((res) => sendNextQuestionRequest(res, url, 0)).catch((error) => {
        console.error(error)
    })
    //TODO: switch to async and return response based on requests validations
    const isCorrect = true
    return buildAnswerResponse(isCorrect)
}

const queations = [
    "Jestem Franek.",
    "Jaka mam na imie?",
    "Gdzie jest Amazonka?",
    "Ma 30 lat?",
    "Ile mam lat?"
]

function logResponse(res) {
    console.log(`statusCode: ${res.statusCode}`)
    console.log(res)
}

function sendNextQuestionRequest(res, url, count) {
    logResponse(res)
    console.log(`Send next question - ${count+2}`)
    axios.post(url, {
        question: queations[count],
    }).then((res2) => {
        if (count < queations.length-1) {
            sendNextQuestionRequest(res2, url, count+1)
        } else {
            logResponse(res2)
        }
    }).catch((error) => {
        console.error(error)
    })
}

module.exports = ownapiproValidate

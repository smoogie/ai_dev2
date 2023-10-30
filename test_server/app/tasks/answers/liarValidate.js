const buildAnswerResponse = require("./buildAnswerResponse");
const liarValidate = function(data) {
    const answer = data.answer
    //TODO: based on randomization from data generation validate if true/false
    let isAnswerCorrect = data.answer === "NO"
    return buildAnswerResponse(isAnswerCorrect)
}

module.exports = liarValidate

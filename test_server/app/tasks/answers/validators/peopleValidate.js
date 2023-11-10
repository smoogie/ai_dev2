const buildAnswerResponse = require("../buildAnswerResponse");

const peopleValidate = function(data) {
    const answer = data.answer
    //TODO: compare to real result with AI
    const isCorrect = answer.includes("indygo")
    return buildAnswerResponse(isCorrect)
}

module.exports = peopleValidate

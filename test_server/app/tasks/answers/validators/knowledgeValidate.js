const buildAnswerResponse = require("../buildAnswerResponse");

const knowledgeValidate = function(data) {
    const answer = data.answer
    //TODO: get question from memory, ask ai to validate response for question
    const isCorrect = answer.length > 0
    return buildAnswerResponse(isCorrect)
}

module.exports = knowledgeValidate

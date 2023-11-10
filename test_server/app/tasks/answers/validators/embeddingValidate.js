const buildAnswerResponse = require("./buildAnswerResponse");
const embeddingValidate = function(data) {
    const answer = data.answer
    //TODO: maybe check if correct values?
    const isEnoughBigArray = Array.isArray(answer) && answer.length == 1536
    return buildAnswerResponse(isEnoughBigArray)
}

module.exports = embeddingValidate

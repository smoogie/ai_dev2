const buildAnswerResponse = require("../buildAnswerResponse");

const gnomeValidate = function(data) {
    const answer = data.answer
    //TODO: get info about the image from memory, and validate response
    const isCorrect = answer.length > "4"
    return buildAnswerResponse(isCorrect)
}

module.exports = gnomeValidate

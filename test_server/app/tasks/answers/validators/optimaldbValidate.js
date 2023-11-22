const buildAnswerResponse = require("../buildAnswerResponse");
const optimaldbValidate = function(data) {
    const answer = data.answer
    //TODO: Use AI to review questions for data
    let isAnswerCorrect = true
    return buildAnswerResponse(isAnswerCorrect)
}

module.exports = optimaldbValidate

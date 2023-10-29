const buildAnswerResponse = require("./buildAnswerResponse");
const moderationValidate = function(data) {
    const answer = data.answer
    let isAnswerCorrect = false
    return buildAnswerResponse(isAnswerCorrect)
}

module.exports = moderationValidate

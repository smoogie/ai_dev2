const buildAnswerResponse = require("../buildAnswerResponse");
const inpromptValidate = function(data) {
    const answer = data.answer
    let isAnswerCorrect = data.answer === "Parpeć"
    return buildAnswerResponse(isAnswerCorrect)
}

module.exports = inpromptValidate

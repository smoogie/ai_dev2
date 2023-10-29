const buildAnswerResponse = require("./buildAnswerResponse");
const liarValidate = function(data) {
    const answer = data.answer
    let isAnswerCorrect = false
    return buildAnswerResponse(isAnswerCorrect)
}

module.exports = liarValidate

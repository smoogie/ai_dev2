const buildAnswerResponse = require("../buildAnswerResponse");
const googleValidate = function(data) {
    const answer = data.answer
    //TODO: Send some request to provided url and verify results
    let isAnswerCorrect = true
    return buildAnswerResponse(isAnswerCorrect)
}

module.exports = googleValidate

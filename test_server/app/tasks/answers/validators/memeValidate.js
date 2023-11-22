const buildAnswerResponse = require("../buildAnswerResponse");
const memeValidate = function(data) {
    const answer = data.answer
    //TODO: Use AI to compare answer and transcription
    let isAnswerCorrect = true
    return buildAnswerResponse(isAnswerCorrect)
}

module.exports = memeValidate

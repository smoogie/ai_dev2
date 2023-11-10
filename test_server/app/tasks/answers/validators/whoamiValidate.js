const buildAnswerResponse = require("./buildAnswerResponse");

const whoamiValidate = function(data) {
    const answer = data.answer
    //TODO: compare to real result with AI
    const isCorrect = answer.includes("Steve Jobs")
    return buildAnswerResponse(isCorrect)
}

module.exports = whoamiValidate

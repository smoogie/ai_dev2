const buildAnswerResponse = require("./buildAnswerResponse");
const bloggerValidate = function(data) {
    const answer = data.answer
    let isAnswerCorrect = false
    return buildAnswerResponse(isAnswerCorrect)
}

module.exports = bloggerValidate

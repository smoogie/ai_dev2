const buildAnswerResponse = require("../buildAnswerResponse");

const searchValidate = function(data) {
    const answer = data.answer
    //TODO: Use AI to check if answer is correct
    return buildAnswerResponse(true)
}

module.exports = searchValidate

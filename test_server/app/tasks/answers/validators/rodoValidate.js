const buildAnswerResponse = require("../buildAnswerResponse");

const rodoValidate = function(data) {
    const answer = data.answer
    //TODO: Use AI to check if answer used as user input in open AI will return good data
    return buildAnswerResponse(true)
}

module.exports = rodoValidate

const buildAnswerResponse = require("./buildAnswerResponse");

const scraperValidate = function(data) {
    const answer = data.answer
    //TODO: Use AI to check if answer is correct based on file
    return buildAnswerResponse(true)
}

module.exports = scraperValidate

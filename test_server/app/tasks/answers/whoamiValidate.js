const buildAnswerResponse = require("./buildAnswerResponse");

const whoamiValidate = function(data) {
    const answer = data.answer
    //TODO: compare to real result
    return buildAnswerResponse(true)
}

module.exports = whoamiValidate

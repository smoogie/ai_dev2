const buildAnswerResponse = require("./buildAnswerResponse");
const correctAnswer = [0,1,1,1]
const moderationValidate = function(data) {
    const answer = data.answer
    let isAnswerCorrect = false
     if (Array.isArray(answer) && answer.length == 4) {
         isAnswerCorrect = answer.reduce(
             (isAllCorrect, response, index) => isAllCorrect && (response === correctAnswer[index]),
             true
         )
     }
    return buildAnswerResponse(isAnswerCorrect)
}

module.exports = moderationValidate

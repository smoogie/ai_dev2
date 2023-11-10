const correctAnswer = require('../../data/json/helloapiData').cookie
const buildAnswerResponse = require('../buildAnswerResponse')
const helloapiValidate = function(data) {
    const answer = data.answer
    const isAnswerCorrect = answer === correctAnswer
    return buildAnswerResponse(isAnswerCorrect)
}

module.exports = helloapiValidate

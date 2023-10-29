const buildAnswerResponse= function(isAnswerCorrect) {
    const response = {
        code: isAnswerCorrect ? 0 : 1,
        msg: isAnswerCorrect ? "OK" : "BAD",
        note: isAnswerCorrect ? "CORRECT" : "wrong answer"
    }
    return response
}

module.exports = buildAnswerResponse

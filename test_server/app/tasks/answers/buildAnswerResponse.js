const buildAnswerResponse= function(isAnswerCorrect) {
    const response = {
        code: isAnswerCorrect ? 0 : -777,
        msg: isAnswerCorrect ? "OK" : "this is NOT the correct answer",
        note: isAnswerCorrect ? "CORRECT" : "wrong answer"
    }
    return response
}

module.exports = buildAnswerResponse

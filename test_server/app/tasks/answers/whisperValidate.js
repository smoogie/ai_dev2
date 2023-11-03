const buildAnswerResponse = require("./buildAnswerResponse");
const transcription = "Test transkrypcji do transkrypcji. Raz, dwa, trzy. Teraz ty!"
const whisperValidate = function(data) {
    const answer = data.answer
    //TODO: Use AI to compare answer and transcription
    let isAnswerCorrect = answer == transcription
    return buildAnswerResponse(isAnswerCorrect)
}

module.exports = whisperValidate

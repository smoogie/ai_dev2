const buildAnswerResponse = require("../buildAnswerResponse");
const bloggerData = require('../../data/json/bloggerData')
const bloggerValidate = function(data) {
    const answer = data.answer
    let isAnswerCorrect = false
    if (answer.length == bloggerData.blog.length) {
        isAnswerCorrect = answer.reduce(
            (areChaptersCorrect, chapter, index) => areChaptersCorrect && validateChapter(chapter, bloggerData.blog[index]),
            true
        )
    }
    return buildAnswerResponse(isAnswerCorrect)
}

function validateChapter(chapter, description) {
    //here we should validate if each chapter has correct text, but we simplify it to check length
    //TODO: use AI to validate responses
    return (typeof chapter === 'string' && chapter.length > 100)
}

module.exports = bloggerValidate

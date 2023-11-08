
require('dotenv').config()
const filePath = process.env.HOST + "/scraper/text_pizza_history.txt"

const scraperData = {
    code: 0,
    msg: "Return answer for the question in POLISH language, based on provided article. Maximum length for the answer is 200 characters",
    input: filePath,
    question: "z którego roku pochodzi łaciński dokument, który pierwszy raz wspomina o pizzy? "
}

module.exports = scraperData

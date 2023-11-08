const helloapi = require('./json/helloapiData')
const blogger = require('./json/bloggerData')
const moderation = require('./json/moderationData')
const inprompt = require('./json/inpromptData')
const embedding = require('./json/embeddingData')
const whisper = require('./json/whisperData')
const functions = require('./json/functionsData')
const rodo = require('./json/rodoData')
const scraper = require('./json/scraperData')

const mapping = {
    helloapi,
    blogger,
    moderation,
    inprompt,
    embedding,
    whisper,
    functions,
    rodo,
    scraper
}

module.exports = (task) => mapping[task]

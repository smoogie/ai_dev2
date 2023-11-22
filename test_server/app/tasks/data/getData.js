const helloapi = require('./json/helloapiData')
const blogger = require('./json/bloggerData')
const moderation = require('./json/moderationData')
const inprompt = require('./json/inpromptData')
const embedding = require('./json/embeddingData')
const whisper = require('./json/whisperData')
const functions = require('./json/functionsData')
const rodo = require('./json/rodoData')
const scraper = require('./json/scraperData')
const search = require('./json/searchData')
const people = require('./json/peopleData')
const ownapi = require('./json/ownapiData')
const ownapipro = require('./json/ownapiproData')
const meme = require('./json/memeData')

const mapping = {
    helloapi,
    blogger,
    moderation,
    inprompt,
    embedding,
    whisper,
    functions,
    rodo,
    scraper,
    search,
    people,
    ownapi,
    ownapipro,
    meme
}

module.exports = (task) => mapping[task]

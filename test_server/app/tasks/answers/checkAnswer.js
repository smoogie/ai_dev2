const helloapi = require('./validators/helloapiValidate')
const blogger = require('./validators/bloggerValidate')
const moderation = require('./validators/moderationValidate')
const liar = require('./validators/liarValidate')
const inprompt = require('./validators/inpromptValidate')
const embedding = require('./validators/embeddingValidate')
const whisper = require('./validators/whisperValidate')
const functions = require('./validators/functionsValidate')
const rodo = require('./validators/rodoValidate')
const scraper = require('./validators/scraperValidate')
const whoami = require('./validators/whoamiValidate')
const search = require('./validators/searchValidate')
const people = require('./validators/peopleValidate')

const mapping = {
    helloapi,
    blogger,
    moderation,
    liar,
    inprompt,
    embedding,
    whisper,
    functions,
    rodo,
    scraper,
    whoami,
    search,
    people
}

module.exports = (task, data) => mapping[task](data)

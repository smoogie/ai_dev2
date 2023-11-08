const helloapi = require('./helloapiValidate')
const blogger = require('./bloggerValidate')
const moderation = require('./moderationValidate')
const liar = require('./liarValidate')
const inprompt = require('./inpromptValidate')
const embedding = require('./embeddingValidate')
const whisper = require('./whisperValidate')
const functions = require('./functionsValidate')
const rodo = require('./rodoValidate')
const scraper = require('./scraperValidate')
const whoami = require('./whoamiValidate')

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
    whoami
}

module.exports = (task, data) => mapping[task](data)

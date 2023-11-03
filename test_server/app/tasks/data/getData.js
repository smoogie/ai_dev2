const helloapi = require('./helloapiData')
const blogger = require('./bloggerData')
const moderation = require('./moderationData')
const inprompt = require('./inpromptData')
const embedding = require('./embeddingData')
const whisper = require('./whisperData')

const mapping = {
    helloapi,
    blogger,
    moderation,
    inprompt,
    embedding,
    whisper
}

module.exports = (task) => mapping[task]

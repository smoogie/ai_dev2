const helloapi = require('./helloapiData')
const blogger = require('./bloggerData')
const moderation = require('./moderationData')
const inprompt = require('./inpromptData')
const embedding = require('./embeddingData')

const mapping = {
    helloapi,
    blogger,
    moderation,
    inprompt,
    embedding
}

module.exports = (task) => mapping[task]

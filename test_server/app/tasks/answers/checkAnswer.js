const helloapi = require('./helloapiValidate')
const blogger = require('./bloggerValidate')
const moderation = require('./moderationValidate')
const liar = require('./liarValidate')
const inprompt = require('./inpromptValidate')
const embedding = require('./embeddingValidate')
const whisper = require('./whisperValidate')

const mapping = {
    helloapi,
    blogger,
    moderation,
    liar,
    inprompt,
    embedding,
    whisper
}

module.exports = (task, data) => mapping[task](data)

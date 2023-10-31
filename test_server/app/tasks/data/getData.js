const helloapi = require('./helloapiData')
const blogger = require('./bloggerData')
const moderation = require('./moderationData')
const inprompt = require('./inpromptData')

const mapping = {
    helloapi,
    blogger,
    moderation,
    inprompt
}

module.exports = (task) => mapping[task]

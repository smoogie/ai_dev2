const helloapi = require('./helloapiData')
const blogger = require('./bloggerData')
const moderation = require('./moderationData')
const liar = require('./liarData')

const mapping = {
    helloapi,
    blogger,
    moderation,
    liar
}

module.exports = (task) => mapping[task]

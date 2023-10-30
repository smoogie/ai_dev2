const helloapi = require('./helloapiData')
const blogger = require('./bloggerData')
const moderation = require('./moderationData')

const mapping = {
    helloapi,
    blogger,
    moderation
}

module.exports = (task) => mapping[task]

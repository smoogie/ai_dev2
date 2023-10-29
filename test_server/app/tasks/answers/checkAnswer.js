const helloapi = require('./helloapiValidate')
const blogger = require('./bloggerValidate')
const moderation = require('./moderationValidate')
const liar = require('./liarValidate')

const mapping = {
    helloapi,
    blogger,
    moderation,
    liar
}

module.exports = (task, data) => mapping[task](data)

const liar = require('./liarData')

const mapping = {
    liar
}

module.exports = (task, req) => mapping[task](req)

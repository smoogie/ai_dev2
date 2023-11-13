const whoami = require('./functionData/whoamiData')
const knowledge = require('./functionData/knowledgeData')

const mapping = {
    whoami,
    knowledge
}

module.exports = (task) => mapping[task]

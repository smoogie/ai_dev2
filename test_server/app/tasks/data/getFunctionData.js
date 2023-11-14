const whoami = require('./functionData/whoamiData')
const knowledge = require('./functionData/knowledgeData')
const tools = require('./functionData/toolsData')

const mapping = {
    whoami,
    knowledge,
    tools
}

module.exports = (task) => mapping[task]

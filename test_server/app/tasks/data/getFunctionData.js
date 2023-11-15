const whoami = require('./functionData/whoamiData')
const knowledge = require('./functionData/knowledgeData')
const tools = require('./functionData/toolsData')
const gnome = require('./functionData/gnomeData')

const mapping = {
    whoami,
    knowledge,
    tools,
    gnome
}

module.exports = (task) => mapping[task]

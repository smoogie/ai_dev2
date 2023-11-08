const whoami = require('./functionData/whoamiData')

const mapping = {
    whoami
}

module.exports = (task) => mapping[task]

const getTaskFromToken = function(req, res) {
    const now = Date.now()
    const token = req.params.token
    if (typeof global.activeTasks[token] === 'undefined'){
        res.status(404).send()
        console.log(`Requested data for not started task`)
        return ""
    }
    const taskToken = global.activeTasks[token]
    const task = taskToken.task
    const time = now-taskToken.time
    console.log(`Token valid from: ${time/1000}`)
    if (time > 120000) {
        res.status(404).send()
        console.log(`Token for task is older than 120 s: ${time/1000}`)
        return ""
    }
    return task
}

module.exports = getTaskFromToken

const testAPIKey = "some-test-key"
const taskMap = require('../tasks/taskTokenMap')
/**
 * @param {*} req
 * @param {Request<P, ResBody, ReqBody, ReqQuery, LocalsObj>|http.ServerResponse} res
 */
const getTaskToken = function(req, res){
    const task = req.params.task
    if (req.body.apikey !== testAPIKey){
        res.status(401).send()
        console.log(`Wrong API key passed on token request`)
        return
    }
    if (!taskMap.hasOwnProperty(task)) {
        res.status(404).send()
        console.log(`Task ${task} does not exist`)
        return
    }
    const token = taskMap[task]
    global.activeTasks[token] = {
        task,
        time: Date.now()
    };
    console.log(`Token for ${task} requested`)
    res.send({code:0, msg:"OK", token})
};

module.exports = getTaskToken

const getData = require('../tasks/data/getData')
const getTaskFromToken = require('./getTaskFromToken')
const {MethodSupportedInTasks, MethodGET} = require('../tasks/methodSupportInTasks')
/**
 * @param {*} req
 * @param {Request<P, ResBody, ReqBody, ReqQuery, LocalsObj>|http.ServerResponse} res
 */
const getTaskData = function(req, res){
    const task = getTaskFromToken(req, res)
    if (task.length <= 0) {
        return
    }
    if (!MethodSupportedInTasks[task].includes(MethodGET)) {
        res.status(404).send()
        console.log(`Requested sent on the wrong HTTP method`)
    }
    const data = getData(task)
    console.log(`Returning data for ${task}`)
    console.log(data)
    res.send(data)
};

module.exports = getTaskData

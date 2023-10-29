const getData = require('../tasks/data/getData')
const getTaskFromToken = require('./getTaskFromToken')
/**
 * @param {*} req
 * @param {Request<P, ResBody, ReqBody, ReqQuery, LocalsObj>|http.ServerResponse} res
 */
const getTaskData = function(req, res){
    const task = getTaskFromToken(req, res)
    if (task.length <= 0) {
        return
    }
    const data = getData(task)
    console.log(`Returning data for ${task}`)
    console.log(data)
    res.send(data)
};

module.exports = getTaskData

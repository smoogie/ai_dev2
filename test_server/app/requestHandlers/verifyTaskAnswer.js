const checkAnswer = require('../tasks/answers/checkAnswer')
const getTaskFromToken = require('./getTaskFromToken')
/**
 * @param {*} req
 * @param {Request<P, ResBody, ReqBody, ReqQuery, LocalsObj>|http.ServerResponse} res
 */
const verifyTaskAnswer = function(req, res){
    const task = getTaskFromToken(req, res)
    if (task.length <= 0) {
        return
    }
    const data = checkAnswer(task, req.body)
    console.log(`Received answer for ${task}`)
    console.log(data)
    res.send(data)
};

module.exports = verifyTaskAnswer

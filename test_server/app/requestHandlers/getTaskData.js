const getData = require('../tasks/data/getData')
const getTaskFromToken = require('./getTaskFromToken')
const {MethodSupportedInTasks, MethodGET} = require('../tasks/methodSupportInTasks')
const {ReturnType, TypeJSON, TypeFile, TypeFunction} = require('../tasks/dataTypeInTasks')
const getFilePath = require('../tasks/data/getFilePath')
const getFunctionData = require('../tasks/data/getFunctionData')
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
    switch (ReturnType[task]) {
        case TypeJSON:
            returnJsonData(task, res)
            break;
        case TypeFile:
            returnFileData(task, res)
            break;
        case TypeFunction:
            returnFunctionData(task, res)
            break;
        default:
            res.status(404).send()
    }
};

function returnJsonData(task, res) {
    const data = getData(task)
    console.log(`Returning data for ${task}`)
    console.log(data)
    res.send(data)
}
function returnFileData(task, res) {
    const path = getFilePath(task)
    const options = {
        root: global.rootPath
    };
    console.log(`Returning data for ${task}`)
    console.log(path)
    res.sendFile(path, options);
}


function returnFunctionData(task, res) {
    const dataFunction = getFunctionData(task)
    const data = dataFunction()
    console.log(`Returning data for ${task}`)
    console.log(data)
    res.send(data)
}

module.exports = getTaskData

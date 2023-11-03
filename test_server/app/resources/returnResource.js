const getFilePath = require("../tasks/data/getFilePath");
const resourcesPaths = require("./resourcesPaths")
const ReturnResource = function(req, res){
    const requestedPath  = req.params.path
    const path = resourcesPaths(requestedPath)
    const options = {
        root: global.rootPath
    };
    console.log(`Returning data for ${requestedPath}`)
    console.log(path)
    res.sendFile(path, options);
}


module.exports = ReturnResource

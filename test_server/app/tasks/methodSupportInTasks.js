const functions = require("./data/json/functionsData");
const MethodGET = "get"
const MethodPOST = "post"
const MethodSupportedInTasks  = {
    "helloapi" : [MethodGET],
    "blogger" : [MethodGET],
    "moderation" : [MethodGET],
    "liar" : [MethodPOST],
    "inprompt":[MethodGET],
    "embedding":[MethodGET],
    "whisper":[MethodGET],
    "functions":[MethodGET],
    "rodo":[MethodGET],
    "scraper":[MethodGET],
    "whoami":[MethodGET],
    "search":[MethodGET],
    "people":[MethodGET],
    "knowledge":[MethodGET],
    "tools":[MethodGET],
    "gnome":[MethodGET],
    "ownapi":[MethodGET],
    "ownapipro":[MethodGET],
    "meme":[MethodGET]
}

module.exports = {
    MethodGET,
    MethodPOST,
    MethodSupportedInTasks
}

const MethodGET = "get"
const MethodPOST = "post"
const MethodSupportedInTasks  = {
    "helloapi" : [MethodGET],
    "blogger" : [MethodGET],
    "moderation" : [MethodGET],
    "liar" : [MethodPOST],
    "inprompt":[MethodGET],
    "embedding":[MethodGET]
}

module.exports = {
    MethodGET,
    MethodPOST,
    MethodSupportedInTasks
}

const buildAnswerResponse = require("../buildAnswerResponse");
const exampleFunctionObject = {
    name: "addUser",
    description: "",
    parameters: {
        type: "object",
        properties: {
            name: {
                type: "string",
                description: ""
            },
            surname: {
                type: "string",
                description: ""
            },
            year: {
                type: "integer",
                description: ""
            }
        }
    }
}
const functionsValidate = function(data) {
    const answer = data.answer
    //TODO: Use AI to run the function like addUser('John','Smith',1974)
    const hasCorrectName = answer.name == exampleFunctionObject.name
    return buildAnswerResponse(hasCorrectName)
}

module.exports = functionsValidate

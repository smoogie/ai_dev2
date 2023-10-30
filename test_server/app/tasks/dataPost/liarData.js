const liarData= function(req) {
    let response = {
        "code" : 0,
        "msg" : "Check field \"answer\" for answer to your question",
        "answer": "Warta is the main river in Poznań."
    }

    const question =  req.body.question
    if (typeof question != "string") {
        response.code = -100
        response.msg = "missing question!"
    } else {
        if  (question.length <= 0) {
            response.code = -100
            response.msg = "missing question!"
        }
    }
    //TODO: randomize if true or false and track for future validation
    //TODO: request info from AI
    return response
}

module.exports = liarData

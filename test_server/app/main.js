const express = require('express')
const app = express()
const port = 3000
const GetTaskToken = require('./requestHandlers/getTaskToken')
const GetTaskData = require('./requestHandlers/getTaskData')
const VerifyTaskAnswer = require('./requestHandlers/verifyTaskAnswer')

global.activeTasks = {}

app.use(express.json());
app.get('/', (req, res) => {
    res.send('Dev AI 2 - test server')
})

app.post('/token/:task', GetTaskToken)
app.get('/task/:token', GetTaskData)
app.post('/answer/:token', VerifyTaskAnswer)

app.listen(port, () => {
    console.log(`Test server listening on port ${port}`)
})

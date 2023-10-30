const express = require('express')
const app = express()
const port = 3000
const GetTaskToken = require('./requestHandlers/getTaskToken')
const GetTaskData = require('./requestHandlers/getTaskData')
const GetTaskDataFromPost = require('./requestHandlers/getTaskDataFromPost')
const VerifyTaskAnswer = require('./requestHandlers/verifyTaskAnswer')
const multer  = require('multer')
const upload = multer()

global.activeTasks = {}

// app.use(express.json());
app.get('/', (req, res) => {
    res.send('Dev AI 2 - test server')
})

app.post('/token/:task', express.json(), GetTaskToken)
app.get('/task/:token', express.json(), GetTaskData)
app.post('/task/:token', upload.none(), GetTaskDataFromPost)
app.post('/answer/:token', express.json(), VerifyTaskAnswer)

app.listen(port, () => {
    console.log(`Test server listening on port ${port}`)
})

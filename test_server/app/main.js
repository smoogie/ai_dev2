const express = require('express')
const app = express()
const port = 3000
const GetTaskToken = require('./requestHandlers/getTaskToken')
const GetTaskData = require('./requestHandlers/getTaskData')
const GetTaskDataFromPost = require('./requestHandlers/getTaskDataFromPost')
const VerifyTaskAnswer = require('./requestHandlers/verifyTaskAnswer')
const ReturnResource = require('./resources/returnResource')
const ReturnScraperFile = require('./tasks/custom/returnScraperFile')
const multer  = require('multer')
const upload = multer()
const path = require('path')


global.activeTasks = {}
global.rootPath = path.join(__dirname)

// app.use(express.json());
app.get('/', (req, res) => {
    res.send('Dev AI 2 - test server')
})

app.post('/token/:task', express.json(), GetTaskToken)
app.get('/task/:token', express.json(), GetTaskData)
app.post('/task/:token', upload.none(), GetTaskDataFromPost)
app.post('/answer/:token', express.json(), VerifyTaskAnswer)
app.get('/resources/:path', express.json(), ReturnResource)

app.get('/scraper/:name', express.json(), ReturnScraperFile)

app.listen(port, () => {
    console.log(`Test server listening on port ${port}`)
})

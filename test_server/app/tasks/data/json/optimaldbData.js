
require('dotenv').config()
const filePath = process.env.HOST + "/resources/3friends.json"

const optimaldbData = {
    code: 0,
    msg: "In a moment you will receive from me a database on three people. It is over 30kb in size. You need to prepare me for an exam in which I will be questioned on this database. Unfortunately, the capacity of my memory is just 9kb. Send me the optimised database",
    database: filePath,
    hint: "I will use GPT-3.5-turbo to answer all test questions"
}

module.exports = optimaldbData

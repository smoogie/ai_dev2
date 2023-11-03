
require('dotenv').config()
const filePath = process.env.HOST + "/resources/whisper.mp3"

const whisperData = {
    code: 0,
    msg: "please return transcription of this file: "+filePath,
    hint: "use WHISPER model - https:\/\/platform.openai.com\/docs\/guides\/speech-to-text"
}

module.exports = whisperData

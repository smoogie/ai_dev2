const resourcesPaths ={
    "whisper.mp3": "./resources/whisper/test.mp3"
}
module.exports = (task) => resourcesPaths[task]

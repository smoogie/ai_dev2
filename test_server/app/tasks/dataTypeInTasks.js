const TypeJSON = "json"
const TypeFile = "file"
const ReturnType = {
    "helloapi" : TypeJSON,
    "blogger" : TypeJSON,
    "moderation" : TypeJSON,
    "liar" : TypeJSON,
    "inprompt":TypeJSON,
    "embedding":TypeJSON,
    // "whisper" :TypeFile
    "whisper" :TypeJSON,
    "functions":TypeJSON
}

module.exports = {
    TypeJSON,
    TypeFile,
    ReturnType
}

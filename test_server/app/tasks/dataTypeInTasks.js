const TypeJSON = "json"
const TypeFile = "file"
const TypeFunction = "function"
const ReturnType = {
    "helloapi" : TypeJSON,
    "blogger" : TypeJSON,
    "moderation" : TypeJSON,
    "liar" : TypeJSON,
    "inprompt":TypeJSON,
    "embedding":TypeJSON,
    // "whisper" :TypeFile
    "whisper" :TypeJSON,
    "functions":TypeJSON,
    "rodo":TypeJSON,
    "scraper":TypeJSON,
    "whoami":TypeFunction,
    "search":TypeJSON,
    "people":TypeJSON
}

module.exports = {
    TypeJSON,
    TypeFile,
    TypeFunction,
    ReturnType
}

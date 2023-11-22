
require('dotenv').config()
const filePath = process.env.HOST + "/resources/monkey.png"

const memeData = {
    code: 0,
    msg: "Create meme using RednerForm API and send me the URL to JPG via \/answer\/ endpoint",
    service: "https:\/\/renderform.io\/",
    // image: filePath,
    image: "https:\/\/zadania.aidevs.pl\/data\/monkey.png",
    text: "Gdy koledzy z pracy mówią, że ta cała automatyzacja to tylko chwilowa moda, a Ty właśnie zastąpiłeś ich jednym, prostym skryptem",
    hint: "https:\/\/zadania.aidevs.pl\/hint\/meme"
}

module.exports = memeData

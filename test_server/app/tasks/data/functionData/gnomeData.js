// require('dotenv').config()
// const filePath = process.env.HOST + "/resources/"
// const images = [
//     "gnome1.png",
//     "gnome2.png",
//     "gnome3.png"
// ]

const images = [
    "https:\/\/zadania.aidevs.pl\/gnome\/3518c516b2dabe06b618dc4fd14d8046.png",
    "https:\/\/zadania.aidevs.pl\/gnome\/788ef0905682f7185c8944cbbaed081d.png",
    "https:\/\/zadania.aidevs.pl\/gnome\/ca1c9f3f65e8c854ab0c5bafd6a96746.png"
]

const gnomeData = function () {
    const image = Math.floor(Math.random() * 3);
    // const url = filePath+images[image]
    const url = images[1]
    //TODO: remember question asked on request for future validation
    return {
        code: 0,
        msg: "I will give you a drawing of a gnome with a hat on his head. Tell me what is the color of the hat in POLISH. If any errors occur, return \"ERROR\" as answer",
        hint: "it won't always be a drawing of a gnome >:)",
        url
    }
}

module.exports = gnomeData

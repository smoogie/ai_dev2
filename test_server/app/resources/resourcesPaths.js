const resourcesPaths ={
    "whisper.mp3": "./resources/whisper/test.mp3",
    "gnome1.png":"./resources/gnome/gnome1.png",
    "gnome2.png":"./resources/gnome/gnome2.png",
    "gnome3.png":"./resources/gnome/gnome3.png"
}
module.exports = (task) => resourcesPaths[task]

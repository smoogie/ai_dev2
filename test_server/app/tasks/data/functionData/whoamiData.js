const hints = [
    "został adoptowany tuż po urodzeniu",
    "był wegetarianinem i eksperymentował z różnymi ekstremalnymi dietami",
    "został wyrzucony z własnej firmy 1985 roku",
    "Jego oficjalna pensja wynosiła 1 dolar rocznie",
    "uczęszczał do college, ale rzucił szkołę po jednym semestrze",
    "W momencie śmierci, jego majątek był szacowany na około 10,2 miliarda dolarów",
    "był wielkim fanem grupy The Beatles",
    "pracował jako technik w firmie Atari"
]

const whoamiData = function () {
    const hint = Math.floor(Math.random() * 8);
    return {
        code: 0,
        msg: "Each time you call up this task, I will return a trivia item about a certain person (the person does not change). Guess who I am",
        hint: hints[hint]
    }
}

module.exports = whoamiData

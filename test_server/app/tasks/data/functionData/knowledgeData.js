const questions = [
    "kto napisał Romeo i Julię?",
    "ile orientacyjnie ludzi mieszka w Polsce?",
    "jaki jest teraz kurs dolara?",
    "podaj populację Francji",
    "podaj aktualny kurs EURO"
]

const knowledgeData = function () {
    const question = Math.floor(Math.random() * 5);
    //TODO: remember question asked on request for future validation
    return {
        code: 0,
        msg: "I will ask you a question about the exchange rate, the current population or general knowledge. Decide whether you will take your knowledge from external sources or from the knowledge of the model",
        question: questions[1],
        "database #1": "Currency http:\/\/api.nbp.pl\/en.html (use table A)",
        "database #2": "Knowledge about countries https:\/\/restcountries.com\/ - field 'population'"
    }
}

module.exports = knowledgeData

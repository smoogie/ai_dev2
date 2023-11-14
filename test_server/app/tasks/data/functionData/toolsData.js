const questions = [
    "Przypomnij mi, abym zapisał się na AI Devs 3.0",
    "W poniedziałek mam wizytę u lekarza",
    "Pojutrze mam kupić 1kg ziemniaków"
]

const toolsData = function () {
    const question = Math.floor(Math.random() * 3);
    //TODO: remember question asked on request for future validation
    return {
        code: 0,
        msg: "Decide whether the task should be added to the ToDo list or to the calendar (if time is provided) and return the corresponding JSON",
        question: questions[question],
        hint: "always use YYYY-MM-DD format for dates",
        "example for ToDo": "Przypomnij mi, że mam kupić mleko = {\"tool\":\"ToDo\",\"desc\":\"Kup mleko\" }",
        "example for Calendar": "Jutro mam spotkanie z Marianem = {\"tool\":\"Calendar\",\"desc\":\"Spotkanie z Marianem\",\"date\":\"2023-11-15\"}"
    }
}

module.exports = toolsData

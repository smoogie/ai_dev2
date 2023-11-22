const googleData = {
    code: 0,
    msg: "Provide me the URL to your API (HTTPS) via \/answer\/ endpoint. I will ask your API a question that requires search engine integration. Your job is to provide me answer to my question",
    hint1: "Please use SerpAPI or similar service. https:\/\/serpapi.com (free account is enough)",
    hint2: "Probably I will ask more than one question, so be prepared for that",
    hint3: "Please return the answer in JSON format, with \"reply\" field!",
    hint4: "Return as concise an answer as possible"
}

module.exports = googleData

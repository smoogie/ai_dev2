const ReturnScraperFile = function(req, res){
    const action = Math.floor(Math.random() * 4);
    switch (action) {
        case 0:
            res.status(500).send()
            break;
        case 1:
            setTimeout(returnFile, 2000,res);
            break;
        case 2:
            res.status(404).send()
            break;
        case 3:
            returnFile(res)
            break;
    }
}

function returnFile(res) {
    const path = "./resources/scraper/text_pizza_history.txt"
    const options = {
        root: global.rootPath
    };
    console.log(path)
    res.sendFile(path, options);

}


module.exports = ReturnScraperFile

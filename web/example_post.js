
// Example of posting a score to the highscores in javascript.

function postScore(address, name, score) {
    fetch(address, {
        method: "POST",
        mode: "cors",
        body: JSON.stringify({name, score}),
        headers: { 'Content-Type': 'application/json;charset=UTF-8' }
    })
    .then(r => {
        console.log("Request complete! response:", r);
    });
}



// Example of retrieving data and calling a processDataFunction(json_data)

function getBoardData(address, processDataFunction) {
    fetch(address, {
        method: "GET",
    })
    .then(response => {
        return response.json();
    })
    .then(myJson => {
        console.log(JSON.stringify(myJson));
        if (processDataFunction !== undefined) {
            processDataFunction(myJson);
        }
    });
}

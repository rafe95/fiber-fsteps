const API_URL = "/todo"

const INPUT = document.querySelector("#taskInput")
const LIST = document.querySelector("#taskList")

function sendToDo() {
    let task = INPUT.value
    let toDo = {id: "", task}
    fetch(API_URL, {
        method: 'POST',
        body: JSON.stringify(toDo),
        headers: {
            'Content-Type': 'application/json'
        }
    }).then(res => res.json())
        .catch(error => console.error('Error:', error))
        .then(response => {
                console.log('Success:', response);
                LIST.innerHTML += `<li>${response.task}</li>`
                INPUT.value = ""
            }
        )
}

function getAll(){
    fetch("/todos").then( response => response.json()).then(
        data => {
            Array.from(data).forEach(
                task => LIST.innerHTML += `<li>${task.task}</li>`
            )
        }
    )
}
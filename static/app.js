const API_URL = "/todo"

const INPUT = document.querySelector("#taskInput")
const CATEGORY = document.querySelector("#categoryInput")
const LIST = document.querySelector("#taskList")

function sendToDo() {
    let toDo = {id: "", task: INPUT.value, completed: false, category: CATEGORY.value}
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
                addTodo(response)
                clearInputs()
            }
        )
}

function addTodo(toDo) {
    LIST.innerHTML += `
            <li class="list-group-item d-flex justify-content-between align-items-center" id>
                <div class="form-check">
                    <input class="form-check-input" type="checkbox" value="${toDo.completed}">
                </div>
            ${toDo.task}
            <span class="badge bg-primary rounded-pill">${toDo.category}</span>
            </li>`
}

function clearInputs() {
    INPUT.value = ""
    CATEGORY.value = ""
}

function getAll() {
    fetch("/todos").then(response => response.json()).then(
        data => {
            Array.from(data).forEach(
                task => addTodo(task)
            )
        }
    )
}
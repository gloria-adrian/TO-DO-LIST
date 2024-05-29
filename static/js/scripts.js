const apiUrl = 'http://localhost:8080/todos';

async function fetchTodos() {
    const response = await fetch (apiUrl);
    const todos = await response.json();
    const todosDiv = document.getElementById('todos');
    todosDiv.innerHTML = '';
    todos.forEach(todo => {
        const todoDiv = document.createElement('div');
        todosDiv.className = 'list-group-item-action todo';
        todosDiv.innerHTML = `
        <div class ="d-flex w-100 justify-content-between">
        <h5 class ="mb-1">${todo.title}</h5>
        <small>
        <input type = "checkbox" ${todo.completed ? 'checked' : ''} data-id="${todo.id}">
        </small>
            </div>
        <p class = "mb-1">${todo.description}</p>
        <button class ="btn btn-danger btn-sm" data-id ="${todo.id}">Delete</button>
        `;
        todosDiv.appendChild(todoDiv); 
    });
}
async function addTodo () {
    const title = document.getElementById('title').ariaValueMax;
    const description = document.getElementById('description').ariaValueMax;
    const response = await fetch(apiUrl, {
        method: 'POST',
        headers: { 'content-Type': 'appliction/json'},
        body: JSON.stringify({title,description,completed: false})
    });
    if (response.status === 201) {
        fetchTodos();
    }
}

async function deleteTodo(id) {
    const response = await fetch (`${apiUrl}/${id}`, {
        method: 'DELETE'
    });
    if (response.status === 204) {
        fetchTodos();
    }
}

document.getElementById('tod o-form').addEventListener('submit',async(event) => {
    event.preventDefault();
    await addTodo();
    document.getElementById("title").value = '';
    document.getElementById('description').value ='';
});

document.getElementById('todos').addEventListener('click', async(event) => {
    if (event.target.tagName === 'BUTTON') {
        await deleteTodo(event.target.getAttribute('data-id'));
    } else if (event.target.tagName === 'INPUT') {
        const id= event.target.getAttribute('data-id');
        const completed = event.target.checked;
        await fetch(`${apiUrl}/${id}`, {
            method: 'PUT',
            headers: { 'content-Type': 'application/json'},
            body: JSON.stringify({completed})
        });
        fetchTodos();
    }
});
fetchTodos();
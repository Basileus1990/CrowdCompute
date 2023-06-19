let taskList = document.getElementById('task-list');


document.getElementById('task-getter').addEventListener('click', () => {
    // clear task list
    taskList.innerHTML = '';
    fetch('/get-tasks').then((response) => {
        response.json().then((data) => {
            //console.log(data);
            data.forEach((task) => {
                let taskItem = document.createElement('li');

                let taskName = document.createElement('p');
                taskName.textContent = task.name;
                taskItem.appendChild(taskName);

                let taskAuthor = document.createElement('p');
                taskAuthor.textContent = "Author: " + task.author;
                taskItem.appendChild(taskAuthor);

                let taskDescription = document.createElement('p');
                taskDescription.textContent = "Description: " + task.description;
                taskItem.appendChild(taskDescription);

                let taskCode = document.createElement('p');
                taskCode.textContent = "Code: " + task.code;
                taskItem.appendChild(taskCode);

                let executeButton = document.createElement('button');
                executeButton.textContent = "Execute the task";
                executeButton.addEventListener('click', () => {
                    if (task.hasOwnProperty('data') == false) {
                        task.data = "";
                    }
                    let taskResult = executeTask(task.code, task.data);
                    if (taskResult.error) {
                        alert("Error: " + taskResult.error);
                        return;
                    }
                    taskResult = JSON.stringify(taskResult);
                    alert("Task result: " + taskResult);
                });
                taskItem.appendChild(executeButton);

                taskList.appendChild(taskItem);
            });
        });
    });

});

// returns an object with the task result
function executeTask(taskCode, taskData) {
    if (taskData != "") {
        taskData = JSON.parse(taskData);
    }
    let taskResult = {};
    try {
        eval(taskCode);
    } catch (e) {
        console.log(e);
        taskResult.error = e;
    }

    return taskResult;
}
let taskList = document.getElementById('task-list');


document.getElementById('task-getter').addEventListener('click', () => {
    // clear task list
    taskList.innerHTML = '';
    fetch('/get-tasks').then((response) => {
        response.json().then((data) => {
            //console.log(data);
            data.forEach((task) => {
                let taskItem = document.createElement('li');
                taskItem.innerHTML = task.code;
                taskList.appendChild(taskItem);
            });
        });
    });

});
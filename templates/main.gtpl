<html>
<head>
    <title>Task list</title>
</head>
<body>
<h1>Task list:</h1>
    <div>
        <ul>
        {{ range $task := .}}
            <li><input type="checkbox" id="{{ $task.Id }}" value="{{ $task.IsFinished }}" /><label for="{{ $task.Id }}">{{ $task.Task }} </label></li>
        {{ end }}
        </ul>

    </div>
    <form action="/add" method="POST">
    <input type="submit" value="Add new task">
</body>
</html>
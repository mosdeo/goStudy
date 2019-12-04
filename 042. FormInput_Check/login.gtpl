<html>
<head>
<title></title>
</head>
<body>
<form action="http://localhost:9090/login" method="post">
    使用者名稱:<input type="text" name="username">
    中文實名:<input type="text" name="realname">
    密碼:<input type="password" name="password">
    年紀:<input type="text" name="age">
    <input type="submit" value="登入">

    </br>
    下拉式功能表
    <select name="fruit">
    <option value="apple">apple<option>
    <option value="pear">pear<option>
    <option value="banana">banana<option>
    </select>

    </br>
    必須選項按鈕
    <input type="radio" name="gender" value="M">男
    <input type="radio" name="gender" value="F">女

    </br>
    核取按鈕
    <input type="checkbox" name="interest" value="interest0">興趣0
    <input type="checkbox" name="interest" value="interest1">興趣1
    <input type="checkbox" name="interest" value="interest2">興趣2
</form>
</body>
</html>

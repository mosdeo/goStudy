<html>
<head>
<title></title>
</head>
<body>
<form action="http://localhost:9090/read" method="post">
    username:<input type="text" name="username">
    age大於:<input type="text" name="ageLower">
    age小於:<input type="text" name="ageUpper">
    必須選項按鈕
    <input type="radio" name="gender" value="M">男
    <input type="radio" name="gender" value="F">女
    <input type="submit" value="查詢">
</form>
</body>
</html>

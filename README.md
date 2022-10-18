<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    
</head>
<body>
<h1 align="center">Hi 👋, </h1>
<h3 align="center">API CAKE GOLANG</h3>
<div>
    <h3>Api Golang:</h3>
</div>
<div>
    <h3>Api Host</h3>
    <p>http://localhost:5000/api/v1</p>
</div>
<div>
		cakeRoutes.GET("/cake", cakeController.All)
		cakeRoutes.POST("/cake", cakeController.Insert)
		cakeRoutes.PATCH("/cake/:id", cakeController.Update)
		cakeRoutes.GET("/cake/:id", cakeController.FindCakeById)
		cakeRoutes.DELETE("/cake/:id", cakeController.Delete)
    <h2>GET ALL CAKE REQUUST</h2>
    <p>http://localhost:5000/api/v1/cake</p>

    <h2>GET DETAIL CAKE REQUEST</H2>
 <p>http://localhost:5000/api/v1/cake/:id</p>

    <h2>CREATE CAKE DATA</H2>
    <p>POST</p>
 <p>http://localhost:5000/api/v1/cake</p>
    <H6>body</H6>
        <p>use form data</p>
        <ul>
            <li>title</li>
            <li>description</li>
            <li>imageFile (to upload file)</li>
            <li>rating</li>
        </ul>

  <h2>EDIT CAKE DATA</H2>
    <p>POST</p>
 <p>http://localhost:5000/api/v1/cake/:id</p>
    <H6>body</H6>
        <p>use form data</p>
        <ul>
            <li>title</li>
            <li>description</li>
            <li>imageFile (to upload file)</li>
            <li>rating</li>
        </ul>

  <h2>DELETE CAKE DATA</H2>
    <p>POST</p>
 <p>http://localhost:5000/api/v1/cake/:id</p>
 
</div>
</body>
</html>
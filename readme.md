## api

```
<!--获取所有-->
router.GET("/api/users", Users)
<!--创建一个-->
router.POST("/api/users", UserCreate)
<!--获取一个用户信息-->
router.GET("/api/users/:id", UserOne)
<!--更新用户信息-->
router.PUT("/api/users/:id", UserUpdate)
<!--删除用户-->
router.POST("/api/users/:id", UserDel)
```
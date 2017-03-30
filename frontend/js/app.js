function Http(obj, cb) {
    $.ajax({
        type: obj.type || "GET",
        url: obj.url,
        dataType: obj.dataType || "json",
        data: obj.data || "",
        success: function (data) {
            cb(data)
        },
        error: function () {
            alert('fail');
        }
    });
}

// 用户列表
function List() {
    let obj = {
        url: "/api/users"
    }

    Http(obj, function (data) {
        console.log(data)
    })
}
// 创建用户
function Create(user = {}) {

    user = {
        name: user.name || "notName",
        age: user.age || 18
    }

    let obj = {
        url: "/api/users",
        type: "POST",
        data: user
    }
    Http(obj, function (data) {
        console.log(data)
    })
}

// 更新用户信息
function Update(user = {}) {

    if (user.id) {
        alert("not id")
        return
    }

    user = {
        id: user.id,
        name: user.name || "update notName",
        age: user.age || 18
    }

    let obj = {
        url: "/api/users/" + user.id,
        type: "PUT",
        data: user
    }
    Http(obj, function (data) {
        console.log(data)
    })
}


// 显示一个用户信息
function Show(id = 1) {
    let obj = {
        url: "/api/users/" + id,
        type: "GET",
    }
    Http(obj, function (data) {
        console.log(data)
    })
}


// 删除一个用户
function Del(id = 1) {
    let obj = {
        url: "/api/users/" + id,
        type: "POST",
    }
    Http(obj, function (data) {
        console.log(data)
    })
}
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
        ReactView(data)
        if(!data.Code){
            ListView(data.Data)

        }
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
        ReactView(data)
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
        ReactView(data)
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
        ReactView(data)
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
        ReactView(data)
    })
}

// 渲染试图
function ReactView(data) {
    let str = JSON.stringify(data)

    console.log(str)
    let app = document.getElementById("app")
    app.innerHTML = str
}

// 渲染 List
function ListView(data) {
    let len = data.length

    let str = ''
    for (var i = 0; i < len; i++) {
        let obj = data[i]
        str += `
        <div>
            <span>Name: ${obj.Name} </span>
            <p>Age: ${obj.Age} </p>
        </div>
        <div>
            <button type="button" onclick="Show(${obj.Id})">详细</button>
            <button type="button" onclick="Del(${obj.Id})">删除</button>
        <div>
        <hr />
        `
    }

    let list = document.getElementById("list")
    list.innerHTML = str
}
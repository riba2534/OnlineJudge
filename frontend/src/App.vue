<template>
  <div id="app">
    <el-menu
      :default-active="activeIndex"
      mode="horizontal"
      @select="handleSelect"
      v-bind:router="true"
      id="nav"
    >
      <el-menu-item index="/" id="title">OnlineJudge</el-menu-item>
      <el-menu-item index="/main">
        <i class="el-icon-s-home"></i>主页
      </el-menu-item>
      <el-menu-item index="/problem">
        <i class="el-icon-grape"></i>问题
      </el-menu-item>
      <el-menu-item index="/status">
        <i class="el-icon-ice-tea"></i>状态
      </el-menu-item>
      <el-menu-item index="/contest">
        <i class="el-icon-bicycle"></i>竞赛
      </el-menu-item>
      <el-menu-item index="/rank">
        <i class="el-icon-medal"></i>排名
      </el-menu-item>
      <el-menu-item index="/about">
        <i class="el-icon-lollipop"></i>关于
      </el-menu-item>
      <el-button
        round
        id="button"
        @click="dialogRegisterVisible = true"
        v-show="!loginshow"
      >Register</el-button>
      <el-button round id="button" @click="dialogLoginVisible = true" v-show="!loginshow">Login</el-button>

      <el-dropdown
        id="user"
        v-show="loginshow"
        @command="handleCommand"
        :show-timeout="100"
        :split-button="true"
        @visible-change="updatename"
      >
        <span class="el-dropdown-link">欢迎 {{name}}</span>
        <el-dropdown-menu slot="dropdown">
          <el-dropdown-item command="home">主页</el-dropdown-item>
          <el-dropdown-item command="submittion">提交记录</el-dropdown-item>
          <el-dropdown-item command="setting">设置</el-dropdown-item>
          <el-dropdown-item command="admin" divided v-show="isadmin">后台管理</el-dropdown-item>
          <el-dropdown-item command="logout" divided>登出</el-dropdown-item>
        </el-dropdown-menu>
      </el-dropdown>
    </el-menu>

    <el-dialog title="注册" :visible.sync="dialogRegisterVisible">
      <el-form :model="form">
        <el-row :gutter="10">
          <el-col :span="3">
            <div style="text-align:center;margin:5px;">用户名</div>
          </el-col>
          <el-col :span="12">
            <el-input v-model="form.username" autocomplete="off"></el-input>
          </el-col>
        </el-row>
        <el-row :gutter="10">
          <el-col :span="3">
            <div style="text-align:center;margin:5px;">昵称</div>
          </el-col>
          <el-col :span="12">
            <el-input v-model="form.name" autocomplete="off"></el-input>
          </el-col>
        </el-row>
        <el-row :gutter="10">
          <el-col :span="3">
            <div style="text-align:center;margin:5px;">密码</div>
          </el-col>
          <el-col :span="12">
            <el-input type="password" v-model="form.password" autocomplete="off"></el-input>
          </el-col>
        </el-row>
        <el-row :gutter="10">
          <el-col :span="3">
            <div style="text-align:center;margin:5px;">确认密码</div>
          </el-col>
          <el-col :span="12">
            <el-input type="password" v-model="form.comfirm" autocomplete="off"></el-input>
          </el-col>
        </el-row>
        <el-row :gutter="10">
          <el-col :span="3">
            <div style="text-align:center;margin:5px;">学校</div>
          </el-col>
          <el-col :span="12">
            <el-input v-model="form.school" autocomplete="off"></el-input>
          </el-col>
        </el-row>
        <el-row :gutter="10">
          <el-col :span="3">
            <div style="text-align:center;margin:5px;">专业</div>
          </el-col>
          <el-col :span="12">
            <el-input v-model="form.course" autocomplete="off"></el-input>
          </el-col>
        </el-row>
        <el-row :gutter="10">
          <el-col :span="3">
            <div style="text-align:center;margin:5px;">班级</div>
          </el-col>
          <el-col :span="12">
            <el-input v-model="form.classes" autocomplete="off"></el-input>
          </el-col>
        </el-row>
        <el-row :gutter="10">
          <el-col :span="3">
            <div style="text-align:center;margin:5px;">学号</div>
          </el-col>
          <el-col :span="12">
            <el-input v-model="form.number" autocomplete="off"></el-input>
          </el-col>
        </el-row>
        <el-row :gutter="10">
          <el-col :span="3">
            <div style="text-align:center;margin:5px;">真实姓名</div>
          </el-col>
          <el-col :span="12">
            <el-input v-model="form.realname" autocomplete="off"></el-input>
          </el-col>
        </el-row>
        <el-row :gutter="10">
          <el-col :span="3">
            <div style="text-align:center;margin:5px;">QQ</div>
          </el-col>
          <el-col :span="12">
            <el-input v-model="form.qq" autocomplete="off"></el-input>
          </el-col>
        </el-row>
        <el-row :gutter="10">
          <el-col :span="3">
            <div style="text-align:center;margin:5px;">Email</div>
          </el-col>
          <el-col :span="12">
            <el-input v-model="form.email" autocomplete="off"></el-input>
          </el-col>
        </el-row>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="dialogRegisterVisible = false">取 消</el-button>
        <el-button type="primary" @click="registerClick">确 定</el-button>
      </div>
    </el-dialog>

    <el-dialog title="登录" :visible.sync="dialogLoginVisible">
      <el-form :model="form">
        <el-row :gutter="10">
          <el-col :span="3">
            <div style="text-align:center;margin:5px;">用户名</div>
          </el-col>
          <el-col :span="12">
            <el-input v-model="form.username" autocomplete="off"></el-input>
          </el-col>
        </el-row>
        <el-row :gutter="10">
          <el-col :span="3">
            <div style="text-align:center;margin:5px;">密码</div>
          </el-col>
          <el-col :span="12">
            <el-input type="password" v-model="form.password" autocomplete="off"></el-input>
          </el-col>
        </el-row>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="dialogLoginVisible = false">取 消</el-button>
        <el-button type="primary" @click="loginClick">确 定</el-button>
      </div>
    </el-dialog>

    <transition name="el-zoom-in-center" mode="out-in">
      <router-view id="route"></router-view>
    </transition>
  </div>
</template>

<script>
export default {
  name: "App",

  data() {
    return {
      activeIndex: "1",
      dialogLoginVisible: false,
      dialogRegisterVisible: false,
      loginshow: sessionStorage.username,
      username: sessionStorage.username,
      name: sessionStorage.name,
      isadmin: false,
      form: {
        username: "",
        password: "",
        comfirm: "",
        name: "",
        school: "",
        course: "",
        classes: "",
        number: "",
        realname: "",
        qq: "",
        email: ""
      }
    };
  },
  created() {
    this.isadmin = sessionStorage.type == 2 || sessionStorage.type == 3;
  },
  methods: {
    updatename() {
      this.name = sessionStorage.name;
    },
    // 选择tab触发的事件
    handleSelect(key, keyPath) {
      console.log(key, keyPath);
    },
    handleCommand(command) {
      if (command == "logout") {
        this.$axios
          .get("http://" + this.$ip + ":" + this.$port + "/logout/")
          .then(response => {
            this.$message({
              message: "登出成功！",
              type: "success"
            });
            sessionStorage.setItem("username", "");
            sessionStorage.setItem("name", "");
            sessionStorage.setItem("auth", "");
            sessionStorage.setItem("rating", "");
            this.loginshow = 0;
            this.username = "";
            this.$router.go(0);
            console.log(response);
          })
          .catch(error => {
            this.$message.error("服务器错误！" + "(" + error + ")");
          });
      }
      if (command == "home") {
        this.$router.push({
          name: "user",
          query: { username: sessionStorage.username }
        });
      }
      if (command == "setting") {
        this.$router.push({
          name: "setting",
          params: { username: sessionStorage.username }
        });
      }
      if (command == "submittion") {
        this.$router.push({
          name: "statue",
          query: { username: sessionStorage.username }
        });
      }
      if (command == "admin") {
        this.$router.push({
          name: "admin"
        });
      }
    },
    registerClick() {
      if (
        !this.form.name ||
        !this.form.school ||
        !this.form.course ||
        !this.form.classes ||
        !this.form.number ||
        !this.form.realname ||
        !this.form.qq ||
        !this.form.email
      ) {
        this.$message.error("字段不能为空！");
        return;
      }
      if (this.form.password != this.form.comfirm) {
        this.$message.error("两次密码不一致！");
        return;
      }
      if (this.form.username.length < 6) {
        this.$message.error("用户名太短！");
        return;
      }
      if (this.form.name.length < 2) {
        this.$message.error("昵称太短！");
        return;
      }
      if (this.form.password.length < 6) {
        this.$message.error("密码太短！");
        return;
      }

      // 对密码进行md5加密
      this.form.password = this.$md5(this.form.password);

      this.$axios
        .post("http://" + this.$ip + ":" + this.$port + "/userdata/", this.form)
        .then(response => {
          this.$axios
            .post(
              "http://" + this.$ip + ":" + this.$port + "/register/",
              this.form
            )
            .then(response => {
              if (response.data == "usererror") {
                this.$message.error("用户名已存在！");
                return;
              }
              this.$message({
                message: "注册成功！",
                type: "success"
              });
              this.dialogRegisterVisible = false;
            })
            .catch(error => {
              this.$message.error("服务器错误！" + "(" + error + ")");
            });
          console.log(response);
        })
        .catch(error => {
          this.$message.error("服务器错误！" + "(" + error + ")");
        });
    },
    loginClick() {
      this.form.password = this.$md5(this.form.password);
      this.$axios
        .post("http://" + this.$ip + ":" + this.$port + "/login/", this.form)
        .then(response => {
          if (response.data == "passworderror") {
            this.$message.error("密码错误");
            return;
          }
          this.$message({
            message: "登录成功！",
            type: "success"
          });

          sessionStorage.setItem("username", this.form.username);
          sessionStorage.setItem("name", response.data.name);
          sessionStorage.setItem("type", response.data.type);
          if (response.data.type == 2 || response.data.type == 3)
            this.isadmin = true;

          this.dialogRegisterVisible = false;
          this.dialogLoginVisible = false;
          this.loginshow = 1;
          this.username = this.form.username;
          this.name = sessionStorage.name;

          this.$axios
            .get(
              "http://" +
                this.$ip +
                ":" +
                this.$port +
                "/userdata/?username=" +
                this.username
            )
            .then(response => {
              sessionStorage.setItem("rating", response.data[0].rating);
              this.$router.go(0);
            });
        })
        .catch(error => {
          this.$message.error("内部错误：" + error);
        });
    }
  }
};
</script>


<style scope>
.el-dropdown-link {
  cursor: pointer;
  color: #409eff;
}
#button {
  float: right;
  margin: 10px;
}
#user {
  float: right;
  margin: 10px;
}
#nav {
  background-color: #ffffff;
  position: fixed;
  left: 0px;
  top: 0px;
  z-index: 5;
  width: 100%;
  box-shadow: 00px 0px 00px rgb(255, 255, 255),
    /*左边阴影*/ 0px 0px 10px rgb(255, 255, 255),
    /*上边阴影*/ 0px 0px 0px rgb(255, 255, 255),
    /*右边阴影*/ 1px 1px 0px rgb(218, 218, 218); /*下边阴影*/
}
#route {
  position: relative;
  top: 80px;
}
#title {
  font-size: 20px;
  font-weight: bold;
}
</style>

<template>
  <div id="app">
    <el-menu
      :default-active="activeIndex"
      mode="horizontal"
      @select="handleSelect"
      v-bind:router="true"
      id="nav"
    >
      <el-menu-item index="/" id="title">NYOJ</el-menu-item>
      <el-menu-item index="/main">
        <i class="el-icon-star-off"></i>主页
      </el-menu-item>
      <el-menu-item index="/problem">
        <i class="el-icon-menu"></i>问题
      </el-menu-item>
      <el-menu-item index="/statue">
        <i class="el-icon-tickets"></i>状态
      </el-menu-item>
      <el-menu-item index="/contest">
        <i class="el-icon-bell"></i>竞赛
      </el-menu-item>
      <el-menu-item index="/rank">
        <i class="el-icon-star-on"></i>排名
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
        <span class="el-dropdown-link">Welcome {{name}}</span>
        <el-dropdown-menu slot="dropdown">
          <el-dropdown-item command="home">Home</el-dropdown-item>
          <el-dropdown-item command="submittion">Submittion</el-dropdown-item>
          <el-dropdown-item command="setting">Setting</el-dropdown-item>
          <el-dropdown-item command="admin" divided v-show="isadmin">Admin</el-dropdown-item>
          <el-dropdown-item command="logout" divided>Logout</el-dropdown-item>
        </el-dropdown-menu>
      </el-dropdown>
    </el-menu>

    <router-view></router-view>
  </div>
</template>

<script>
// import HelloWorld from "./components/HelloWorld.vue";

export default {
  name: "app",
  components: {}
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

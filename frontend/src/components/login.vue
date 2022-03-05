<template>
  <div class="hello">
    <!-- <h1>{{ msg }}</h1> -->

    <div>
      <h1 class="title">Login</h1>
    </div>
    <hr>

    <div>
      <b-container fluid>
        <b-row class="text-center">
          <b-col cols="4" class="text-center"></b-col>
          <b-col cols="4" class="text-center">
            <form>

              <div class="field">
                <label class="label">用户名</label>
                <b-form-input name="userName" v-model="userName" v-validate="'required'" class="input"
                              type="text" placeholder="输入用户名"></b-form-input>
              </div>

              <div class="field" style="margin-top:13px;">
                <label class="label">密码</label>
                <b-form-input type="password" name="passWord" v-model="passWord" class="input"
                              v-validate="'required|digits'" placeholder="输入密码"></b-form-input>
              </div>
            </form>
          </b-col>
        </b-row>
        <hr>
        <b-row class="text-center">
          <b-col cols="4" class="text-center"></b-col>
          <b-col cols="4" class="text-center">
            <div><label class="label">{{ isPass }}</label></div>
          </b-col>
        </b-row>
      </b-container>
    </div>

    &nbsp;
    <div>
      <b-button variant="primary" v-on:click="signIn()">登陆</b-button>&nbsp;&nbsp;
      <b-button variant="primary" v-on:click="signUp()">注册</b-button>
    </div>
  </div>
</template>

<script>

import axios from 'axios';
import Vue from 'vue'
import VeeValidate from 'vee-validate'

/* eslint-disable */
Vue.use(VeeValidate)

export default {
  name: 'Login',

  data: function () {
    return {
      userName: "", passWord: "", isPass: ""
    }
  },

  methods: {
    signIn: function () {
      var data = {"user_name": this.userName, "pass_word": this.passWord}

      /*eslint-disable*/
      console.log(data)
      /*eslint-enable*/

      axios({
        method: "POST",
        url: "http://127.0.0.1:8090/signin",
        data: data,
        headers: {"content-type": "text/plain"}
      }).then(result => {
        // this.response = result.data;
        this.isPass = result.data['is_pass']

        /*eslint-disable*/
        console.log(result.data)
        /*eslint-enable*/

      }).catch(error => {
        /*eslint-disable*/
        console.error(error);
        /*eslint-enable*/
      });
    },
    signUp: function () {
      var data = {"user_name": this.userName, "pass_word": this.passWord}

      /*eslint-disable*/
      console.log(data)
      /*eslint-enable*/

      axios({
        method: "POST",
        url: "http://127.0.0.1:8090/signup",
        data: data,
        headers: {"content-type": "text/plain"}
      }).then(result => {
        // this.response = result.data;
        this.isPass = result.data['is_pass']

        /*eslint-disable*/
        console.log(result.data)
        /*eslint-enable*/

      }).catch(error => {
        /*eslint-disable*/
        console.error(error);
        /*eslint-enable*/
      });
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h3 {
  margin: 40px 0 0;
}

ul {
  list-style-type: none;
  padding: 0;
}

li {
  display: inline-block;
  margin: 0 10px;
}

a {
  color: #42b983;
}
</style>
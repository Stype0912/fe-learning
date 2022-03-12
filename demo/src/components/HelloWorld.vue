<template>
  <!--  <div id="counter">-->
  <!--    Counter: {{ counter }}-->
  <!--  </div>-->
  <div id="bind-attribute">
  <span v-bind:title="message">
    鼠标悬停几秒钟查看此处动态绑定的提示信息！
  </span>
    <p>{{ message }}</p>
    <button v-on:click="reverseMessage">反转 Message</button>
  </div>

</template>

<script>
export default {
  data() {
    return {
      message: 'You loaded this page on ' + new Date().toLocaleString()
    }
  },
  methods: {
    reverseMessage() {
      this.message = this.message
        .split('')
        .reverse()
        .join('')
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h1, h2 {
  font-weight: normal;
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
button {
  
}
</style>

<template>
  <div class="hello">
    <h1>{{ msg }}</h1>
    <p>
      If you don't understand what this site is for,<br>
      check out the
      <a
        href="https://github.com/dave/jennifer"
        target="_blank"
        rel="noopener"
      >dave/jennifer repo</a> for code generation.
      Here's our <a
        href="https://github.com/rwlist/go2jen"
        target="_blank"
        rel="noopener"
      >GitHub</a>.
    </p>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue'

export default defineComponent({
  name: 'HelloWorld',
  props: {
    msg: String
  }
})
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

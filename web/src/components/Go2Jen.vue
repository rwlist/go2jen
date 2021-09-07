<template>
  <div style="width: 100%;">
    <HelloWorld msg="Go2Jen" />
    <div style="width: 100%; overflow:auto;">
      <textarea
        onkeydown="if(event.keyCode===9){var v=this.value,s=this.selectionStart,e=this.selectionEnd;this.value=v.substring(0, s)+'\t'+v.substring(e);this.selectionStart=this.selectionEnd=s+1;return false;}"
        id="input"
        v-model="code"
        class="code"
        style="float:left; width: 49%;"
      />
      <div
        id="output"
        class="code"
        style="float:right; width: 50%;"
      >
        <pre style="float: left;"><code>{{ jenCode }}</code></pre>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import HelloWorld from './HelloWorld.vue'

export default defineComponent({
  name: 'Go2Jen',
  components: { HelloWorld },
  data () {
    return {
      code: 'package a\n\nfunc hello(arg string) {\n\t// enter code here\n\tfmt.Println("Hello", arg)\t\t\n}' as string,
      jenCode: 'jen codegen will appear here' as string
    }
  },
  watch: {
    code (newCode: string, oldCode: string) {
      console.log(newCode, '|', oldCode)
      this.doConvert(newCode)
    }
  },
  methods: {
    async doConvert (code: string) {
      const response = await fetch('https://functions.yandexcloud.net/d4eibao4k3sldaho1opd', {
        method: 'POST',
        headers: {
          'content-type': 'application/json;charset=UTF-8'
        },
        body: JSON.stringify({
          Code: code
        })
      })

      const data = await response.json()
      if (response.ok) {
        console.log(data)
        this.jenCode = data.JenniferGen
      } else {
        console.log(response.status)
      }
    }
  }
})
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
#input {
    background: #ECF2F5;
    border-right: 2px solid #AAA;
}
#output {
    background: #E0EBF5;
}
.code {
    height: 500px;
    font: 14px/1.5em 'Consolas', 'Menlo', 'Monaco', 'Courier New', monospace;
    overflow: scroll;
}
</style>

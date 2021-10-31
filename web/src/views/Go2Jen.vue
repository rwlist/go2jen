<template>
  <div style="width: 100%;">
    <InfoHeader msg="Go2Jen" />
    <div style="">
      <CodeElement
        id="input"
        v-model="code"
        class="code"
      />

      <!-- this div is horizontally centered -->
      <div class="generated-hint">
        \/ \/ \/ \/ GENERATED CODE WILL APPEAR BELOW \/ \/ \/ \/
      </div>

      <CodeElement
        id="output"
        v-model="jenCode"
        class="code"
      />
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import CodeElement from '../components/CodeElement.vue'
import InfoHeader from '../components/InfoHeader.vue'

export default defineComponent({
  name: 'Go2Jen',
  components: { InfoHeader, CodeElement },
  data () {
    /* eslint-disable no-tabs */
    return {
      code: `package a

func hello(arg string) {
	// enter code here
	fmt.Println("Hello", arg)
}` as string,
      /* eslint-enable no-tabs */
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
.generated-hint {
    text-align: center;
    font-size: 1.2em;
    color: #999;
    font-weight: bold;
}
</style>

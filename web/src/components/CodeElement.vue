<template>
  <div class="container">
    <div ref="codemirrorRef" />
  </div>
</template>

<script lang="ts">
import {
  defineComponent,
  computed,
  ref,
  onMounted,
  watch,
  PropType
} from 'vue'
import CodeMirror, { ModeSpec } from 'codemirror'

import 'codemirror/mode/meta'
import 'codemirror/mode/go/go.js'
import 'codemirror/addon/display/placeholder.js'
import 'codemirror/addon/selection/active-line.js'
import 'codemirror/addon/selection/mark-selection.js'
import 'codemirror/addon/search/match-highlighter.js'
import 'codemirror/addon/display/autorefresh.js'
import 'codemirror/addon/edit/matchbrackets.js'
import 'codemirror/lib/codemirror.css'

export default defineComponent({
  name: 'CodeElement',
  props: {
    modelValue: {
      type: [String] as PropType<string>,
      default: null
    }
  },
  emits: ['update:model-value'],
  setup (props, { emit }) {
    const language = 'go'
    const placeholder = 'Write some code here...'

    const codemirrorRef = ref<HTMLTextAreaElement | null>(null)
    let codemirror: CodeMirror.Editor | null

    onMounted(() => {
      if (codemirrorRef.value) {
        codemirror = CodeMirror(codemirrorRef.value, {
          ...options.value,
          value: stringValue.value
        })

        codemirror.setOption('mode', { name: language })

        codemirror.on('change', (cm, { origin }) => {
          if (origin === 'setValue') return
          emit('update:model-value', cm.getValue())
        })
      }
    })

    const stringValue = computed<string>(() => {
      if (props.modelValue === null) return ''
      return props.modelValue as string
    })

    watch(stringValue, () => {
      if (codemirror == null) {
        return
      }

      if (codemirror.getValue() !== stringValue.value) {
        codemirror.setValue(stringValue.value || '')
      }
    })

    const options = computed<Record<string, any>>(() => {
      const defaultOptions: CodeMirror.EditorConfiguration = {
        autoRefresh: true,
        styleActiveLine: true,
        matchBrackets: true,
        mode: language,
        theme: 'default',
        tabSize: 4,
        indentUnit: 4,
        lineNumbers: true,
        readOnly: false,
        placeholder: placeholder,
        showCursorWhenSelecting: true
      }

      return Object.assign({}, defaultOptions)
    })

    return { options, codemirrorRef, stringValue }
  }
})
</script>

<style scoped>
.container {
  position: relative;
  width: 100%;
  font-size: 14px;
}
</style>

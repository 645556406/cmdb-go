<!-- src/components/Terminal.vue -->
<template>
  <div ref="terminal" class="terminal-container" />
</template>

<script>
import { Terminal } from 'xterm'
import { FitAddon } from 'xterm-addon-fit'
import 'xterm/css/xterm.css'

export default {
  // eslint-disable-next-line vue/require-prop-types
  // props: ['host', 'username', 'password'],
  data() {
    return {
      term: null,
      socket: null,
      fitAddon: new FitAddon()
    }
  },
  mounted() {
    this.initTerminal()
    this.connectWebSocket()
    window.addEventListener('resize', this.handleResize)
  },
  beforeDestroy() {
    if (this.socket) this.socket.close()
    window.removeEventListener('resize', this.handleResize)
  },
  methods: {
    initTerminal() {
      this.term = new Terminal({
        cursorBlink: true,
        fontSize: 14,
        fontFamily: 'Courier New'
      })
      this.term.loadAddon(this.fitAddon)
      this.term.open(this.$refs.terminal)
      this.fitAddon.fit()

      this.term.onData(data => {
        if (this.socket) this.socket.send(data)
      })
    },
    connectWebSocket() {
      // 从 Vuex 获取参数
      // const { host, username, password } = this.$store.state.terminal.sshParams
      // if (!host || !username || !password) {
      //   console.error('缺少 SSH 参数！')
      //   return
      // }
      const { host, username, password } = JSON.parse(sessionStorage.getItem('sshParams'))
      this.socket = new WebSocket('ws://localhost:8080/api/v1/ssh/connect')
      this.socket.onopen = () => {
        // 发送SSH连接参数
        this.socket.send(JSON.stringify({
          host: host,
          username: username,
          password: password // 实际项目应使用加密传输
        }))
      }

      this.socket.onmessage = (event) => {
        this.term.write(event.data)
      }

      this.socket.onclose = () => {
        this.term.writeln('\r\n连接已关闭')
      }
    },
    handleResize() {
      if (this.fitAddon) {
        this.fitAddon.fit()
        // 通知后端调整终端大小（可选）
        const { cols, rows } = this.term
        this.socket.send(JSON.stringify({ type: 'resize', cols, rows }))
      }
    }
  }
}
</script>

<style scoped>
    .terminal-container {
        width: 100%;
        height: 100%;
      /* height: calc(100vh - 60px); /* 计算 terminal 窗口大小 -60px 是 header的高度 */
        padding: 10px;
        background: #000;
        overflow: hidden; /* 禁止滚动 */
    }
</style>

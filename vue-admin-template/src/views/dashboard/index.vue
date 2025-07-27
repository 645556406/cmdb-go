<template>
  <div class="dashboard-container">
    <div>
      <el-row :gutter="12" style=" font-weight: bold;">
        <el-col :span="8">
          <el-card shadow="hover" style="color: blue;">
            服务器数量：{{ data.total }}
          </el-card>
        </el-col>
        <el-col :span="8">
          <el-card shadow="hover" style="color: green;">
            在线：{{ data.online }}
          </el-card>
        </el-col>
        <el-col :span="8">
          <el-card shadow="hover" style="color: red;">
            离线：{{ data.offline }}
          </el-card>
        </el-col>
      </el-row>
    </div>
  </div>
</template>

<script>
import { countServer } from '@/api/table'

export default {
  name: 'Dashboard',
  data() {
    return {
      data: {
        total: 0,
        online: 0,
        offline: 0
      }
    }
  },
  mounted() {
    this.conn = new WebSocket('ws://localhost:8080/api/v1/server/update')
    this.conn.onmessage = (e) => {
      this.data = { ...this.data, ...JSON.parse(e.data) }
    }
    // var conn = new WebSocket('ws://localhost:8080/api/v1/server/update')
    // conn.onopen = function(e) {
    //   console.log('连接已打开')
    //   conn.send('Hello, World!')
    // }
    // conn.onmessage = function(e) {
    //   console.log('从服务器收到消息:', e.data)
    //   this.data = JSON.parse(e.data)
    //   console.log(this.data)
    // }
    // conn.onclose = function(e) {
    //   console.log('连接已关闭')
    // }
  },
  beforeDestroy() {
    this.socket.disconnect()
  },
  methods: {
    fetchCountServer() {
      countServer().then(response => {
        this.data = response.data
      })
    }
  }
}
</script>

<style lang="scss" scoped>
.dashboard {
  &-container {
    margin: 30px;
  }
  &-text {
    font-size: 30px;
    line-height: 46px;
  }
}
</style>

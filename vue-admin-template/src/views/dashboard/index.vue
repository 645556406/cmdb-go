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
    this.initWebSocket()
  },
  beforeDestroy() {
    // this.socket.disconnect()
    if (this.conn && this.conn.readyState === WebSocket.OPEN) {
      console.log('WebSocket 连接关闭成功')
      this.conn.close()
    }
  },
  methods: {
    fetchCountServer() {
      countServer().then(response => {
        this.data = response.data
      })
    },
    initWebSocket() {
      this.conn = new WebSocket('ws://localhost:8080/api/v1/server/update')
      this.conn.onopen = () => {
        console.log('WebSocket 连接建立成功')
        this.reconnectAttempts = 0
      }
      this.conn.onmessage = (e) => {
        try {
          this.data = { ...this.data, ...JSON.parse(e.data) }
        } catch (error) {
          console.error('消息解析失败: ', error)
        }
      }
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

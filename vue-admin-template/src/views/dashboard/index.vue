<template>
  <div class="dashboard-container">
    <div>
      <el-row :gutter="12" style=" font-weight: bold;">
        <el-col :span="8">
          <el-card shadow="hover" style="color: blue;">
            服务器数量：{{ total }}
          </el-card>
        </el-col>
        <el-col :span="8">
          <el-card shadow="hover" style="color: green;">
            在线：{{ online }}
          </el-card>
        </el-col>
        <el-col :span="8">
          <el-card shadow="hover" style="color: red;">
            离线：{{ offline }}
          </el-card>
        </el-col>
      </el-row>
    </div>
  </div>
</template>

<script>
import { countServer, onlineServer, offlineServer } from '@/api/table'

export default {
  name: 'Dashboard',
  data() {
    return {
      total: 0,
      online: 0,
      offline: 0
    }
  },
  created() {
    setInterval(this.fetchCountServer, 10000)
    this.fetchCountServer()
  },
  methods: {
    fetchCountServer() {
      countServer().then(response => {
        this.total = response.data
      })
      onlineServer().then(response => {
        this.online = response.data
      })
      offlineServer().then(response => {
        this.offline = response.data
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

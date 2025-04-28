<template>
  <div class="app-container">
    <div style="display: flex;justify-content: left; align-items: center; gap: 10px; box-shadow: 0 2px 4px rgba(0, 0, 0, .12), 0 0 6px rgba(0, 0, 0, .04); height: 100px;">
      <!-- 搜索 -->
      <el-input v-model="searchQuery" size="mini" placeholder="请输入要搜索的IP" style="width: 200px; margin-left: 30px;" />
      <el-button type="primary" size="mini" @click="searchOneByIP">查询</el-button>
      <!-- 触发按钮 -->
      <el-button type="primary" size="mini" @click="dialogVisible = true">新增服务器</el-button>
    </div>
    <div>
      <!-- 对话框表单 -->
      <el-dialog title="填写服务器信息" :visible.sync="dialogVisible" width="auto;">
        <el-form :model="form" label-width="auto">
          <el-form-item label="IP">
            <el-input v-model="form.IP" />
          </el-form-item>
          <el-form-item label="HostName">
            <el-input v-model="form.HostName" />
          </el-form-item>
          <el-form-item label="Env">
            <el-input v-model="form.Env" />
          </el-form-item>
          <el-form-item label="Owner">
            <el-input v-model="form.Owner" />
          </el-form-item>
          <el-form-item label="OS">
            <el-input v-model="form.OS" />
          </el-form-item>
          <el-form-item label="Area">
            <el-input v-model="form.Area" />
          </el-form-item>
          <el-form-item label="Username">
            <el-input v-model="form.Username" />
          </el-form-item>
          <el-form-item label="Password">
            <el-input v-model="form.Password" type="password" />
          </el-form-item>
          <el-form-item label="PublicKey">
            <el-input v-model="form.PublicKey" />
          </el-form-item>
        </el-form>
        <span slot="footer">
          <el-button @click="dialogVisible = false">取消</el-button>
          <el-button type="primary" @click="submitForm">提交</el-button>
        </span>
      </el-dialog>
    </div>
    <div>
      <el-table
        v-loading="listLoading"
        :data="list"
        :height="620"
        element-loading-text="Loading"
        border
        fit
        highlight-current-row
        style="width: 100%;margin-top:30px; gap: 10px; box-shadow: 0 2px 4px rgba(0, 0, 0, .12), 0 0 6px rgba(0, 0, 0, .04);align-items: center;"
      >
        <el-table-column align="center" label="ID" fixed="left">
          <template slot-scope="scope">
            {{ scope.row.ID }}
          </template>
        </el-table-column>
        <el-table-column label="HostName" align="center" min-width="120">
          <template slot-scope="scope">
            {{ scope.row.Hostname }}
          </template>
        </el-table-column>
        <el-table-column label="IP" align="center" min-width="120">
          <template slot-scope="scope">
            <span>{{ scope.row.IP }}</span>
          </template>
        </el-table-column>
        <el-table-column label="Env" align="center">
          <template slot-scope="scope">
            {{ scope.row.Env }}
          </template>
        </el-table-column>
        <el-table-column class-name="status-col" label="OS" align="center">
          <template slot-scope="scope">
            <el-tag :type="scope.row.status | statusFilter">{{ scope.row.OS }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column label="CPU" align="center">
          <template slot-scope="scope">
            {{ scope.row.CPU }}
          </template>
        </el-table-column>
        <el-table-column label="Memory" align="center">
          <template slot-scope="scope">
            {{ scope.row.Memory }}
          </template>
        </el-table-column>
        <el-table-column label="Owner" align="center">
          <template slot-scope="scope">
            {{ scope.row.Owner }}
          </template>
        </el-table-column>
        <el-table-column label="Area" align="center">
          <template slot-scope="scope">
            {{ scope.row.Area }}
          </template>
        </el-table-column>
        <el-table-column label="Username" align="center">
          <template slot-scope="scope">
            {{ scope.row.Username }}
          </template>
        </el-table-column>
        <el-table-column label="Password" align="center">
          <template slot-scope="scope">
            {{ scope.row.Password }}
          </template>
        </el-table-column>
        <el-table-column label="PublicKey" align="center">
          <template slot-scope="scope">
            {{ truncateText(scope.row.PublicKey) }}
          </template>
        </el-table-column>
        <el-table-column label="Status" align="center">
          <template slot-scope="scope">
            {{ scope.row.Status }}
          </template>
        </el-table-column>
        <!-- <el-table-column align="center" prop="created_at" label="CreateTime" width="250">
          <template slot-scope="scope">
            <i class="el-icon-time" />
            <span>{{ scope.row.CreatedAt }}</span>
          </template>
        </el-table-column>
        <el-table-column align="center" prop="updated_at" label="UpdateTime" width="250">
          <template slot-scope="scope">
            <i class="el-icon-time" />
            <span>{{ scope.row.UpdatedAt }}</span>
          </template>
        </el-table-column> -->
        <el-table-column align="center" label="Actions" width="250" fixed="right" class-name="small-padding fixed-width">
          <template slot-scope="scope">
            <el-button type="primary" size="mini" @click="openTerminalSafe(scope.row)">安全连接</el-button>
            <el-button type="primary" size="mini" @click="dialog = true, row = scope.row">编辑</el-button>
            <el-button type="danger" size="mini" @click="confirmRole(scope.row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
    </div>
    <div>
      <el-drawer
        ref="drawer"
        title="修改服务器信息"
        :before-close="handleClose"
        :visible.sync="dialog"
        direction="ltr"
        custom-class="demo-drawer"
      >
        <div class="demo-drawer__content" style="display: flex; flex-direction: column; align-items: center; height: 100%;">
          <el-form :model="row">
            <el-form-item label="ID" :label-width="formLabelWidth">
              <el-input v-model="row.ID" autocomplete="off" />
            </el-form-item>
            <el-form-item label="HostName" :label-width="formLabelWidth">
              <el-input v-model="row.Hostname" autocomplete="off" />
            </el-form-item>
            <el-form-item label="IP" :label-width="formLabelWidth">
              <el-input v-model="row.IP" autocomplete="off" />
            </el-form-item>
            <el-form-item label="Env" :label-width="formLabelWidth">
              <el-input v-model="row.Env" autocomplete="off" />
            </el-form-item>
            <el-form-item label="OS" :label-width="formLabelWidth">
              <el-input v-model="row.OS" autocomplete="off" />
            </el-form-item>
            <el-form-item label="CPU" :label-width="formLabelWidth">
              <el-input v-model="row.CPU" autocomplete="off" />
            </el-form-item>
            <el-form-item label="Memory" :label-width="formLabelWidth">
              <el-input v-model="row.Memory" autocomplete="off" />
            </el-form-item>
            <el-form-item label="Owner" :label-width="formLabelWidth">
              <el-input v-model="row.Owner" autocomplete="off" />
            </el-form-item>
            <el-form-item label="Username" :label-width="formLabelWidth">
              <el-input v-model="row.Username" autocomplete="off" />
            </el-form-item>
            <el-form-item label="Password" :label-width="formLabelWidth">
              <el-input v-model="row.Password" autocomplete="off" />
            </el-form-item>
            <el-form-item label="Area" :label-width="formLabelWidth">
              <el-input v-model="row.Area" autocomplete="off" />
            </el-form-item>
            <el-form-item label="PublicKey" :label-width="formLabelWidth">
              <el-input v-model="row.PublicKey" autocomplete="off" />
            </el-form-item>
          </el-form>
          <div class="demo-drawer__footer">
            <el-button @click="cancelForm">取 消</el-button>
            <el-button type="primary" :loading="loading" @click="$refs.drawer.closeDrawer()">{{ loading ? '提交中 ...' : '确 定' }}</el-button>
          </div>
        </div>
      </el-drawer>
    </div>
    <br>
    <div>
      <el-pagination
        :current-page="currentPage"
        :page-sizes="[10, 20, 30, 50]"
        :page-size="pageSize"
        layout="total, sizes, prev, pager, next, jumper"
        :total="total"
        @size-change="handleSizeChange"
        @current-change="handleCurrentChange"
      />
    </div>
  </div>
</template>

<script>
import { getList, updateServer, delServer, addServer, getOneByID } from '@/api/table'

export default {
  filters: {
    statusFilter(status) {
      const statusMap = {
        published: 'success',
        draft: 'gray',
        deleted: 'danger'
      }
      return statusMap[status]
    }
  },
  data() {
    return {
      currentPage: 1, // 当前页码
      pageSize: 10, // 每页条数
      total: 0, // 总数据量
      searchQuery: '',
      row: {},
      table: false,
      dialog: false,
      loading: false,
      formLabelWidth: '80px',
      timer: null,
      drawer: false,
      direction: 'rtl',
      list: null,
      listLoading: true,
      tableKey: 0,
      form: {
        IP: '',
        Env: '',
        Owner: ''
      },
      dialogVisible: false
    }
  },
  created() {
    this.fetchData()
  },
  methods: {
    truncateText(text) {
      return text.length > 15 ? text.slice(0, 15) + '...' : text
    },
    openTerminalSafe(row) {
      console.log(row.IP)
      const params = {
        host: row.IP,
        username: row.Username
        // password: 'sls123'
      }
      // 1. 先存储到 Vuex，注意因为 Vuex 是单页面存储，不能跨页面使用
      this.$store.commit('terminal/setSSHParams', params)
      sessionStorage.setItem('sshParams', JSON.stringify(params))
      // 2. 打开新窗口（不带敏感参数）
      const route = this.$router.resolve({ name: 'Terminal' })
      window.open(route.href)
    },
    // 每页条数改变
    handleSizeChange(val) {
      this.pageSize = val
      this.fetchData()
    },
    // 当前页改变
    handleCurrentChange(val) {
      this.currentPage = val
      this.fetchData()
    },
    // 搜索还是有问题
    searchOneByIP() {
      this.listLoading = true
      getOneByID(this.searchQuery.trim()).then((response) => {
        if (response.code === 20000) {
          const serverList = []
          serverList.push(response.data)
          this.list = serverList
          this.listLoading = false
        } else {
          this.$message({
            type: 'faild',
            message: response.message
          })
        }
      }).catch(
        console.log('数据未查询到！'),
        this.listLoading = false
      )
    },
    submitForm() {
      addServer(this.form).then((response) => {
        console.log(response)
        if (response.code === 20000) {
          this.$message({
            type: 'success',
            message: response.message
          })
          this.dialogVisible = false
          this.fetchData()
        } else {
          this.$message({
            type: 'faild',
            message: response.message
          })
        }
      }).catch(error => {
        // 错误处理
        console.log(error)
      })
    },
    handleClose(done) {
      if (this.loading) {
        return
      }
      this.$confirm('确定要提交表单吗？')
        .then(_ => {
          this.loading = true
          // 提交前明确转换数据类型
          const payload = {
            ID: Number(this.row.ID),
            HostName: String(this.row.Hostname),
            CPU: Number(this.row.CPU),
            IP: String(this.row.IP),
            OS: String(this.row.OS),
            Memory: Number(this.row.Memory),
            Owner: String(this.row.Owner),
            Env: String(this.row.Env),
            Status: Number(this.row.Status),
            Username: String(this.row.Username),
            Password: String(this.row.Password),
            Area: String(this.row.Area),
            PublicKey: String(this.row.PublicKey)
          }
          // 调用提交方法
          updateServer(payload).then(response => {
            this.$message({
              type: 'success',
              message: response.message
            })
            this.loading = false
          })
          done()
        })
        .catch(_ => {
          done()
        })
    },
    cancelForm() {
      this.loading = false
      this.dialog = false
      clearTimeout(this.timer)
    },
    fetchData() {
      this.listLoading = true
      getList().then(response => {
        this.total = (response.data).length // 假设总共有100条数据
        // 根据分页参数截取数据
        const start = (this.currentPage - 1) * this.pageSize
        const end = start + this.pageSize
        this.list = response.data.slice(start, end)
        this.listLoading = false
      })
    },
    // fetchData() {
    //   this.listLoading = true
    //   getList().then(response => {
    //     console.log(response)
    //     this.list = response.data
    //     this.listLoading = false
    //   })
    // },
    confirmRole(row) {
      this.$confirm('此操作将永久删除该服务器, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        delServer(row).then((response) => {
          if (response.code === 20000) {
            this.$message({
              type: 'success',
              message: response.message
            })
            this.fetchData()
          } else {
            this.$message({
              type: 'faild',
              message: response.message
            })
          }
        }).catch(error => {
          // 错误处理
          console.log(error)
        })
      }).catch(() => {
        this.$message({
          type: 'info',
          message: '已取消删除'
        })
      })
    }
  }
}
</script>
<style lang="scss" scoped>
.app-container {
  .roles-table {
    margin-top: 30px;
  }
  .permission-tree {
    margin-bottom: 30px;
  }
}
</style>

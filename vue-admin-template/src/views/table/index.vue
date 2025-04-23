<template>
  <div class="app-container">
    <div style="display: flex;justify-content: left; align-items: center; gap: 10px; ">
      <!-- 触发按钮 -->
      <el-button type="primary" @click="dialogVisible = true">新增服务器</el-button>
      <!-- 搜索 -->
      <el-input v-model="searchQuery" placeholder="请输入搜索内容" style="width: 200px" />
      <el-button type="primary" @click="dialogVisible = true">查询</el-button>
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
        element-loading-text="Loading"
        border
        fit
        highlight-current-row
        style="width: 100%;margin-top:30px;"
      >
        <el-table-column align="center" label="ID" width="95">
          <template slot-scope="scope">
            {{ scope.row.ID }}
          </template>
        </el-table-column>
        <el-table-column label="HostName">
          <template slot-scope="scope">
            {{ scope.row.Hostname }}
          </template>
        </el-table-column>
        <el-table-column label="IP" width="110" align="center">
          <template slot-scope="scope">
            <span>{{ scope.row.IP }}</span>
          </template>
        </el-table-column>
        <el-table-column label="Env" width="110" align="center">
          <template slot-scope="scope">
            {{ scope.row.Env }}
          </template>
        </el-table-column>
        <el-table-column class-name="status-col" label="OS" width="110" align="center">
          <template slot-scope="scope">
            <el-tag :type="scope.row.status | statusFilter">{{ scope.row.OS }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column label="CPU ">
          <template slot-scope="scope">
            {{ scope.row.CPU }}
          </template>
        </el-table-column>
        <el-table-column label="Memory ">
          <template slot-scope="scope">
            {{ scope.row.Memory }}
          </template>
        </el-table-column>
        <el-table-column label="Owner ">
          <template slot-scope="scope">
            {{ scope.row.Owner }}
          </template>
        </el-table-column>
        <el-table-column label="Status ">
          <template slot-scope="scope">
            {{ scope.row.Status }}
          </template>
        </el-table-column>
        <el-table-column align="center" prop="created_at" label="CreateTime" width="200">
          <template slot-scope="scope">
            <i class="el-icon-time" />
            <span>{{ scope.row.CreatedAt }}</span>
          </template>
        </el-table-column>
        <el-table-column align="center" prop="updated_at" label="UpdateTime" width="200">
          <template slot-scope="scope">
            <i class="el-icon-time" />
            <span>{{ scope.row.UpdatedAt }}</span>
          </template>
        </el-table-column>
        <el-table-column align="center" label="Actions" width="235" class-name="small-padding fixed-width">
          <template slot-scope="scope">
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
          </el-form>
          <div class="demo-drawer__footer">
            <el-button @click="cancelForm">取 消</el-button>
            <el-button type="primary" :loading="loading" @click="$refs.drawer.closeDrawer()">{{ loading ? '提交中 ...' : '确 定' }}</el-button>
          </div>
        </div>
      </el-drawer>
    </div>
  </div>
</template>

<script>
import { getList, updateServer, delServer, addServer } from '@/api/table'

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
            ID: Number(this.row.id),
            HostName: String(this.row.HostName),
            CPU: Number(this.row.CPU),
            IP: String(this.row.IP),
            OS: String(this.row.OS),
            Memory: Number(this.row.Memory),
            Owner: String(this.row.Owner),
            Env: String(this.row.Env),
            Status: Number(this.row.Status)
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
        console.log(response)
        this.list = response.data
        this.listLoading = false
      })
    },
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

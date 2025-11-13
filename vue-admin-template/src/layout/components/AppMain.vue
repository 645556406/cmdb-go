<template>
  <section class="app-main">
    <transition name="fade-transform" mode="out-in">
      <router-view :key="key" />
    </transition>
  </section>
</template>

<script>
export default {
  name: 'AppMain',
  computed: {
    key() {
      return this.$route.path
    }
  }
}
</script>

<style lang="scss" scoped>
@import '@/styles/variables.scss';

.app-main {
  /*60 = navbar  */
  min-height: calc(100vh - 60px);
  width: 100%;
  position: relative;
  overflow: hidden;
  background: linear-gradient(135deg, rgba(248, 250, 252, 0.8) 0%, rgba(241, 245, 249, 0.8) 100%);
  padding: 20px;
}

.fixed-header+.app-main {
  padding-top: 80px;
}

// 页面切换动画
.fade-transform-leave-active,
.fade-transform-enter-active {
  transition: all 0.3s ease;
}

.fade-transform-enter {
  opacity: 0;
  transform: translateX(30px);
}

.fade-transform-leave-to {
  opacity: 0;
  transform: translateX(-30px);
}
</style>

<style lang="scss">
@import '@/styles/variables.scss';

// fix css style bug in open el-dialog
.el-popup-parent--hidden {
  .fixed-header {
    padding-right: 15px;
  }
}

// 全局页面容器样式优化
.app-main {
  ::v-deep .app-container {
    background: rgba(255, 255, 255, 0.9);
    backdrop-filter: blur(10px);
    border-radius: 16px;
    box-shadow: 0 8px 32px rgba(0, 0, 0, 0.08);
    border: 1px solid rgba(255, 255, 255, 0.2);
    padding: 24px;
    margin-bottom: 20px;
    transition: all 0.3s ease;
    
    &:hover {
      box-shadow: 0 12px 40px rgba(0, 0, 0, 0.12);
      transform: translateY(-2px);
    }
    
    .filter-container {
      background: rgba(248, 250, 252, 0.5);
      border-radius: 12px;
      padding: 16px;
      margin-bottom: 20px;
      border: 1px solid $border-color;
      
      .filter-item {
        margin-right: 12px;
        margin-bottom: 8px;
        
        &.el-button {
          border-radius: 8px;
          transition: all 0.3s ease;
          
          &:hover {
            transform: translateY(-1px);
            box-shadow: 0 4px 12px rgba(16, 185, 129, 0.2);
          }
        }
      }
    }
    
    // 表格容器样式
    .table-container {
      background: white;
      border-radius: 12px;
      padding: 16px;
      box-shadow: 0 2px 8px rgba(0, 0, 0, 0.06);
      
      .el-table {
        border-radius: 8px;
        overflow: hidden;
        
        &::before {
          display: none;
        }
        
        th {
          background: linear-gradient(135deg, rgba(16, 185, 129, 0.05) 0%, rgba(59, 130, 246, 0.05) 100%);
          color: $text-primary;
          font-weight: 600;
          border-bottom: 2px solid $border-color;
        }
        
        td {
          border-bottom: 1px solid $border-color;
          transition: all 0.3s ease;
          
          &:hover {
            background: rgba(16, 185, 129, 0.02);
          }
        }
        
        tr {
          transition: all 0.3s ease;
          
          &:hover {
            background: rgba(16, 185, 129, 0.03);
          }
        }
      }
    }
    
    // 分页样式
    .pagination-container {
      display: flex;
      justify-content: center;
      margin-top: 20px;
      
      .el-pagination {
        .el-pager li {
          border-radius: 6px;
          margin: 0 2px;
          transition: all 0.3s ease;
          
          &:hover {
            transform: translateY(-1px);
            box-shadow: 0 2px 8px rgba(16, 185, 129, 0.2);
          }
          
          &.active {
            background: linear-gradient(135deg, $primary-color 0%, $secondary-color 100%);
            color: white;
          }
        }
        
        .btn-prev,
        .btn-next {
          border-radius: 6px;
          transition: all 0.3s ease;
          
          &:hover {
            transform: translateY(-1px);
            box-shadow: 0 2px 8px rgba(16, 185, 129, 0.2);
          }
        }
      }
    }
  }
}
</style>

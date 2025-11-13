<template>
  <div :class="{'has-logo':showLogo}">
    <logo v-if="showLogo" :collapse="isCollapse" />
    <el-scrollbar wrap-class="scrollbar-wrapper">
      <el-menu
        :default-active="activeMenu"
        :collapse="isCollapse"
        :background-color="variables.menuBg"
        :text-color="variables.menuText"
        :unique-opened="false"
        :active-text-color="variables.menuActiveText"
        :collapse-transition="false"
        mode="vertical"
        class="sidebar-menu"
      >
        <sidebar-item v-for="route in routes" :key="route.path" :item="route" :base-path="route.path" />
      </el-menu>
    </el-scrollbar>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
import Logo from './Logo'
import SidebarItem from './SidebarItem'
import variables from '@/styles/variables.scss'

export default {
  components: { SidebarItem, Logo },
  computed: {
    ...mapGetters([
      'sidebar'
    ]),
    routes() {
      return this.$router.options.routes
    },
    activeMenu() {
      const route = this.$route
      const { meta, path } = route
      // if set path, the sidebar will highlight the path you set
      if (meta.activeMenu) {
        return meta.activeMenu
      }
      return path
    },
    showLogo() {
      return this.$store.state.settings.sidebarLogo
    },
    variables() {
      return variables
    },
    isCollapse() {
      return !this.sidebar.opened
    }
  }
}
</script>

<style lang="scss" scoped>
@import '@/styles/variables.scss';

.sidebar-menu {
  border-right: none;
  
  // 菜单项样式优化
  ::v-deep .el-menu-item {
    margin: 4px 8px;
    border-radius: 8px;
    transition: all 0.3s ease;
    border: 1px solid transparent;
    
    &:hover {
      background: rgba(16, 185, 129, 0.1) !important;
      color: $primary-color !important;
      border-color: rgba(16, 185, 129, 0.2);
      transform: translateX(2px);
    }
    
    &.is-active {
      background: linear-gradient(135deg, $primary-color 0%, $secondary-color 100%) !important;
      color: white !important;
      border-color: $primary-color;
      box-shadow: 0 4px 12px rgba(16, 185, 129, 0.3);
      
      &::before {
        content: '';
        position: absolute;
        left: 0;
        top: 50%;
        transform: translateY(-50%);
        width: 3px;
        height: 20px;
        background: white;
        border-radius: 0 2px 2px 0;
      }
    }
  }
  
  // 子菜单样式优化
  ::v-deep .el-submenu {
    margin: 4px 8px;
    border-radius: 8px;
    overflow: hidden;
    
    .el-submenu__title {
      transition: all 0.3s ease;
      border: 1px solid transparent;
      border-radius: 8px;
      
      &:hover {
        background: rgba(16, 185, 129, 0.1) !important;
        color: $primary-color !important;
        border-color: rgba(16, 185, 129, 0.2);
        transform: translateX(2px);
      }
    }
    
    &.is-active {
      .el-submenu__title {
        background: linear-gradient(135deg, $primary-color 0%, $secondary-color 100%) !important;
        color: white !important;
        border-color: $primary-color;
        box-shadow: 0 4px 12px rgba(16, 185, 129, 0.3);
      }
    }
    
    .el-menu {
      background: rgba(255, 255, 255, 0.5);
      
      .el-menu-item {
        margin: 2px 8px;
        padding-left: 40px !important;
        
        &:hover {
          background: rgba(16, 185, 129, 0.1) !important;
          color: $primary-color !important;
        }
        
        &.is-active {
          background: rgba(16, 185, 129, 0.2) !important;
          color: $primary-color !important;
        }
      }
    }
  }
  
  // 折叠状态样式
  &.el-menu--collapse {
    ::v-deep .el-menu-item,
    ::v-deep .el-submenu__title {
      text-align: center;
      padding: 0 20px;
    }
  }
}

// 滚动条样式
::v-deep .scrollbar-wrapper {
  .el-scrollbar__bar {
    &.is-vertical {
      right: 2px;
      width: 4px;
      
      .el-scrollbar__thumb {
        background: rgba(16, 185, 129, 0.3);
        border-radius: 2px;
        
        &:hover {
          background: rgba(16, 185, 129, 0.5);
        }
      }
    }
  }
}
</style>

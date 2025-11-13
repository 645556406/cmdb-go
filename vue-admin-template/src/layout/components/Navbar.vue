<template>
  <div class="navbar">
    <hamburger :is-active="sidebar.opened" class="hamburger-container" @toggleClick="toggleSideBar" />

    <breadcrumb class="breadcrumb-container" />

    <div class="right-menu">
      <theme-switcher class="right-menu-item hover-effect" />
      
      <el-dropdown class="avatar-container" trigger="click">
        <div class="avatar-wrapper">
          <img :src="avatar+'?imageView2/1/w/80/h/80'" class="user-avatar">
          <i class="el-icon-caret-bottom" />
        </div>
        <el-dropdown-menu slot="dropdown" class="user-dropdown">
          <router-link to="/">
            <el-dropdown-item>
              首页
            </el-dropdown-item>
          </router-link>
          <el-dropdown-item divided @click.native="logout">
            <span style="display:block;">退出登录</span>
          </el-dropdown-item>
        </el-dropdown-menu>
      </el-dropdown>
    </div>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
import Breadcrumb from '@/components/Breadcrumb'
import Hamburger from '@/components/Hamburger'
import ThemeSwitcher from '@/components/ThemeSwitcher'

export default {
  components: {
    Breadcrumb,
    Hamburger,
    ThemeSwitcher
  },
  data() {
    return {
      imageView2: 'vue-admin-template/src/icons/svg/user.svg'
    }
  },
  computed: {
    ...mapGetters([
      'sidebar',
      'avatar'
    ])
  },
  methods: {
    toggleSideBar() {
      this.$store.dispatch('app/toggleSideBar')
    },
    async logout() {
      await this.$store.dispatch('/logout')
      this.$router.push(`/login?redirect=${this.$route.fullPath}`)
    }
  }
}
</script>

<style lang="scss" scoped>
@import '@/styles/variables.scss';

.navbar {
  height: 60px;
  overflow: hidden;
  position: relative;
  background: linear-gradient(135deg, rgba(255, 255, 255, 0.95) 0%, rgba(248, 250, 252, 0.95) 100%);
  backdrop-filter: blur(10px);
  box-shadow: 0 2px 12px rgba(0, 0, 0, 0.08);
  border-bottom: 1px solid $border-color;

  .hamburger-container {
    line-height: 56px;
    height: 100%;
    float: left;
    cursor: pointer;
    transition: all 0.3s ease;
    -webkit-tap-highlight-color: transparent;
    padding: 0 16px;
    border-radius: 8px;
    margin: 2px 8px;

    &:hover {
      background: rgba(16, 185, 129, 0.1);
      transform: scale(1.05);
    }
  }

  .breadcrumb-container {
    float: left;
    margin-left: 8px;
  }

  .right-menu {
    float: right;
    height: 100%;
    line-height: 60px;
    display: flex;
    align-items: center;
    padding-right: 20px;

    &:focus {
      outline: none;
    }

    .right-menu-item {
      display: inline-flex;
      align-items: center;
      justify-content: center;
      padding: 0 12px;
      height: 36px;
      margin: 0 4px;
      font-size: 16px;
      color: $text-secondary;
      vertical-align: middle;
      border-radius: 8px;
      transition: all 0.3s ease;

      &.hover-effect {
        cursor: pointer;

        &:hover {
          background: rgba(16, 185, 129, 0.1);
          color: $primary-color;
          transform: translateY(-1px);
        }
      }
    }

    .avatar-container {
      margin-right: 8px;

      .avatar-wrapper {
        display: flex;
        align-items: center;
        padding: 6px 12px;
        border-radius: 12px;
        background: rgba(16, 185, 129, 0.05);
        border: 1px solid rgba(16, 185, 129, 0.2);
        cursor: pointer;
        transition: all 0.3s ease;

        &:hover {
          background: rgba(16, 185, 129, 0.1);
          border-color: $primary-color;
          transform: translateY(-1px);
          box-shadow: 0 4px 12px rgba(16, 185, 129, 0.2);
        }

        .user-avatar {
          cursor: pointer;
          width: 36px;
          height: 36px;
          border-radius: 50%;
          border: 2px solid $primary-color;
          margin-right: 8px;
          transition: all 0.3s ease;

          &:hover {
            transform: scale(1.1);
          }
        }

        .el-icon-caret-bottom {
          cursor: pointer;
          color: $text-secondary;
          font-size: 14px;
          transition: all 0.3s ease;

          &:hover {
            color: $primary-color;
          }
        }
      }
    }
  }
}

// 用户下拉菜单样式
::v-deep .user-dropdown {
  margin-top: 8px;
  border-radius: 12px;
  box-shadow: 0 12px 32px rgba(0, 0, 0, 0.12);
  border: 1px solid $border-color;
  overflow: hidden;

  .el-dropdown-menu__item {
    padding: 12px 20px;
    transition: all 0.3s ease;
    color: $text-secondary;

    &:hover {
      background: rgba(16, 185, 129, 0.1);
      color: $primary-color;
    }

    &.is-divided {
      border-top: 1px solid $border-color;
      margin-top: 4px;
      padding-top: 12px;
    }
  }
}
</style>

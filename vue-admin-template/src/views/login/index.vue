<template>
  <div class="login-container">
    <div class="login-background">
      <div class="floating-shapes">
        <div class="shape shape-1" />
        <div class="shape shape-2" />
        <div class="shape shape-3" />
        <div class="shape shape-4" />
      </div>
    </div>
    
    <el-form ref="loginForm" :model="loginForm" :rules="loginRules" class="login-form card-container" auto-complete="on" label-position="left">
      <div class="title-container">
        <div class="logo">
          <svg-icon icon-class="server" class="logo-icon" />
        </div>
        <h3 class="title">运维管理平台</h3>
        <p class="subtitle">安全 · 高效 · 智能</p>
      </div>

      <el-form-item prop="username">
        <span class="svg-container">
          <svg-icon icon-class="user" />
        </span>
        <el-input
          ref="username"
          v-model="loginForm.username"
          placeholder="请输入用户名"
          name="username"
          type="text"
          tabindex="1"
          auto-complete="on"
          class="form-input"
        />
      </el-form-item>

      <el-form-item prop="password">
        <span class="svg-container">
          <svg-icon icon-class="password" />
        </span>
        <el-input
          :key="passwordType"
          ref="password"
          v-model="loginForm.password"
          :type="passwordType"
          placeholder="请输入密码"
          name="password"
          tabindex="2"
          auto-complete="on"
          class="form-input"
          @keyup.enter.native="handleLogin"
        />
        <span class="show-pwd" @click="showPwd">
          <svg-icon :icon-class="passwordType === 'password' ? 'eye' : 'eye-open'" />
        </span>
      </el-form-item>

      <el-button :loading="loading" type="primary" class="login-btn" @click.native.prevent="handleLogin">
        <span v-if="!loading">登 录</span>
        <span v-else>登录中...</span>
      </el-button>

      <div class="tips">
        <span>请使用有效的管理员账户登录系统</span>
      </div>

    </el-form>
  </div>
</template>

<script>
import { validUsername } from '@/utils/validate'

export default {
  name: 'Login',
  data() {
    const validateUsername = (rule, value, callback) => {
      if (!validUsername(value)) {
        callback(new Error('Please enter the correct user name'))
      } else {
        callback()
      }
    }
    const validatePassword = (rule, value, callback) => {
      if (value.length < 6) {
        callback(new Error('The password can not be less than 6 digits'))
      } else {
        callback()
      }
    }
    return {
      loginForm: {
        username: '',
        password: ''
      },
      loginRules: {
        username: [{ required: true, trigger: 'blur', validator: validateUsername }],
        password: [{ required: true, trigger: 'blur', validator: validatePassword }]
      },
      loading: false,
      passwordType: 'password',
      redirect: undefined
    }
  },
  watch: {
    $route: {
      handler: function(route) {
        this.redirect = route.query && route.query.redirect
      },
      immediate: true
    }
  },
  methods: {
    showPwd() {
      if (this.passwordType === 'password') {
        this.passwordType = ''
      } else {
        this.passwordType = 'password'
      }
      this.$nextTick(() => {
        this.$refs.password.focus()
      })
    },
    handleLogin() {
      this.$refs.loginForm.validate(valid => {
        if (valid) {
          this.loading = true
          this.$store.dispatch('user/login', this.loginForm).then(() => {
            this.$router.push({ path: this.redirect || '/' })
            this.loading = false
          }).catch(() => {
            this.loading = false
          })
        } else {
          console.log('error submit!!')
          return false
        }
      })
    }
  }
}
</script>

<style lang="scss">
// 引入主题变量
@import '@/styles/variables.scss';

/* 修复input 背景不协调 和光标变色 */
$bg: transparent;
$light_gray: $text-primary;
$cursor: $primary-color;

@supports (-webkit-mask: none) and (not (cater-color: $cursor)) {
  .login-container .el-input input {
    color: $cursor;
  }
}

/* reset element-ui css */
.login-container {
  .el-input {
    display: inline-block;
    height: 50px;
    width: 85%;

    input {
      background: rgba(255, 255, 255, 0.9);
      border: 2px solid $border-color;
      -webkit-appearance: none;
      border-radius: 8px;
      padding: 12px 5px 12px 15px;
      color: $text-primary;
      height: 50px;
      caret-color: $cursor;
      font-size: 14px;
      transition: all 0.3s ease;

      &:focus {
        border-color: $primary-color;
        box-shadow: 0 0 0 3px rgba(16, 185, 129, 0.1);
        outline: none;
      }

      &:-webkit-autofill {
        box-shadow: 0 0 0 1000px rgba(255, 255, 255, 0.9) inset !important;
        -webkit-text-fill-color: $text-primary !important;
      }
    }
  }

  .el-form-item {
    border: none;
    background: rgba(255, 255, 255, 0.8);
    border-radius: 12px;
    margin-bottom: 24px;
    backdrop-filter: blur(10px);
    border: 1px solid rgba(255, 255, 255, 0.3);
    transition: all 0.3s ease;

    &:hover {
      background: rgba(255, 255, 255, 0.95);
      box-shadow: 0 8px 24px rgba(0, 0, 0, 0.1);
    }
  }
}
</style>

<style lang="scss" scoped>
// 引入主题变量
@import '@/styles/variables.scss';

.login-container {
  min-height: 100vh;
  width: 100%;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  overflow: hidden;
  position: relative;
  display: flex;
  align-items: center;
  justify-content: center;

  .login-background {
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    overflow: hidden;
    z-index: 0;

    .floating-shapes {
      position: absolute;
      width: 100%;
      height: 100%;

      .shape {
        position: absolute;
        border-radius: 50%;
        background: rgba(255, 255, 255, 0.1);
        animation: float 6s ease-in-out infinite;

        &.shape-1 {
          width: 80px;
          height: 80px;
          top: 20%;
          left: 10%;
          animation-delay: 0s;
        }

        &.shape-2 {
          width: 120px;
          height: 120px;
          top: 60%;
          right: 10%;
          animation-delay: 2s;
        }

        &.shape-3 {
          width: 60px;
          height: 60px;
          bottom: 20%;
          left: 20%;
          animation-delay: 4s;
        }

        &.shape-4 {
          width: 100px;
          height: 100px;
          top: 10%;
          right: 30%;
          animation-delay: 1s;
        }
      }
    }
  }

  .login-form {
    position: relative;
    width: 420px;
    max-width: 90%;
    padding: 48px 40px;
    margin: 0 auto;
    z-index: 10;
    backdrop-filter: blur(20px);
    background: rgba(255, 255, 255, 0.95);
    border: 1px solid rgba(255, 255, 255, 0.3);
    box-shadow: 0 20px 40px rgba(0, 0, 0, 0.1);
  }

  .tips {
    font-size: 13px;
    color: $text-secondary;
    text-align: center;
    margin-top: 20px;
    line-height: 1.5;
  }

  .svg-container {
    padding: 6px 5px 6px 15px;
    color: $primary-color;
    vertical-align: middle;
    width: 30px;
    display: inline-block;
    font-size: 16px;
  }

  .title-container {
    position: relative;
    text-align: center;
    margin-bottom: 40px;

    .logo {
      margin-bottom: 16px;
      .logo-icon {
        font-size: 48px;
        color: $primary-color;
        animation: pulse 2s infinite;
      }
    }

    .title {
      font-size: 28px;
      color: $text-primary;
      margin: 0 0 8px 0;
      font-weight: 700;
      letter-spacing: 1px;
    }

    .subtitle {
      font-size: 14px;
      color: $text-secondary;
      margin: 0;
      font-weight: 400;
      letter-spacing: 2px;
    }
  }

  .show-pwd {
    position: absolute;
    right: 15px;
    top: 50%;
    transform: translateY(-50%);
    font-size: 16px;
    color: $text-secondary;
    cursor: pointer;
    user-select: none;
    transition: color 0.3s ease;

    &:hover {
      color: $primary-color;
    }
  }

  .login-btn {
    width: 100%;
    height: 50px;
    background: linear-gradient(135deg, $primary-color 0%, $primary-light 100%);
    border: none;
    border-radius: 8px;
    font-size: 16px;
    font-weight: 600;
    letter-spacing: 1px;
    transition: all 0.3s ease;
    margin-top: 10px;

    &:hover {
      background: linear-gradient(135deg, $primary-dark 0%, $primary-color 100%);
      transform: translateY(-2px);
      box-shadow: 0 8px 24px rgba(16, 185, 129, 0.4);
    }

    &:active {
      transform: translateY(0);
    }
  }
}

// 动画效果
@keyframes float {
  0%, 100% {
    transform: translateY(0px) rotate(0deg);
  }
  50% {
    transform: translateY(-20px) rotate(180deg);
  }
}

@keyframes pulse {
  0%, 100% {
    transform: scale(1);
  }
  50% {
    transform: scale(1.05);
  }
}

// 响应式设计
@media (max-width: 768px) {
  .login-container {
    .login-form {
      width: 90%;
      padding: 32px 24px;
    }

    .title-container {
      .title {
        font-size: 24px;
      }

      .logo-icon {
        font-size: 40px;
      }
    }
  }
}
</style>

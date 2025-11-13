<template>
  <div class="theme-switcher">
    <el-tooltip content="切换主题" placement="bottom">
      <el-button
        type="text"
        class="theme-button"
        @click="toggleTheme"
      >
        <i :class="themeIcon" class="theme-icon"></i>
      </el-button>
    </el-tooltip>
  </div>
</template>

<script>
export default {
  name: 'ThemeSwitcher',
  data() {
    return {
      currentTheme: localStorage.getItem('theme') || 'light'
    }
  },
  computed: {
    themeIcon() {
      return this.currentTheme === 'light' ? 'el-icon-moon' : 'el-icon-sunny'
    }
  },
  mounted() {
    this.applyTheme()
  },
  methods: {
    toggleTheme() {
      this.currentTheme = this.currentTheme === 'light' ? 'dark' : 'light'
      this.applyTheme()
      localStorage.setItem('theme', this.currentTheme)
      this.$message.success(`已切换到${this.currentTheme === 'light' ? '浅色' : '深色'}主题`)
    },
    applyTheme() {
      const root = document.documentElement
      if (this.currentTheme === 'dark') {
        root.classList.add('dark-theme')
      } else {
        root.classList.remove('dark-theme')
      }
    }
  }
}
</script>

<style lang="scss" scoped>
@import '@/styles/variables.scss';

.theme-switcher {
  display: inline-block;
  
  .theme-button {
    padding: 8px 12px;
    border-radius: 8px;
    transition: all 0.3s ease;
    background: rgba(16, 185, 129, 0.05);
    border: 1px solid rgba(16, 185, 129, 0.2);
    
    &:hover {
      background: rgba(16, 185, 129, 0.1);
      border-color: $primary-color;
      transform: translateY(-1px);
      box-shadow: 0 4px 12px rgba(16, 185, 129, 0.2);
    }
    
    .theme-icon {
      font-size: 18px;
      color: $primary-color;
      transition: all 0.3s ease;
      
      &:hover {
        transform: scale(1.1);
      }
    }
  }
}
</style>
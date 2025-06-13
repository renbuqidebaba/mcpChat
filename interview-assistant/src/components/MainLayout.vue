<template>
  <div class="main-layout">
    <!-- 导航栏 -->
    <nav class="navbar">
      <div class="nav-brand">
        <h2>MCP Chat</h2>
      </div>
      <div class="nav-menu">
        <router-link to="/main" class="nav-link" active-class="active">聊天</router-link>
        <router-link to="/user" class="nav-link" active-class="active">用户中心</router-link>
        <span class="user-info">欢迎, {{ currentUser }}</span>
        <button @click="handleLogout" class="logout-btn">退出登录</button>
      </div>
    </nav>

    <!-- 主要内容区域 -->
    <main class="main-content">
      <ChatInterface />
    </main>
  </div>
</template>

<script>
import ChatInterface from './ChatInterface.vue';

export default {
  name: 'MainLayout',
  components: {
    ChatInterface
  },
  data() {
    return {
      currentUser: ''
    };
  },
  mounted() {
    this.loadCurrentUser();
  },
  methods: {
    loadCurrentUser() {
      this.currentUser = localStorage.getItem('currentUser') || '';
    },

    handleLogout() {
      // 清除本地存储
      localStorage.removeItem('currentUser');
      // 跳转到登录页
      this.$router.push('/');
    }
  }
};
</script>

<style scoped>
.main-layout {
  min-height: 100vh;
  display: flex;
  flex-direction: column;
}

.navbar {
  background: white;
  box-shadow: 0 2px 10px rgba(0,0,0,0.1);
  padding: 15px 30px;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.nav-brand h2 {
  color: #333;
  margin: 0;
}

.nav-menu {
  display: flex;
  align-items: center;
  gap: 20px;
}

.nav-link {
  text-decoration: none;
  color: #666;
  font-weight: 500;
  padding: 8px 16px;
  border-radius: 6px;
  transition: all 0.3s ease;
}

.nav-link:hover,
.nav-link.active {
  color: #667eea;
  background: #f8f9ff;
}

.user-info {
  color: #666;
  font-size: 14px;
}

.logout-btn {
  padding: 8px 16px;
  background: #dc3545;
  color: white;
  border: none;
  border-radius: 6px;
  cursor: pointer;
  font-weight: 500;
  transition: background 0.3s ease;
}

.logout-btn:hover {
  background: #c82333;
}

.main-content {
  flex: 1;
  padding: 0;
}

@media (max-width: 768px) {
  .navbar {
    padding: 15px 20px;
  }
  
  .nav-menu {
    gap: 10px;
  }
  
  .nav-link {
    padding: 6px 12px;
    font-size: 14px;
  }
  
  .user-info {
    display: none;
  }
}
</style>
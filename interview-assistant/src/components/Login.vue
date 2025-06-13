<template>
  <div class="login-container">
    <div class="login-card">
      <div class="login-header">
        <h1>{{ isLogin ? '欢迎回来' : '创建账户' }}</h1>
        <p>{{ isLogin ? '登录您的账户继续使用' : '注册新账户开始使用' }}</p>
      </div>

      <form @submit.prevent="handleSubmit" class="login-form">
        <div class="form-group">
          <label>用户名</label>
          <input 
            type="text" 
            v-model="form.username" 
            required 
            placeholder="请输入用户名"
            :disabled="loading"
          />
        </div>

        <div class="form-group" v-if="!isLogin">
          <label>邮箱</label>
          <input 
            type="email" 
            v-model="form.email" 
            required 
            placeholder="请输入邮箱"
            :disabled="loading"
          />
        </div>

        <div class="form-group">
          <label>密码</label>
          <input 
            type="password" 
            v-model="form.password" 
            required 
            placeholder="请输入密码"
            :disabled="loading"
          />
        </div>

        <button type="submit" :disabled="loading" class="submit-btn">
          {{ loading ? '处理中...' : (isLogin ? '登录' : '注册') }}
        </button>

        <div v-if="errorMessage" class="error-message">
          {{ errorMessage }}
        </div>

        <div v-if="successMessage" class="success-message">
          {{ successMessage }}
        </div>
      </form>

      <div class="login-footer">
        <p>
          {{ isLogin ? '还没有账户？' : '已有账户？' }}
          <a href="#" @click.prevent="toggleMode" class="toggle-link">
            {{ isLogin ? '立即注册' : '立即登录' }}
          </a>
        </p>
      </div>
    </div>
  </div>
</template>

<script>
import { login, register } from '../services/apiService';

export default {
  name: 'Login',
  data() {
    return {
      isLogin: true,
      loading: false,
      errorMessage: '',
      successMessage: '',
      form: {
        username: '',
        password: '',
        email: ''
      }
    };
  },
  mounted() {
    // 检查是否已登录
    this.checkLoginStatus();
  },
  methods: {
    checkLoginStatus() {
      const currentUser = localStorage.getItem('currentUser');
      if (currentUser && currentUser !== 'null' && currentUser !== 'undefined') {
        // 已登录，直接跳转到主页
        this.$router.replace('/main');
      }
    },

    toggleMode() {
      this.isLogin = !this.isLogin;
      this.errorMessage = '';
      this.successMessage = '';
      this.form.email = '';
    },

    clearMessages() {
      this.errorMessage = '';
      this.successMessage = '';
    },
    
    async handleSubmit() {
      this.loading = true;
      this.clearMessages();
      
      try {
        let response;
        
        if (this.isLogin) {
          response = await login(this.form.username, this.form.password);
        } else {
          response = await register(this.form.username, this.form.password, this.form.email);
        }
        
        console.log('API响应:', response);
        
        if (response && response.code === 200) {
          if (this.isLogin) {
            // 登录成功，保存用户名到本地存储
            localStorage.setItem('currentUser', this.form.username);
            
            this.successMessage = '登录成功，正在跳转...';
            
            // 延迟跳转
            setTimeout(() => {
              this.$router.replace('/main');
            }, 1000);
            
          } else {
            // 注册成功，切换到登录模式
            this.isLogin = true;
            this.successMessage = '注册成功，请登录';
            this.form.password = '';
            this.form.email = '';
          }
        } else {
          this.errorMessage = response?.msg || '操作失败';
        }
      } catch (error) {
        console.error('请求失败:', error);
        this.errorMessage = error.response?.data?.msg || '网络错误，请稍后再试';
      } finally {
        this.loading = false;
      }
    }
  }
};
</script>

<style scoped>
.login-container {
  min-height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  padding: 20px;
}

.login-card {
  background: white;
  border-radius: 20px;
  padding: 40px;
  width: 100%;
  max-width: 400px;
  box-shadow: 0 20px 60px rgba(0,0,0,0.15);
}

.login-header {
  text-align: center;
  margin-bottom: 30px;
}

.login-header h1 {
  font-size: 28px;
  color: #333;
  margin-bottom: 8px;
}

.login-header p {
  color: #666;
  font-size: 16px;
}

.login-form {
  margin-bottom: 20px;
}

.form-group {
  margin-bottom: 20px;
}

.form-group label {
  display: block;
  margin-bottom: 8px;
  color: #555;
  font-weight: 500;
}

.form-group input {
  width: 100%;
  padding: 12px 15px;
  border: 2px solid #e0e0e0;
  border-radius: 8px;
  font-size: 16px;
  transition: border-color 0.3s ease;
}

.form-group input:focus {
  outline: none;
  border-color: #667eea;
  box-shadow: 0 0 0 3px rgba(102, 126, 234, 0.1);
}

.form-group input:disabled {
  background: #f5f5f5;
  cursor: not-allowed;
}

.submit-btn {
  width: 100%;
  padding: 14px;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
  border: none;
  border-radius: 8px;
  font-size: 16px;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.3s ease;
}

.submit-btn:hover:not(:disabled) {
  transform: translateY(-1px);
  box-shadow: 0 4px 15px rgba(102, 126, 234, 0.4);
}

.submit-btn:disabled {
  opacity: 0.6;
  cursor: not-allowed;
  transform: none;
  box-shadow: none;
}

.error-message {
  margin-top: 15px;
  padding: 10px;
  background: #fee;
  border: 1px solid #fcc;
  border-radius: 6px;
  color: #c33;
  font-size: 14px;
  text-align: center;
}

.success-message {
  margin-top: 15px;
  padding: 10px;
  background: #efe;
  border: 1px solid #cfc;
  border-radius: 6px;
  color: #363;
  font-size: 14px;
  text-align: center;
}

.login-footer {
  text-align: center;
}

.login-footer p {
  color: #666;
  font-size: 14px;
}

.toggle-link {
  color: #667eea;
  text-decoration: none;
  font-weight: 500;
}

.toggle-link:hover {
  text-decoration: underline;
}

@media (max-width: 480px) {
  .login-card {
    padding: 30px 20px;
  }
  
  .login-header h1 {
    font-size: 24px;
  }
}
</style>
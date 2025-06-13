<template>
  <div class="user-center">
    <!-- 用户信息卡片 -->
    <div class="user-card">
      <div class="avatar-container">
        <div class="avatar" @click="selectAvatar">
          <img v-if="userInfo.avatar" :src="getAvatarUrl(userInfo.avatar)" alt="头像" />
          <div v-else class="default-avatar">{{ userInfo.username?.charAt(0).toUpperCase() }}</div>
          <div class="avatar-overlay">
            <i class="camera-icon">📷</i>
          </div>
        </div>
        <input 
          type="file" 
          ref="avatarInput" 
          @change="handleAvatarChange" 
          accept="image/*" 
          style="display: none"
        />
      </div>
      
      <div class="user-info">
        <h2 class="username">{{ userInfo.username }}</h2>
        <p class="email">{{ userInfo.email }}</p>
        <div class="user-stats">
          <div class="stat-item">
            <span class="stat-label">LLM调用余额</span>
            <span class="stat-value">{{ llmTokenBalance }}</span>
            <span class="stat-unit">tokens</span>
          </div>
        </div>
      </div>
      
      <div class="user-actions">
        <button @click="showEditModal = true" class="edit-btn">
          <i class="edit-icon">✏️</i>
          编辑资料
        </button>
        <button @click="refreshTokenBalance" class="refresh-btn">
          <i class="refresh-icon">🔄</i>
          刷新余额
        </button>
        <button @click="showRechargeModal = true" class="recharge-btn">
          <i class="money-icon">💰</i>
          充值Token
        </button>
      </div>
    </div>

    <!-- 编辑用户信息模态框 -->
    <div v-if="showEditModal" class="modal-overlay" @click="closeEditModal">
      <div class="modal-content" @click.stop>
        <div class="modal-header">
          <h3>编辑个人信息</h3>
          <button @click="closeEditModal" class="close-btn">×</button>
        </div>
        
        <form @submit.prevent="updateInfo" class="edit-form">
          <div class="form-group">
            <label>用户名</label>
            <input 
              type="text" 
              v-model="editForm.username" 
              required 
              placeholder="请输入用户名"
            />
          </div>
          
          <div class="form-group">
            <label>邮箱</label>
            <input 
              type="email" 
              v-model="editForm.email" 
              required 
              placeholder="请输入邮箱"
            />
          </div>
          
          <div class="form-group">
            <label>新密码（不修改请留空）</label>
            <input 
              type="password" 
              v-model="editForm.newPassword" 
              placeholder="请输入新密码"
            />
          </div>
          
          <div class="form-group">
            <label>确认新密码</label>
            <input 
              type="password" 
              v-model="editForm.confirmPassword" 
              placeholder="请确认新密码"
            />
          </div>
          
          <div class="form-actions">
            <button type="button" @click="closeEditModal" class="cancel-btn">取消</button>
            <button type="submit" :disabled="updating" class="save-btn">
              {{ updating ? '保存中...' : '保存' }}
            </button>
          </div>
        </form>
      </div>
    </div>

    <!-- 充值Token模态框 -->
    <div v-if="showRechargeModal" class="modal-overlay" @click="closeRechargeModal">
      <div class="modal-content" @click.stop>
        <div class="modal-header">
          <h3>充值LLM调用Token</h3>
          <button @click="closeRechargeModal" class="close-btn">×</button>
        </div>
        
        <div class="recharge-form">
          <div class="balance-info">
            <p>当前余额: <strong>{{ llmTokenBalance }} tokens</strong></p>
          </div>
          
          <div class="recharge-options">
            <div class="recharge-option" 
                 v-for="option in rechargeOptions" 
                 :key="option.tokens"
                 :class="{ active: selectedRecharge === option }"
                 @click="selectedRecharge = option">
              <div class="option-tokens">{{ option.tokens }} tokens</div>
              <div class="option-price">¥{{ option.price }}</div>
            </div>
          </div>
          
          <div class="form-actions">
            <button @click="closeRechargeModal" class="cancel-btn">取消</button>
            <button @click="handleRecharge" :disabled="!selectedRecharge || recharging" class="recharge-submit-btn">
              {{ recharging ? '充值中...' : '立即充值' }}
            </button>
          </div>
        </div>
      </div>
    </div>

    <!-- 消息提示 -->
    <div v-if="message" :class="['toast-message', messageType]">
      {{ message }}
    </div>
  </div>
</template>

<script>
import { getUserInfo, updateUserInfo, getTokenBalance, uploadAvatar, rechargeTokens } from '../services/apiService';

export default {
  name: 'UserCenter',
  data() {
    return {
      currentUser: '',
      userInfo: {
        username: '',
        email: '',
        avatar: ''
      },
      editForm: {
        username: '',
        email: '',
        newPassword: '',
        confirmPassword: ''
      },
      llmTokenBalance: 0,
      updating: false,
      message: '',
      messageType: 'success',
      showEditModal: false,
      showRechargeModal: false,
      uploadingAvatar: false,
      recharging: false,
      selectedRecharge: null,
      rechargeOptions: [
        { tokens: 1000, price: 10 },
        { tokens: 5000, price: 45 },
        { tokens: 10000, price: 80 },
        { tokens: 50000, price: 350 }
      ]
    };
  },
  async mounted() {
    this.currentUser = localStorage.getItem('currentUser') || '';
    await this.loadUserInfo();
    await this.loadTokenBalance();
  },
  methods: {
    async loadUserInfo() {
      try {
        if (!this.currentUser) {
          this.showMessage('未找到用户信息', 'error');
          return;
        }

        const response = await getUserInfo(this.currentUser);
        if (response && response.code === 200) {
          this.userInfo = {
            username: response.data.username || '',
            email: response.data.email || '',
            avatar: response.data.avatar || ''
          };
          this.resetEditForm();
        } else {
          this.showMessage(response?.msg || '获取用户信息失败', 'error');
        }
      } catch (error) {
        console.error('加载用户信息失败:', error);
        this.showMessage('加载用户信息失败', 'error');
      }
    },
    
    async loadTokenBalance() {
      try {
        if (!this.currentUser) return;

        const response = await getTokenBalance(this.currentUser);
        if (response && response.code === 200) {
          this.llmTokenBalance = response.data.balance || 0;
        } else {
          this.showMessage(response?.msg || '获取Token余额失败', 'error');
        }
      } catch (error) {
        console.error('加载Token余额失败:', error);
        this.showMessage('加载Token余额失败', 'error');
      }
    },
    
    async refreshTokenBalance() {
      await this.loadTokenBalance();
      this.showMessage('余额已刷新', 'success');
    },
    
    async handleRecharge() {
      if (!this.selectedRecharge) {
        this.showMessage('请选择充值套餐', 'error');
        return;
      }
      
      this.recharging = true;
      try {
        const response = await rechargeTokens(this.currentUser, this.selectedRecharge.tokens);
        if (response && response.code === 200) {
          this.showMessage(`成功充值 ${this.selectedRecharge.tokens} tokens`, 'success');
          await this.loadTokenBalance();
          this.closeRechargeModal();
        } else {
          this.showMessage(response?.msg || '充值失败', 'error');
        }
      } catch (error) {
        this.showMessage('充值失败', 'error');
        console.error('Error recharging:', error);
      } finally {
        this.recharging = false;
      }
    },
    
    closeEditModal() {
      this.showEditModal = false;
      this.resetEditForm();
    },
    
    closeRechargeModal() {
      this.showRechargeModal = false;
      this.selectedRecharge = null;
    },
    
    selectAvatar() {
      this.$refs.avatarInput.click();
    },
    
    async handleAvatarChange(event) {
      const file = event.target.files[0];
      if (!file) return;
      
      // 验证文件类型
      if (!file.type.startsWith('image/')) {
        this.showMessage('请选择图片文件', 'error');
        return;
      }
      
      // 验证文件大小（限制为2MB）
      if (file.size > 2 * 1024 * 1024) {
        this.showMessage('图片大小不能超过2MB', 'error');
        return;
      }
      
      this.uploadingAvatar = true;
      try {
        const response = await uploadAvatar(this.currentUser, file);
        if (response && response.code === 200) {
          this.userInfo.avatar = response.data.avatarUrl;
          this.showMessage('头像更新成功', 'success');
        } else {
          this.showMessage(response?.msg || '头像上传失败', 'error');
        }
      } catch (error) {
        this.showMessage('头像上传失败', 'error');
        console.error('Error uploading avatar:', error);
      } finally {
        this.uploadingAvatar = false;
      }
    },
    
    resetEditForm() {
      this.editForm = {
        username: this.userInfo.username || '',
        email: this.userInfo.email || '',
        newPassword: '',
        confirmPassword: ''
      };
    },
    
    async updateInfo() {
      // 验证密码
      if (this.editForm.newPassword && this.editForm.newPassword !== this.editForm.confirmPassword) {
        this.showMessage('两次输入的密码不一致', 'error');
        return;
      }
      
      // 验证密码强度
      if (this.editForm.newPassword && this.editForm.newPassword.length < 6) {
        this.showMessage('密码长度至少6位', 'error');
        return;
      }
      
      this.updating = true;
      try {
        const updateData = {
          oldUsername: this.currentUser,
          username: this.editForm.username,
          email: this.editForm.email
        };
        
        if (this.editForm.newPassword) {
          updateData.password = this.editForm.newPassword;
        }
        
        const response = await updateUserInfo(updateData);
        if (response && response.code === 200) {
          this.showMessage('信息更新成功', 'success');
          
          // 如果用户名发生变化，更新本地存储
          if (updateData.username !== this.currentUser) {
            localStorage.setItem('currentUser', updateData.username);
            this.currentUser = updateData.username;
          }
          
          await this.loadUserInfo();
          this.closeEditModal();
        } else {
          this.showMessage(response?.msg || '更新失败', 'error');
        }
      } catch (error) {
        this.showMessage(error.response?.data?.msg || '网络错误', 'error');
        console.error('Error updating user info:', error);
      } finally {
        this.updating = false;
      }
    },
    
    getAvatarUrl(avatar) {
      if (!avatar) return '';
      if (avatar.startsWith('http')) return avatar;
      return `http://localhost:8081${avatar}`;
    },
    
    showMessage(text, type = 'success') {
      this.message = text;
      this.messageType = type;
      setTimeout(() => {
        this.message = '';
      }, 3000);
    }
  }
};
</script>

<style scoped>
/* 保持原有样式不变，这里省略了样式代码 */
.user-center {
  max-width: 900px;
  margin: 0 auto;
  padding: 20px;
  min-height: 100vh;
  background: #f5f7fa;
}

.user-card {
  background: white;
  border-radius: 20px;
  padding: 40px;
  box-shadow: 0 8px 32px rgba(0,0,0,0.1);
  display: flex;
  align-items: center;
  gap: 40px;
  margin-bottom: 30px;
}

.avatar-container {
  position: relative;
}

.avatar {
  width: 120px;
  height: 120px;
  border-radius: 50%;
  overflow: hidden;
  cursor: pointer;
  position: relative;
  border: 4px solid #e0e0e0;
  transition: all 0.3s ease;
}

.avatar:hover {
  border-color: #667eea;
  transform: scale(1.05);
}

.avatar img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.default-avatar {
  width: 100%;
  height: 100%;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  display: flex;
  align-items: center;
  justify-content: center;
  color: white;
  font-size: 48px;
  font-weight: bold;
}

.avatar-overlay {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: rgba(0,0,0,0.5);
  display: flex;
  align-items: center;
  justify-content: center;
  opacity: 0;
  transition: opacity 0.3s ease;
}

.avatar:hover .avatar-overlay {
  opacity: 1;
}

.camera-icon {
  font-size: 24px;
}

.user-info {
  flex: 1;
}

.username {
  margin: 0 0 8px 0;
  font-size: 32px;
  font-weight: 700;
  color: #333;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
}

.email {
  margin: 0 0 20px 0;
  color: #666;
  font-size: 18px;
}

.user-stats {
  display: flex;
  gap: 30px;
}

.stat-item {
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 15px 25px;
  background: #f8f9fa;
  border-radius: 12px;
  min-width: 120px;
}

.stat-label {
  font-size: 14px;
  color: #666;
  margin-bottom: 5px;
}

.stat-value {
  font-size: 24px;
  font-weight: bold;
  color: #667eea;
}

.stat-unit {
  font-size: 12px;
  color: #999;
  margin-left: 4px;
}

.user-actions {
  display: flex;
  flex-direction: column;
  gap: 15px;
}

.edit-btn, .refresh-btn, .recharge-btn {
  padding: 12px 24px;
  border: none;
  border-radius: 10px;
  cursor: pointer;
  font-size: 16px;
  font-weight: 500;
  transition: all 0.3s ease;
  display: flex;
  align-items: center;
  gap: 8px;
  min-width: 140px;
  justify-content: center;
}

.edit-btn {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
}

.edit-btn:hover {
  transform: translateY(-2px);
  box-shadow: 0 4px 15px rgba(102, 126, 234, 0.4);
}

.refresh-btn {
  background: #f8f9fa;
  color: #333;
  border: 2px solid #e0e0e0;
}

.refresh-btn:hover {
  background: #e9ecef;
  border-color: #667eea;
}

.recharge-btn {
  background: linear-gradient(135deg, #28a745 0%, #20c997 100%);
  color: white;
}

.recharge-btn:hover {
  transform: translateY(-2px);
  box-shadow: 0 4px 15px rgba(40, 167, 69, 0.4);
}

/* 模态框样式 */
.modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: rgba(0,0,0,0.5);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 1000;
  animation: fadeIn 0.3s ease;
}

.modal-content {
  background: white;
  border-radius: 15px;
  width: 90%;
  max-width: 500px;
  max-height: 90vh;
  overflow-y: auto;
  animation: slideIn 0.3s ease;
}

.modal-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 25px 30px 15px;
  border-bottom: 1px solid #e0e0e0;
}

.modal-header h3 {
  margin: 0;
  font-size: 20px;
  color: #333;
}

.close-btn {
  background: none;
  border: none;
  font-size: 24px;
  cursor: pointer;
  color: #666;
  width: 30px;
  height: 30px;
  display: flex;
  align-items: center;
  justify-content: center;
  border-radius: 50%;
  transition: background 0.3s ease;
}

.close-btn:hover {
  background: #f0f0f0;
}

.edit-form, .recharge-form {
  padding: 25px 30px;
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

.form-actions {
  display: flex;
  gap: 15px;
  justify-content: flex-end;
  margin-top: 25px;
  padding-top: 20px;
  border-top: 1px solid #e0e0e0;
}

.cancel-btn, .save-btn, .recharge-submit-btn {
  padding: 12px 24px;
  border: none;
  border-radius: 8px;
  cursor: pointer;
  font-size: 16px;
  font-weight: 500;
  transition: all 0.3s ease;
}

.cancel-btn {
  background: #f8f9fa;
  color: #666;
  border: 2px solid #e0e0e0;
}

.cancel-btn:hover {
  background: #e9ecef;
}

.save-btn {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
}

.save-btn:hover {
  transform: translateY(-1px);
  box-shadow: 0 4px 15px rgba(102, 126, 234, 0.4);
}

.save-btn:disabled {
  opacity: 0.6;
  cursor: not-allowed;
  transform: none;
  box-shadow: none;
}

.balance-info {
  text-align: center;
  margin-bottom: 25px;
  padding: 15px;
  background: #f8f9fa;
  border-radius: 8px;
}

.recharge-options {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: 15px;
  margin-bottom: 25px;
}

.recharge-option {
  padding: 20px;
  border: 2px solid #e0e0e0;
  border-radius: 10px;
  text-align: center;
  cursor: pointer;
  transition: all 0.3s ease;
}

.recharge-option:hover {
  border-color: #667eea;
}

.recharge-option.active {
  border-color: #667eea;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
}

.option-tokens {
  font-size: 18px;
  font-weight: bold;
  margin-bottom: 5px;
}

.option-price {
  font-size: 16px;
  color: #666;
}

.recharge-option.active .option-price {
  color: #fff;
}

.recharge-submit-btn {
  background: linear-gradient(135deg, #28a745 0%, #20c997 100%);
  color: white;
}

.recharge-submit-btn:hover {
  transform: translateY(-1px);
  box-shadow: 0 4px 15px rgba(40, 167, 69, 0.4);
}

.recharge-submit-btn:disabled {
  opacity: 0.6;
  cursor: not-allowed;
  transform: none;
  box-shadow: none;
}

/* 消息提示 */
.toast-message {
  position: fixed;
  top: 20px;
  right: 20px;
  padding: 15px 25px;
  border-radius: 8px;
  font-weight: 500;
  z-index: 1001;
  animation: slideInRight 0.3s ease;
}

.toast-message.success {
  background: #4caf50;
  color: white;
}

.toast-message.error {
  background: #f44336;
  color: white;
}

/* 动画 */
@keyframes fadeIn {
  from { opacity: 0; }
  to { opacity: 1; }
}

@keyframes slideIn {
  from { 
    opacity: 0;
    transform: translateY(-50px) scale(0.95);
  }
  to { 
    opacity: 1;
    transform: translateY(0) scale(1);
  }
}

@keyframes slideInRight {
  from { 
    opacity: 0;
    transform: translateX(100px);
  }
  to { 
    opacity: 1;
    transform: translateX(0);
  }
}

/* 响应式设计 */
@media (max-width: 768px) {
  .user-card {
    flex-direction: column;
    text-align: center;
    gap: 25px;
    padding: 30px 20px;
  }
  
  .username {
    font-size: 28px;
  }
  
  .user-stats {
    justify-content: center;
  }
  
  .user-actions {
    flex-direction: row;
    justify-content: center;
  }
  
  .modal-content {
    width: 95%;
    margin: 20px;
  }
  
  .edit-form, .recharge-form {
    padding: 20px;
  }
  
  .form-actions {
    flex-direction: column;
  }
}
</style>
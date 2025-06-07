<template>
  <div class="chat-container">
    <div class="chat-header">
      <h1>面试助手</h1>
    </div>
    
    <div class="message-container" ref="messageContainer">
      <MessageItem 
        v-for="(msg, index) in messages" 
        :key="index" 
        :message="msg" 
      />
      
      <div v-if="loading" class="loading-indicator">
        <div class="dot"></div>
        <div class="dot"></div>
        <div class="dot"></div>
      </div>
    </div>
    
    <div class="input-container">
      <textarea 
        v-model="userInput" 
        placeholder="输入面试问题..." 
        @keydown.enter.exact.prevent="sendMessage"
        :disabled="loading"
        rows="3"
        ref="inputField"
      ></textarea>
      <button 
        @click="sendMessage" 
        :disabled="loading || !userInput.trim()"
        class="send-button"
      >
        <svg viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
          <path d="M22 2L11 13" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"></path>
          <path d="M22 2L15 22L11 13L2 9L22 2Z" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"></path>
        </svg>
      </button>
    </div>
  </div>
</template>

<script>
import MessageItem from './MessageItem.vue';
import { sendChatMessage } from '../services/apiService';

export default {
  name: 'ChatInterface',
  components: {
    MessageItem
  },
  data() {
    return {
      messages: [
        {
          role: 'assistant',
          content: '你好！我是面试助手，可以帮你准备计算机专业的面试问题。请问有什么可以帮助你的？'
        }
      ],
      userInput: '',
      loading: false
    };
  },
  methods: {
    async sendMessage() {
      if (!this.userInput.trim() || this.loading) return;
      
      const userMessage = this.userInput.trim();
      this.messages.push({
        role: 'user',
        content: userMessage
      });
      
      this.userInput = '';
      this.loading = true;
      this.scrollToBottom();
      
      try {
        // 调用后端API
        const response = await sendChatMessage(userMessage);
        if (response.code === 200 || response.code === 0) {
          this.messages.push({
            role: 'assistant',
            content: response.msg || response.message || '收到回复'
          });
        } else {
          // 错误情况，显示错误信息
          this.messages.push({
            role: 'assistant',
            content: response.msg || '服务器返回错误',
            error: true
          });
        }
        
      } catch (error) {
        console.error('发送消息失败:', error);
        // 添加错误提示
        this.messages.push({
          role: 'assistant',
          content: '抱歉，发生了错误，请稍后再试。',
          error: true
        });
      } finally {
        this.loading = false;
        this.scrollToBottom();
        this.$nextTick(() => {
          this.$refs.inputField.focus();
        });
      }
    },
    
    scrollToBottom() {
      this.$nextTick(() => {
        const container = this.$refs.messageContainer;
        container.scrollTop = container.scrollHeight;
      });
    }
  },
  mounted() {
    // 初始聚焦到输入框
    this.$refs.inputField.focus();
  }
};
</script>

<style scoped>
.chat-container {
  display: flex;
  flex-direction: column;
  height: 100vh;
  max-width: 800px;
  margin: 0 auto;
  background-color: #ffffff;
  box-shadow: 0 0 20px rgba(0,0,0,0.1);
}

.chat-header {
  padding: 16px;
  text-align: center;
  border-bottom: 1px solid #e6e6e6;
}

.chat-header h1 {
  margin: 0;
  font-size: 1.5rem;
  color: #333;
}

.message-container {
  flex: 1;
  overflow-y: auto;
  padding: 20px;
  display: flex;
  flex-direction: column;
}

.input-container {
  display: flex;
  align-items: flex-end;
  padding: 16px;
  border-top: 1px solid #e6e6e6;
  background-color: #fff;
}

textarea {
  flex: 1;
  padding: 12px 16px;
  border: 1px solid #d9d9d9;
  border-radius: 8px;
  resize: none;
  font-family: inherit;
  font-size: 16px;
  line-height: 1.5;
  outline: none;
  transition: border-color 0.2s;
}

textarea:focus {
  border-color: #8c8c8c;
}

.send-button {
  width: 40px;
  height: 40px;
  margin-left: 8px;
  display: flex;
  align-items: center;
  justify-content: center;
  background-color: #2e7d32;
  border: none;
  border-radius: 8px;
  color: white;
  cursor: pointer;
  transition: background-color 0.2s;
}

.send-button:hover {
  background-color: #1b5e20;
}

.send-button:disabled {
  background-color: #c8e6c9;
  cursor: not-allowed;
}

.send-button svg {
  width: 20px;
  height: 20px;
}

.loading-indicator {
  align-self: flex-start;
  display: flex;
  gap: 4px;
  padding: 12px 16px;
  background-color: #f0f2f5;
  border-radius: 16px;
  margin-top: 10px;
}

.dot {
  width: 8px;
  height: 8px;
  border-radius: 50%;
  background-color: #8c8c8c;
  animation: pulse 1.5s infinite;
}

.dot:nth-child(2) {
  animation-delay: 0.3s;
}

.dot:nth-child(3) {
  animation-delay: 0.6s;
}

@keyframes pulse {
  0%, 100% {
    opacity: 0.5;
    transform: scale(0.8);
  }
  50% {
    opacity: 1;
    transform: scale(1);
  }
}

@media (max-width: 768px) {
  .chat-container {
    height: 100vh;
    max-width: 100%;
    border-radius: 0;
    box-shadow: none;
  }
}
</style>
<template>
  <div :class="['message', messageClass]">
    <div class="avatar">
      <div v-if="message.role === 'assistant'" class="assistant-avatar">A</div>
      <div v-else class="user-avatar">U</div>
    </div>
    <div class="content">
      <div class="markdown-content" v-html="formattedContent"></div>
    </div>
  </div>
</template>

<script>
import { marked } from 'marked';

export default {
  name: 'MessageItem',
  props: {
    message: {
      type: Object,
      required: true
    }
  },
  computed: {
    messageClass() {
      return {
        'assistant-message': this.message.role === 'assistant',
        'user-message': this.message.role === 'user',
        'error-message': this.message.error
      };
    },
    formattedContent() {
      // 仅对助手消息应用Markdown
      if (this.message.role === 'assistant') {
        return marked(this.message.content);
      }
      return this.message.content;
    }
  }
};
</script>

<style scoped>
.message {
  display: flex;
  margin-bottom: 20px;
  gap: 12px;
}

.avatar {
  flex-shrink: 0;
  width: 36px;
  height: 36px;
  border-radius: 4px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-weight: bold;
}

.assistant-avatar {
  background-color: #10a37f;
  color: white;
}

.user-avatar {
  background-color: #e9e9e9;
  color: #333;
}

.content {
  flex: 1;
  padding: 12px 16px;
  background-color: #f7f7f8;
  border-radius: 8px;
  line-height: 1.5;
  font-size: 16px;
}

.user-message .content {
  background-color: #e9f5ff;
}

.error-message .content {
  background-color: #ffebee;
}

/* 应用深度选择器来影响Markdown渲染内容的样式 */
:deep(.markdown-content pre) {
  background-color: rgba(0, 0, 0, 0.05);
  padding: 12px;
  border-radius: 6px;
  overflow-x: auto;
}

:deep(.markdown-content code) {
  font-family: 'SFMono-Regular', Consolas, 'Liberation Mono', Menlo, Courier, monospace;
  font-size: 14px;
}

:deep(.markdown-content p) {
  margin: 8px 0;
}

:deep(.markdown-content p:first-child) {
  margin-top: 0;
}

:deep(.markdown-content p:last-child) {
  margin-bottom: 0;
}
</style>
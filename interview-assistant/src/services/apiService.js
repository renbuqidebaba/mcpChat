// src/services/apiService.js
import axios from "axios";

// 后端API基础URL，如果后端运行在不同端口，请相应修改
const BASE_URL = "http://localhost:8081/api";

const apiClient = axios.create({
  baseURL: BASE_URL,
  headers: {
    "Content-Type": "application/json",
  },
  timeout: 30000, // 超时时间30秒
});

/**
 * 发送聊天消息到后端
 * @param {string} message - 用户输入的消息
 * @returns {Promise} - 返回后端响应
 */
export async function sendChatMessage(message) {
  try {
    const response = await apiClient.post("/chat", { message });
    return response.data;
  } catch (error) {
    console.error("API调用失败:", error);
    throw error;
  }
}

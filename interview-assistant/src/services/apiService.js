import axios from "axios";

// 后端API基础URL
const BASE_URL = "http://localhost:8081";

const apiClient = axios.create({
  baseURL: BASE_URL,
  headers: {
    "Content-Type": "application/json",
  },
  timeout: 30000,
});

/**
 * 用户登录
 * @param {string} username - 用户名
 * @param {string} password - 密码
 */
export async function login(username, password) {
  try {
    const response = await apiClient.post("/user/login", { username, password });
    return response.data;
  } catch (error) {
    console.error("登录失败:", error);
    throw error;
  }
}

/**
 * 用户注册
 * @param {string} username - 用户名
 * @param {string} password - 密码
 * @param {string} email - 邮箱
 */
export async function register(username, password, email) {
  try {
    const response = await apiClient.post("/user/register", { username, password, email });
    return response.data;
  } catch (error) {
    console.error("注册失败:", error);
    throw error;
  }
}

/**
 * 获取用户信息
 * @param {string} username - 用户名
 */
export async function getUserInfo(username) {
  try {
    const response = await apiClient.get(`/user/info/${username}`);
    return response.data;
  } catch (error) {
    console.error("获取用户信息失败:", error);
    throw error;
  }
}

/**
 * 更新用户信息
 * @param {Object} userInfo - 用户信息
 */
export async function updateUserInfo(userInfo) {
  try {
    const response = await apiClient.put("/user/update", userInfo);
    return response.data;
  } catch (error) {
    console.error("更新用户信息失败:", error);
    throw error;
  }
}

/**
 * 上传用户头像
 * @param {string} username - 用户名
 * @param {File} file - 头像文件
 */
export async function uploadAvatar(username, file) {
  try {
    const formData = new FormData();
    formData.append('username', username);
    formData.append('avatar', file);
    
    const response = await apiClient.post("/user/avatar", formData, {
      headers: {
        'Content-Type': 'multipart/form-data'
      }
    });
    return response.data;
  } catch (error) {
    console.error("头像上传失败:", error);
    throw error;
  }
}

/**
 * 获取用户LLM调用余额
 * @param {string} username - 用户名
 */
export async function getTokenBalance(username) {
  try {
    const response = await apiClient.get(`/user/token-balance/${username}`);
    return response.data;
  } catch (error) {
    console.error("获取token余额失败:", error);
    throw error;
  }
}

/**
 * 充值LLM token
 * @param {string} username - 用户名
 * @param {number} amount - 充值token数量
 */
export async function rechargeTokens(username, amount) {
  try {
    const response = await apiClient.post("/user/recharge", { username, amount });
    return response.data;
  } catch (error) {
    console.error("充值失败:", error);
    throw error;
  }
}

/**
 * 发送聊天消息到后端
 * @param {string} username - 用户名
 * @param {string} message - 用户输入的消息
 * @returns {Promise} - 返回后端响应
 */
export async function sendChatMessage(username, message) {
  try {
    const response = await apiClient.post("/chat", { username, message });
    return response.data;
  } catch (error) {
    console.error("API调用失败:", error);
    throw error;
  }
}
import axios from 'axios';

const apiClient = axios.create({
  baseURL: process.env.NEXT_PUBLIC_API_BASE_URL || 'http://localhost:8080', // 環境変数を使用
  headers: {
    'Content-Type': 'application/json',
  },
});

export default apiClient;

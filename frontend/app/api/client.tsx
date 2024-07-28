import axios from 'axios';

const apiClient = axios.create({
  baseURL: 'http://localhost:8080', // バックエンドのベースURLに置き換えてください
  headers: {
    'Content-Type': 'application/json',
  },
});

export default apiClient;

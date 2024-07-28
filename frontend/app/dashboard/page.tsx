// frontend/app/dashboard/page.tsx
'use client';

import React, { useEffect, useState } from 'react';
import apiClient from '../api/client';

export default function DashboardPage() {
  const [userData, setUserData] = useState<any>(null);
  const [loading, setLoading] = useState<boolean>(true);

  useEffect(() => {
    // ユーザーデータを取得するためのAPIリクエストを送信
    const fetchUserData = async () => {
      try {
        const response = await apiClient.get('/users/me');
        setUserData(response.data);
      } catch (error) {
        console.error('Error fetching user data:', error);
      } finally {
        setLoading(false);
      }
    };

    fetchUserData();
  }, []);

  if (loading) {
    return <div>Loading...</div>;
  }

  if (!userData) {
    return <div>Error loading user data.</div>;
  }

  return (
    <div className="p-4">
      <h1 className="text-2xl font-bold mb-4">ダッシュボード</h1>
      <div className="bg-white p-6 rounded-lg shadow-md">
        <h2 className="text-xl font-semibold">こんにちは、{userData.name}さん！</h2>
        <p className="mt-2">Email: {userData.email}</p>
        {/* 他のユーザー情報やダッシュボードコンテンツをここに追加 */}
      </div>
    </div>
  );
}

// DashboardPage.tsx

'use client';

import React, { useEffect, useState } from 'react';
import apiClient from '../api/client';
import withAuth from '../../components/hocs/withAuth';
import PageTitle from '../../components/atoms/PageTitle'; // PageTitleをインポート

const DashboardPage: React.FC = () => {
  const [userData, setUserData] = useState<any>(null);
  const [loading, setLoading] = useState<boolean>(true);

  useEffect(() => {
    const fetchUserData = async () => {
      try {
        const response = await apiClient.get('/api/users/me', {
          headers: {
            Authorization: `Bearer ${localStorage.getItem('token')}`,
          },
        });
        setUserData(response.data.user);
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
      <PageTitle title="ダッシュボード" />
      <div className="bg-white p-6 rounded-lg shadow-md">
        <h2 className="text-xl font-semibold">{userData.name} 様</h2>
        <p className="mt-2">Email: {userData.email}</p>
      </div>
    </div>
  );
}

export default withAuth(DashboardPage);

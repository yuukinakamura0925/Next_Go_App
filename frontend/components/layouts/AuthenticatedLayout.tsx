// components/layouts/AuthenticatedLayout.tsx

'use client';

import { ReactNode, useState } from 'react';
import Sidebar from '../organisms/Sidebar';

// 認証されたレイアウトコンポーネント
export default function AuthenticatedLayout({ children }: { children: ReactNode }) {
  const [isSidebarOpen, setIsSidebarOpen] = useState(true); // サイドバーの開閉状態を管理

  const handleLogout = () => {
    localStorage.removeItem('token');
    window.dispatchEvent(new Event('storage'));
    window.location.href = '/auth/sign_in';
  };

  // サイドバーの開閉を切り替える関数
  const toggleSidebar = () => {
    setIsSidebarOpen((prev) => !prev);
  };

  return (
    <div className="flex flex-col min-h-screen">
    {/* ヘッダー */}
    <header className="bg-blue-600 text-white p-4">
      <div className="flex justify-between items-center" style={{ paddingLeft: '20px', paddingRight: '20px' }}>
        <h1 className="text-xl font-bold">My App</h1>
        <nav>
          <ul className="flex space-x-4">
            <li>
              <button onClick={handleLogout}>ログアウト</button>
            </li>
          </ul>
        </nav>
      </div>
    </header>

    {/* メインコンテンツ */}
    <div className="flex flex-1 m-8">
      {/* サイドバー */}
      <Sidebar isOpen={isSidebarOpen} toggleSidebar={toggleSidebar} />

      {/* 子コンポーネント */}
      <main className="bg-white text-black p-4 flex-1 ml-8 overflow-auto">
        {children}
      </main>
    </div>
  </div>
  );
}
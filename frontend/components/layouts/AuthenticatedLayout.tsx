'use client';

import { ReactNode, useState } from 'react';

// アイコンのスタイル用コンポーネント（仮のアイコン用）
const DashboardIcon = () => <span className="material-icons">D</span>;
const BookIcon = () => <span className="material-icons">L</span>;

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
        <div className="container mx-auto flex justify-between items-center">
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
        <div className={`transition-all duration-300 bg-white text-black ${isSidebarOpen ? 'w-64' : 'w-0'} relative`}>
          {/* 開閉トリガー */}
          {isSidebarOpen ? (
            <button
              onClick={toggleSidebar}
              className="absolute top-0 right-0 mt-4 mr-4 bg-gray-700 text-white rounded-full w-8 h-8 flex items-center justify-center cursor-pointer"
            >
              ←
            </button>
          ) : (
            <button
              onClick={toggleSidebar}
              className="absolute top-0 -right-4 mt-4 transform translate-x-full bg-gray-700 text-white rounded-full w-8 h-8 flex items-center justify-center cursor-pointer"
            >
              →
            </button>
          )}

          {/* サイドバーの内容 */}
          {isSidebarOpen && (
            <nav className="p-4">
              <ul className="space-y-4">
                <li className="flex items-center">
                  <DashboardIcon />
                  <span className="ml-2">ダッシュボード</span>
                </li>
                <li className="flex items-center">
                  <BookIcon />
                  <span className="ml-2">本追加</span>
                </li>
              </ul>
            </nav>
          )}
        </div>

        {/* 子コンポーネント */}
        <main className="bg-white text-black p-4 flex-1 ml-8 overflow-auto">
          {children}
        </main>
      </div>
    </div>
  );
}

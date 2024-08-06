// components/Sidebar.tsx

'use client';

import React from 'react';

// アイコンのスタイル用コンポーネント（仮のアイコン用）
const DashboardIcon = () => <span className="material-icons">D</span>;
const BookIcon = () => <span className="material-icons">M</span>;

interface SidebarProps {
  isOpen: boolean;
  toggleSidebar: () => void;
}

const Sidebar: React.FC<SidebarProps> = ({ isOpen, toggleSidebar }) => {
  return (
    <div className={`transition-all duration-300 bg-white text-black ${isOpen ? 'w-64' : 'w-0'} relative pt-10 pb-10`}>
      {/* 開閉トリガー */}
      {isOpen ? (
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
      {isOpen && (
        <nav className="p-4">
          <ul className="space-y-8">
            <li className="flex items-center">
              <DashboardIcon />
              <a href="/dashboard" className="ml-2">ダッシュボード</a>
            </li>
            <li className="flex items-center">
              <BookIcon />
              <a href="/menu_category" className="ml-2">メニューカテゴリー</a>
            </li>
            <li>
            </li>
          </ul>
        </nav>
      )}
    </div>
  );
};

export default Sidebar;

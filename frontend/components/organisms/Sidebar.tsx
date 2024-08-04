'use client';

import React from 'react';

const Sidebar: React.FC = () => {
  return (
    <div className="bg-gray-800 text-white w-64 p-4">
      <ul>
        <li className="py-2"><a href="/dashboard">ダッシュボード</a></li>
        <li className="py-2"><a href="/books">本追加</a></li>
        {/* 他のリンクを追加 */}
      </ul>
    </div>
  );
};

export default Sidebar;

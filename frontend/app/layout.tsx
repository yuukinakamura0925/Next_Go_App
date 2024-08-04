'use client'; // クライアントコンポーネントとしてマーク

import './globals.css';
import { ReactNode, useEffect, useState } from 'react';
import { useAuth } from '../hooks/useAuth';
import AuthenticatedLayout from '../components/layouts/AuthenticatedLayout'; // AuthenticatedLayoutをインポート

// ルートレイアウトコンポーネント
export default function RootLayout({ children }: { children: ReactNode }) {
  const auth = useAuth(); // カスタムフックuseAuthを使用して認証状態を取得
  const [isAuthenticated, setIsAuthenticated] = useState(auth.isAuthenticated); // 初期認証状態を設定

  // localStorageのトークンの有無を確認し、認証状態を更新する関数
  const handleStorageChange = () => {
    const token = localStorage.getItem('token');
    setIsAuthenticated(!!token); // トークンが存在すればtrue、存在しなければfalse
  };

  useEffect(() => {
    // ページ読み込み時にトークンの有無を確認
    handleStorageChange();

    // ストレージの変更を監視
    window.addEventListener('storage', handleStorageChange);

    // クリーンアップ関数
    return () => {
      window.removeEventListener('storage', handleStorageChange);
    };
  }, []);

  return (
    <html lang="en">
      <head>
        <title>My App</title>
      </head>
      <body>
        {isAuthenticated ? (
          <AuthenticatedLayout>
            {children}
          </AuthenticatedLayout>
        ) : (
          <UnauthenticatedLayout>
            {children}
          </UnauthenticatedLayout>
        )}
      </body>
    </html>
  );
}

function UnauthenticatedLayout({ children }: { children: ReactNode }) {
  return (
    <div>
      <nav className="bg-blue-600 p-4 text-white">
        <ul className="flex space-x-4">
          <li><a href="/auth/sign_up">新規登録</a></li>
          <li><a href="/auth/sign_in">サインイン</a></li>
        </ul>
      </nav>
      <main className="p-4">{children}</main>
    </div>
  );
}

'use client'; // クライアントコンポーネントとしてマーク

import { useEffect, useState } from 'react';
import jwt_decode, { JwtPayload } from 'jwt-decode';

// 認証状態を保持するためのインターフェース
interface AuthState {
  isAuthenticated: boolean; // 認証されているかどうか
  user: JwtPayload | null; // デコードされたトークンのペイロード
  loading: boolean; // 認証状態を読み込んでいるかどうか
}

// 認証状態を取得するカスタムフック
export const useAuth = (): AuthState => {
  // 認証状態を管理するステートを定義
  const [auth, setAuth] = useState<AuthState>({ isAuthenticated: false, user: null, loading: true });

  useEffect(() => {
    // ローカルストレージからトークンを取得
    const token = localStorage.getItem('token');
    if (token) {
      try {
        // トークンをデコード
        const decoded = jwt_decode<JwtPayload>(token);
        // 認証状態を更新（認証されている）
        setAuth({ isAuthenticated: true, user: decoded, loading: false });
      } catch (error) {
        // トークンのデコードに失敗した場合のエラーハンドリング
        console.error('Token decode failed', error);
        localStorage.removeItem('token'); // トークンを削除
        // 認証状態を更新（認証されていない）
        setAuth({ isAuthenticated: false, user: null, loading: false });
      }
    } else {
      // トークンが存在しない場合の処理
      setAuth({ isAuthenticated: false, user: null, loading: false });
    }
  }, []);

  return auth; // 認証状態を返す
};

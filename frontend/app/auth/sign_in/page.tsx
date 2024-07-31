'use client';

import React, { useState } from 'react';
import { useRouter } from 'next/navigation'; 
import apiClient from '../../api/client';
import AuthForm from '../../../components/organisms/AuthForm';
import FlashMessage from '../../../components/atoms/FlashMessage';

export default function SignInPage() {
  const [email, setEmail] = useState<string>('');
  const [password, setPassword] = useState<string>('');
  const [message, setMessage] = useState<string | null>(null);
  const [isError, setIsError] = useState<boolean>(false);
  const router = useRouter();

  const handleSignIn = async (e: React.FormEvent) => {
    e.preventDefault();
    setMessage(null);
    try {
      const response = await apiClient.post('/users/sign_in', {
        email,
        password,
      });

      if (response.status === 200) {
        localStorage.setItem('token', response.data.token); // トークンを保存
        window.dispatchEvent(new Event('storage')); // ストレージイベントをトリガー
        setMessage('Sign in successful!');
        setIsError(false);
        router.push('/dashboard'); // サインイン成功時にリダイレクト
      } else {
        setMessage('Sign in failed!');
        setIsError(true);
      }
    } catch (error: any) {
      console.error("Sign in error:", error);
      setMessage(error.response?.data?.message || 'Sign in failed!');
      setIsError(true);
    }
  };

  return (
    <div className="flex flex-col items-center justify-center min-h-screen bg-gray-100">
      <div className="w-full max-w-md p-8 bg-white rounded shadow-md">
        <AuthForm
          title="サインイン"
          fields={[
            { id: 'email', label: 'Email', type: 'email', value: email, onChange: (e) => setEmail(e.target.value) },
            { id: 'password', label: 'パスワード', type: 'password', value: password, onChange: (e) => setPassword(e.target.value) },
          ]}
          onSubmit={handleSignIn}
          buttonText="Sign In"
        />
        <FlashMessage message={message} isError={isError} />
      </div>
    </div>
  );
}

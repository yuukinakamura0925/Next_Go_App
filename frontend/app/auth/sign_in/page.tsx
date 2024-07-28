'use client';

import { useState } from 'react';
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
        setMessage('Sign in successful!');
        setIsError(false);
        router.push('/dashboard'); // サインイン成功時にリダイレクト
      } else {
        setMessage('Sign in failed!');
        setIsError(true);
      }
    } catch (error: any) {
      setMessage(error.response?.data?.message || 'Sign in failed!');
      setIsError(true);
      console.error('Sign in failed:', error);
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

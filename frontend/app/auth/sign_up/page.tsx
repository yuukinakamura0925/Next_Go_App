'use client';

import { useState } from 'react';
import apiClient from '../../api/client';
import AuthForm from '../../../components/organisms/AuthForm';
import FlashMessage from '@/components/atoms/FlashMessage';

export default function SignupPage() {
  const [name, setName] = useState<string>('');
  const [email, setEmail] = useState<string>('');
  const [password, setPassword] = useState<string>('');
  const [message, setMessage] = useState<string | null>(null);
  const [isError, setIsError] = useState<boolean>(false);

  const handleSignup = async (e: React.FormEvent) => {
    e.preventDefault();
    setMessage(null);
    try {
      const response = await apiClient.post('/api/auth/users/sign_up', {
        name,
        email,
        password,
      });

      if (response.status === 201) {
        setMessage('Signup successful!');
        setIsError(false);
      } else {
        setMessage('Signup failed!');
        setIsError(true);
      }
    } catch (error: any) {
      setMessage(error.response?.data?.message || 'Signup failed!');
      setIsError(true);
    }
  };

  return (
    <div className="flex flex-col items-center justify-center min-h-screen bg-gray-100">
      <div className="w-full max-w-md p-8 bg-white rounded shadow-md">
        <AuthForm
          title="ユーザー登録"
          fields={[
            { id: 'name', label: '名前', type: 'text', value: name, onChange: (e) => setName(e.target.value) },
            { id: 'email', label: 'Email', type: 'email', value: email, onChange: (e) => setEmail(e.target.value) },
            { id: 'password', label: 'パスワード', type: 'password', value: password, onChange: (e) => setPassword(e.target.value) },
          ]}
          onSubmit={handleSignup}
          buttonText="Sign Up"
        />
        <FlashMessage message={message} isError={isError} />
      </div>
    </div>
  );
}

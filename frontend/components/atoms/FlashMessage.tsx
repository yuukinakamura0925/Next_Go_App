import React from 'react';

interface FlashMessageProps {
  message: string | null;
  isError: boolean;
}

const FlashMessage: React.FC<FlashMessageProps> = ({ message, isError }) => {
  if (!message) return null;

  return (
    <div className={`p-4 mt-4 text-sm rounded ${isError ? 'bg-red-100 text-red-700' : 'bg-green-100 text-green-700'}`}>
      {message}
    </div>
  );
};

export default FlashMessage;

import React from 'react';
import FormField from '../molecules/FormField';
import Button from '../atoms/Button';

interface AuthFormProps {
  title: string;
  fields: { id: string; label: string; type: string; value: string; onChange: (e: React.ChangeEvent<HTMLInputElement>) => void; }[];
  onSubmit: (e: React.FormEvent) => void;
  buttonText: string;
}

const AuthForm: React.FC<AuthFormProps> = ({ title, fields, onSubmit, buttonText }) => {
  return (
    <div className="bg-white p-6 rounded-lg w-full max-w-md">
      <h1 className="text-2xl font-bold mb-4">{title}</h1>
      <form onSubmit={onSubmit} className="space-y-4">
        {fields.map(field => (
          <FormField
            key={field.id}
            id={field.id}
            label={field.label}
            type={field.type}
            value={field.value}
            onChange={field.onChange}
          />
        ))}
        <Button type="submit">{buttonText}</Button>
      </form>
    </div>
  );
};

export default AuthForm;

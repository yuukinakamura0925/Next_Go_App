import React from 'react';
import InputField from '../atoms/InputField';

interface FormFieldProps {
  id: string;
  label: string;
  type: string;
  value: string;
  onChange: (e: React.ChangeEvent<HTMLInputElement>) => void;
}

const FormField: React.FC<FormFieldProps> = ({ id, label, type, value, onChange }) => {
  return <InputField
          id={id}
          label={label}
          type={type}
          value={value}
          onChange={onChange}
        />;
};

export default FormField;

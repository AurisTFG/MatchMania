import {
  FormControl,
  InputLabel,
  MenuItem,
  Select as MuiSelect,
} from '@mui/material';
import { useState } from 'react';
import { useFieldContext } from '../../../hooks/form/useAppForm';
import FormErrors from '../Helpers/FormErrors';

type SelectFieldProps = {
  label: string;
  options: { value: string | number; label: string }[];
};

export default function Select({ label, options }: SelectFieldProps) {
  const field = useFieldContext<string | number>();
  const [isOpen, setIsOpen] = useState(false);

  const errorMessages = field.state.meta.errors.map(
    (error: { message: string }) => error.message,
  );

  const enableLabelLayout = isOpen || !!field.state.value;

  return (
    <FormControl
      fullWidth
      margin="normal"
      error={errorMessages.length > 0}
    >
      <InputLabel>{label}</InputLabel>
      <MuiSelect
        label={enableLabelLayout ? label : ''}
        value={field.state.value}
        onChange={(e) => {
          field.handleChange(e.target.value);
        }}
        onOpen={() => {
          setIsOpen(true);
        }}
        onClose={() => {
          setIsOpen(false);
        }}
        displayEmpty
      >
        {options.map((option) => (
          <MenuItem
            key={option.value}
            value={option.value}
          >
            {option.label}
          </MenuItem>
        ))}
      </MuiSelect>
      <FormErrors messages={errorMessages} />
    </FormControl>
  );
}

import { TextField as MuiTextField } from '@mui/material';
import { useFieldContext } from 'hooks/form/useAppForm';
import FieldErrorText from './FieldErrorText';

type TextFieldProps = {
  label: string;
  type?: string;
  placeholder?: string;
};

export default function Text({
  label,
  type = 'text',
  placeholder,
}: TextFieldProps) {
  const field = useFieldContext<string>();

  const errorMessages = field.state.meta.errors.map(
    (error: { message: string }) => error.message,
  );

  return (
    <>
      <MuiTextField
        label={label}
        placeholder={placeholder ?? label}
        type={type}
        value={field.state.value}
        onChange={(e) => {
          field.handleChange(e.target.value);
        }}
        error={errorMessages.length > 0}
        variant="outlined"
        margin="normal"
        fullWidth
        sx={{
          color: '#ffffff',
          '&:-webkit-autofill': {
            WebkitBoxShadow: '0 0 0 100px #307ECC inset',
            WebkitTextFillColor: 'ffffff',
          },
        }}
      />

      <FieldErrorText messages={errorMessages} />
    </>
  );
}

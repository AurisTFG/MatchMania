import { SxProps, TextField as MuiTextField, Theme } from '@mui/material';
import { useFieldContext } from 'hooks/form/useAppForm';
import FormErrors from '../Helpers/FormErrors';

type TextFieldProps = {
  label: string;
  type?: string;
  placeholder?: string;
  sx?: SxProps<Theme>;
};

export default function Text({
  label,
  type = 'text',
  placeholder,
  sx,
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
        sx={sx}
        margin="normal"
        fullWidth
      />
      <FormErrors messages={errorMessages} />
    </>
  );
}

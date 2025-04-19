import { TextField as MuiTextField } from '@mui/material';
import { useFieldContext } from '../../../hooks/form/useAppForm';
import FormErrors from '../Helpers/FormErrors';

export default function TextField({
  label,
  type = 'text',
}: {
  label: string;
  type?: string;
}) {
  const field = useFieldContext<string>();

  const errorMessages = field.state.meta.errors.map(
    (error: { message: string }) => error.message,
  );

  return (
    <>
      <MuiTextField
        label={label}
        placeholder={label}
        type={type}
        value={field.state.value}
        onChange={(e) => {
          field.handleChange(e.target.value);
        }}
        error={errorMessages.length > 0}
        fullWidth
        margin="normal"
      />
      <FormErrors messages={errorMessages} />
    </>
  );
}

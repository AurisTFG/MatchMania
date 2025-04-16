import { TextField as MuiTextField } from '@mui/material';
import { useFieldContext } from '../../../hooks/form/useAppForm';

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
      {errorMessages.length > 0 && (
        <div
          style={{
            marginTop: '-5px',
            marginBottom: '10px',
            width: '100%',
            textAlign: 'left',
          }}
        >
          {errorMessages.map((message, index) => (
            <div
              key={index}
              style={{
                color: '#d32f2f',
                fontSize: '0.75rem',
                marginLeft: '14px',
              }}
            >
              {message}
            </div>
          ))}
        </div>
      )}
    </>
  );
}

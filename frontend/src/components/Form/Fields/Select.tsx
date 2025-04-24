import {
  FormControl,
  InputLabel,
  MenuItem,
  Select as MuiSelect,
} from '@mui/material';
import { SELECT_OPTIONS } from '../../../constants/selectOptions';
import { useFieldContext } from '../../../hooks/form/useAppForm';
import FormErrors from '../Helpers/FormErrors';

type SelectFieldProps = {
  label: string;
  options: { key: string; value: string }[];
};

export default function Select({ label, options }: SelectFieldProps) {
  const field = useFieldContext<string>();

  const errorMessages = field.state.meta.errors.map(
    (error: { message: string }) => error.message,
  );

  const optionsWithEmpty = [SELECT_OPTIONS.NOT_SELECTED, ...options];

  return (
    <FormControl
      fullWidth
      margin="normal"
      error={errorMessages.length > 0}
    >
      <InputLabel>{label}</InputLabel>
      <MuiSelect
        label={label}
        value={field.state.value}
        onChange={(e) => {
          field.handleChange(e.target.value);
        }}
      >
        {optionsWithEmpty.map((option) => (
          <MenuItem
            key={option.key}
            value={option.key}
          >
            {option.value}
          </MenuItem>
        ))}
      </MuiSelect>
      <FormErrors
        marginTop="0.18rem"
        messages={errorMessages}
      />
    </FormControl>
  );
}

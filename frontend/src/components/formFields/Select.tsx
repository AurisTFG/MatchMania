import { Autocomplete, TextField } from '@mui/material';
import { SELECT_OPTIONS } from 'constants/selectOptions';
import { useFieldContext } from 'hooks/form/useAppForm';
import FieldErrorText from './FieldErrorText';

type SelectFieldProps = {
  label: string;
  options: { key: string; value: string }[];
  notSelectedOption?: { key: string; value: string };
  renderOptionText?: (option: {
    key: string;
    value: string;
  }) => React.ReactNode;
};

export default function Select({
  label,
  options,
  notSelectedOption = SELECT_OPTIONS.NOT_SELECTED,
  renderOptionText,
}: SelectFieldProps) {
  const field = useFieldContext<string>();

  const errorMessages = field.state.meta.errors.map(
    (error: { message: string }) => error.message,
  );

  const sortedOptions = [...options].sort((a, b) =>
    a.value.localeCompare(b.value),
  );

  const optionsWithEmpty = [notSelectedOption, ...sortedOptions];

  const selectedOption =
    optionsWithEmpty.find((opt) => opt.key === field.state.value) ?? null;

  return (
    <>
      <Autocomplete
        options={optionsWithEmpty}
        getOptionLabel={(option) => option.value}
        isOptionEqualToValue={(option, value) => option.key === value.key}
        value={selectedOption}
        onChange={(_, newValue) => {
          field.handleChange(newValue?.key ?? '');
        }}
        renderInput={(params) => (
          <TextField
            {...params}
            placeholder="Search..."
            label={label}
            error={errorMessages.length > 0}
            margin="normal"
            fullWidth
          />
        )}
        renderOption={(props, option) => (
          <li
            {...props}
            key={option.key}
          >
            {renderOptionText ? renderOptionText(option) : option.value}
          </li>
        )}
      />

      <FieldErrorText
        marginTop="0.18rem"
        messages={errorMessages}
      />
    </>
  );
}

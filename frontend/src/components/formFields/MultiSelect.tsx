import { Autocomplete, Checkbox, Chip, TextField } from '@mui/material';
import { useFieldContext } from 'hooks/form/useAppForm';
import FieldErrorText from './FieldErrorText';

type MultiSelectFieldProps = {
  label: string;
  options: { key: string; value: string }[];
  renderOptionText?: (option: {
    key: string;
    value: string;
  }) => React.ReactNode;
};

export default function MultiSelect({
  label,
  options,
  renderOptionText,
}: MultiSelectFieldProps) {
  const field = useFieldContext<string[]>();

  const errorMessages = field.state.meta.errors.map(
    (error: { message: string }) => error.message,
  );

  const sortedOptions = [...options].sort((a, b) =>
    a.value.localeCompare(b.value),
  );

  const selectedOptions = options.filter((opt) =>
    field.state.value.includes(opt.key),
  );

  return (
    <>
      <Autocomplete
        multiple
        options={sortedOptions}
        getOptionLabel={(option) => option.value}
        isOptionEqualToValue={(option, value) => option.key === value.key}
        value={selectedOptions}
        onChange={(_, newValue) => {
          const newKeys = newValue.map((v) => v.key);
          field.handleChange(newKeys);
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
        renderTags={(value, getTagProps) =>
          value.map((option, index) => (
            <Chip
              {...getTagProps({ index })}
              key={option.key}
              label={renderOptionText ? renderOptionText(option) : option.value}
            />
          ))
        }
        renderOption={(props, option, { selected }) => (
          <li
            {...props}
            key={option.key}
          >
            <Checkbox
              key={Math.random()}
              style={{ marginRight: 8 }}
              checked={selected}
            />
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

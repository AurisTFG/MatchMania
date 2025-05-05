import { styled } from '@mui/material/styles';
import { DatePicker as MuiDatePicker } from '@mui/x-date-pickers/DatePicker';
import dayjs from 'dayjs';
import { useFieldContext } from 'hooks/form/useAppForm';
import FieldErrorText from './FieldErrorText';

type DatePickerProps = {
  label: string;
};

const StyledDatePicker = styled(MuiDatePicker)(({ theme }) => ({
  '& .MuiPickersOutlinedInput-notchedOutline': {
    borderColor:
      theme.palette.mode === 'dark'
        ? 'rgba(255,255,255,0.6)'
        : 'rgba(0,0,0,0.23)',
  },
}));

export default function DatePicker({ label }: DatePickerProps) {
  const field = useFieldContext<Date | null>();

  const errorMessages = field.state.meta.errors.map(
    (error: { message: string }) => error.message,
  );

  return (
    <>
      <StyledDatePicker
        label={label}
        value={field.state.value ? dayjs(field.state.value) : null}
        onChange={(date: dayjs.Dayjs | null) => {
          field.handleChange(date ? date.toDate() : null);
        }}
        slotProps={{
          textField: {
            error: errorMessages.length > 0,
            fullWidth: true,
            margin: 'normal',
          },
        }}
      />

      <FieldErrorText messages={errorMessages} />
    </>
  );
}

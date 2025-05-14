import { styled } from '@mui/material/styles';
import { DateTimePicker as MuiDateTimePicker } from '@mui/x-date-pickers/DateTimePicker';
import dayjs from 'dayjs';
import { useFieldContext } from 'hooks/form/useAppForm';
import FieldErrorText from './FieldErrorText';

type DateTimePickerProps = {
  label: string;
};

const StyledDateTimePicker = styled(MuiDateTimePicker)(({ theme }) => ({
  '& .MuiPickersOutlinedInput-notchedOutline': {
    borderColor:
      theme.palette.mode === 'dark'
        ? 'rgba(255,255,255,0.6)'
        : 'rgba(0,0,0,0.23)',
  },
}));

export default function DateTimePicker({ label }: DateTimePickerProps) {
  const field = useFieldContext<Date | null>();

  const errorMessages = field.state.meta.errors.map(
    (error: { message: string }) => error.message,
  );

  return (
    <>
      <StyledDateTimePicker
        label={label}
        value={field.state.value ? dayjs(field.state.value) : null}
        onChange={(date: dayjs.Dayjs | null) => {
          field.handleChange(date ? date.toDate() : null);
        }}
        ampm={false}
        disablePast={false}
        format="YYYY-MM-DD HH:mm"
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

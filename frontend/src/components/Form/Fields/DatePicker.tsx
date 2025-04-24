import { DatePicker as MuiDatePicker } from '@mui/x-date-pickers/DatePicker';
import dayjs from 'dayjs';
import { useFieldContext } from 'hooks/form/useAppForm';
import FormErrors from '../Helpers/FormErrors';

export default function DatePicker({ label }: { label: string }) {
  const field = useFieldContext<Date | null>();

  const errorMessages = field.state.meta.errors.map(
    (error: { message: string }) => error.message,
  );

  return (
    <>
      <MuiDatePicker
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
      <FormErrors messages={errorMessages} />
    </>
  );
}

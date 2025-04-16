import { Button } from '@mui/material'; // Import MUI Button
import { useFormContext } from '../../../hooks/form/useAppForm';

export default function SubmitButton({ label }: { label: string }) {
  const form = useFormContext();

  return (
    <form.Subscribe selector={(state) => state.isSubmitting}>
      {(isSubmitting) => (
        <Button
          type="submit"
          disabled={isSubmitting}
          variant="contained"
          color="primary"
        >
          {label}
        </Button>
      )}
    </form.Subscribe>
  );
}

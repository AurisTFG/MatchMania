import { Button } from '@mui/material'; // Import MUI Button
import { useFormContext } from '../../../hooks/form/useAppForm';

export default function SubmitButton({ label }: { label: string }) {
  const form = useFormContext();

  return (
    <form.Subscribe selector={(state) => [state.canSubmit, state.isSubmitting]}>
      {([canSubmit, isSubmitting]) => (
        <Button
          type="submit"
          disabled={!canSubmit}
          variant="contained"
          color="primary"
        >
          {isSubmitting ? '...' : label}
        </Button>
      )}
    </form.Subscribe>
  );
}

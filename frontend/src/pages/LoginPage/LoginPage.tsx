import { Box, Typography } from '@mui/material';
import { useEffect } from 'react';
import { useSearchParams } from 'react-router-dom';
import { toast } from 'sonner';
import { useLogIn } from '../../api/hooks/authHooks';
import { useAppForm } from '../../hooks/form/useAppForm';
import { loginDtoValidator } from '../../validators/auth/loginDtoValidator';

export default function LoginPage() {
  const [searchParams] = useSearchParams();
  const errorMessage = searchParams.get('error');

  const { mutateAsync: loginAsync } = useLogIn();

  const form = useAppForm({
    defaultValues: {
      email: '',
      password: '',
    },
    validators: {
      onSubmit: loginDtoValidator,
    },
    onSubmit: async ({ value }) => {
      await loginAsync(value);
    },
  });

  useEffect(() => {
    if (errorMessage) {
      toast.error(errorMessage);
    }
  }, [errorMessage]);

  return (
    <Box sx={{ maxWidth: 400, mx: 'auto', mt: 5 }}>
      <form
        onSubmit={(e) => {
          e.preventDefault();
          void form.handleSubmit();
        }}
      >
        <Typography
          variant="h4"
          sx={{
            mb: 3,
            fontWeight: 'bold',
          }}
        >
          Login
        </Typography>

        <form.AppField name="email">
          {(field) => <field.TextField label="Email" />}
        </form.AppField>

        <form.AppField name="password">
          {(field) => (
            <field.TextField
              label="Password"
              type="password"
            />
          )}
        </form.AppField>

        <form.AppForm>
          <form.SubmitButton label="Login" />
        </form.AppForm>
      </form>
    </Box>
  );
}

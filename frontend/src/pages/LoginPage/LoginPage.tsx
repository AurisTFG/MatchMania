import { Box, Button, Paper, Typography } from '@mui/material';
import { useEffect } from 'react';
import { useSearchParams } from 'react-router-dom';
import { toast } from 'sonner';
import { useLogIn } from 'api/hooks/authHooks';
import { useAppForm } from 'hooks/form/useAppForm';
import { loginDtoValidator } from 'validators/auth/loginDtoValidator';

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
    <Box
      sx={{
        display: 'flex',
        flexDirection: 'column',
        alignItems: 'center',
        justifyContent: 'center',
        height: 'calc(100vh - 112px)',
      }}
    >
      <Paper
        elevation={3}
        sx={{
          width: '100%',
          maxWidth: 400,
          p: 4,
          borderRadius: 3,
          backgroundColor: 'background.paper',
        }}
      >
        <Typography
          variant="h4"
          component="h1"
          sx={{
            mb: 4,
            fontWeight: 'bold',
            textAlign: 'center',
            color: 'text.primary',
          }}
        >
          Log in
        </Typography>

        <form
          onSubmit={(e) => {
            e.preventDefault();
            void form.handleSubmit();
          }}
        >
          <form.AppField name="email">
            {(field) => <field.Text label="Email" />}
          </form.AppField>

          <form.AppField name="password">
            {(field) => (
              <field.Text
                label="Password"
                type="password"
              />
            )}
          </form.AppField>

          <Box sx={{ mt: 3 }}>
            <form.AppForm>
              <Button
                type="submit"
                fullWidth
                variant="contained"
                color="primary"
                size="large"
                sx={{ borderRadius: 2 }}
              >
                Log in
              </Button>
            </form.AppForm>
          </Box>
        </form>
      </Paper>
    </Box>
  );
}

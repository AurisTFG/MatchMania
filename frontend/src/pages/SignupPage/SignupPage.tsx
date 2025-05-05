import { Box, Button, Paper, Typography } from '@mui/material';
import { useSignUp } from 'api/hooks/authHooks';
import withAuth from 'hocs/withAuth';
import { useAppForm } from 'hooks/form/useAppForm';
import { Permission } from 'types/enums/permission';
import { signupDtoValidator } from 'validators/auth/signupDtoValidator';

function SignupPage() {
  const { mutateAsync: signupAsync } = useSignUp();

  const form = useAppForm({
    defaultValues: {
      username: '',
      email: '',
      password: '',
      confirmPassword: '',
    },
    validators: {
      onSubmit: signupDtoValidator,
    },
    onSubmit: async ({ value }) => {
      await signupAsync(value);
    },
  });

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
          Sign up
        </Typography>

        <form
          onSubmit={(e) => {
            e.preventDefault();
            void form.handleSubmit();
          }}
        >
          <form.AppField name="username">
            {(field) => <field.Text label="Username" />}
          </form.AppField>

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

          <form.AppField name="confirmPassword">
            {(field) => (
              <field.Text
                label="Confirm Password"
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
                Sign up
              </Button>
            </form.AppForm>
          </Box>
        </form>
      </Paper>
    </Box>
  );
}

export default withAuth(SignupPage, {
  permission: Permission.LoggedOut,
});

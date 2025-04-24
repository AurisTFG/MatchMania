import { Box } from '@mui/material';
import { useSignUp } from 'api/hooks/authHooks';
import { useAppForm } from 'hooks/form/useAppForm';
import { signupDtoValidator } from 'validators/auth/signupDtoValidator';

export default function SignupPage() {
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
    <Box sx={{ maxWidth: 400, mx: 'auto', mt: 5 }}>
      <form
        onSubmit={(e) => {
          e.preventDefault();
          void form.handleSubmit();
        }}
      >
        <h1>Sign up</h1>

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

        <form.AppForm>
          <form.SubmitButton label="Sign up" />
        </form.AppForm>
      </form>
    </Box>
  );
}

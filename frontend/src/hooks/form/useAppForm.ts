import { createFormHook, createFormHookContexts } from '@tanstack/react-form';
import { lazy } from 'react';

const TextField = lazy(() => import('../../components/Form/Fields/TextField'));
const DatePicker = lazy(
  () => import('../../components/Form/Fields/DatePicker'),
);

const SubmitButton = lazy(
  () => import('../../components/Form/Controls/SubmitButton'),
);

export const { fieldContext, useFieldContext, formContext, useFormContext } =
  createFormHookContexts();

export const { useAppForm } = createFormHook({
  fieldComponents: {
    TextField,
    DatePicker,
  },
  formComponents: {
    SubmitButton,
  },
  fieldContext,
  formContext,
});

import { createFormHook, createFormHookContexts } from '@tanstack/react-form';
import { lazy } from 'react';

const Text = lazy(() => import('components/Form/Fields/Text'));
const Select = lazy(() => import('components/Form/Fields/Select'));
const DatePicker = lazy(() => import('components/Form/Fields/DatePicker'));

const SubmitButton = lazy(
  () => import('components/Form/Controls/SubmitButton'),
);

export const { fieldContext, useFieldContext, formContext, useFormContext } =
  createFormHookContexts();

export const { useAppForm } = createFormHook({
  fieldComponents: {
    Text,
    Select,
    DatePicker,
  },
  formComponents: {
    SubmitButton,
  },
  fieldContext,
  formContext,
});

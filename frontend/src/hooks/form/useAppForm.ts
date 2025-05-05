import { createFormHook, createFormHookContexts } from '@tanstack/react-form';
import { lazy } from 'react';

const Text = lazy(() => import('components/formFields/Text'));
const Select = lazy(() => import('components/formFields/Select'));
const MultiSelect = lazy(() => import('components/formFields/MultiSelect'));
const DatePicker = lazy(() => import('components/formFields/DatePicker'));
const DateTimePicker = lazy(
  () => import('components/formFields/DateTimePicker'),
);

const SubmitButton = lazy(() => import('components/formControls/SubmitButton'));

export const { fieldContext, useFieldContext, formContext, useFormContext } =
  createFormHookContexts();

export const { useAppForm } = createFormHook({
  fieldComponents: {
    Text,
    Select,
    MultiSelect,
    DatePicker,
    DateTimePicker,
  },
  formComponents: {
    SubmitButton,
  },
  fieldContext,
  formContext,
});

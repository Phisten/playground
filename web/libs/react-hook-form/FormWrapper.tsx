import { PropsWithChildren } from 'react';
import { FieldValues, FormProvider, SubmitHandler, UseFormReturn, useForm } from 'react-hook-form';
import { devConsoleLog } from '../devConsoleLog';

// eslint-disable-next-line @typescript-eslint/no-explicit-any
export interface FormWrapperProps<T extends FieldValues = any>
  extends Omit<React.DetailedHTMLProps<React.FormHTMLAttributes<HTMLFormElement>, HTMLFormElement>, 'onSubmit'> {
  methods?: UseFormReturn<T>;
  onSubmit?: SubmitHandler<T>;
}

export const FormWrapper = (props: PropsWithChildren<FormWrapperProps>) => {
  const {
    children,
    onSubmit = () => {
      devConsoleLog('FormWrapper:DefaultEmptyOnSubmit');
    },
    methods: methodsProp,
    ...other
  } = props;
  const emptyMethods = useForm();
  const methods = methodsProp ?? emptyMethods;

  return (
    <FormProvider {...methods}>
      <form
        {...other}
        onSubmit={methods.handleSubmit(onSubmit, (error) => {
          devConsoleLog('handleSubmit onInvalid: ', error);
        })}
      >
        {children}
      </form>
    </FormProvider>
  );
};

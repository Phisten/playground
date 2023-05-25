import { useIsClient } from 'usehooks-ts';
import styles from './mzn.module.scss';
import clsx from 'clsx';
import {
  FormFieldsDebug,
  FormFieldsWrapper,
  InputField,
} from '@mezzanine-ui/react-hook-form';
import { useForm } from 'react-hook-form';

type DataType = typeof dataList[0];
const dataList = [
  { id: '0', name: 'a1', age: 12 },
  { id: '2', name: 'a12', age: 13123 },
  { id: '3', name: 'bb2', age: 22 },
  { id: '4', name: 'a3332', age: 6 },
  { id: '5', name: 'a144242', age: 20 },
  { id: '6', name: 'a2312', age: 19 },
];

export const MznRfh = () => {
  const methods = useForm();
  return (
    <>
      <FormFieldsWrapper methods={methods}>
        <FormFieldsDebug />
        <InputField registerName="InputField1"></InputField>
      </FormFieldsWrapper>
    </>
  );
};

export const Page = () => {
  const isClient = useIsClient();

  return (
    <>
      <p>mzn rhf highlight</p>
      <a
        href={
          'https://codesandbox.io/s/mzn-table-row-highlight-j76m42?file=/src/App.js'
        }
        target={'_blank'}
        rel="noreferrer"
      >
        <p>codeSandbox</p>
      </a>
      <div className={clsx('w-screen h-screen', styles['wrapper'])}>
        {isClient && <MznRfh />}
      </div>
    </>
  );
};

export default Page;

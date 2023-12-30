import { is } from 'ramda';
import { PropsWithChildren, ReactNode } from 'react';

/* eslint-disable-next-line */
export interface BaseFieldProps {
  label?: ReactNode;
}

export function BaseField(props: PropsWithChildren<BaseFieldProps>) {
  const { children, label = null } = props;

  return (
    <div className="grid gap-2">
      {is(String, label) ? <label className="text-left">{label}</label> : label}
      {children}
    </div>
  );
}

export default BaseField;
